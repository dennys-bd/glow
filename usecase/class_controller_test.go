package usecase_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"strconv"
	"testing"

	"github.com/dennys-bd/glow/api/router"
	"github.com/dennys-bd/glow/entity"
	"github.com/dennys-bd/glow/entity/factories"
	"github.com/dennys-bd/glow/ops/projectpath"
	"github.com/dennys-bd/glow/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
	"gorm.io/gorm"
)

var Cleaner = dbcleaner.New()

type ClassesSuite struct {
	suite.Suite
	routes *gin.Engine
	dbCtx  context.Context
}

func (s *ClassesSuite) SetupSuite() {
	sqlite := engine.NewSqliteEngine(projectpath.Root() + "/test.db")
	Cleaner.SetEngine(sqlite)

	db := repository.GetDB()
	ctx := context.WithValue(context.Background(), factories.DBContextKey, db)
	s.dbCtx = ctx
}

func (s *ClassesSuite) SetupTest() {
	s.routes = gin.Default()
	router.SetApiRouter(s.routes)
	Cleaner.Acquire("classes")
}

func (s *ClassesSuite) TearDownTest() {
	Cleaner.Clean("classes")
}

func TestRunClassesSuite(t *testing.T) {
	suite.Run(t, new(ClassesSuite))
}

func (s *ClassesSuite) TestClassesIndex() {
	classGroup := factories.ClassGroupFactory.MustCreateWithContextAndOption(s.dbCtx, nil).(*factories.ClassGroup)
	classes := classGroup.Classes
	sort.Slice(classes, func(i, j int) bool {
		return classes[i].ID < classes[j].ID
	})

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("GET", "/api/classes", nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	expectedResponse := Must(json.Marshal(map[string]interface{}{"data": classes})).([]byte)

	s.Equal(http.StatusOK, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *ClassesSuite) TestClassesShow() {
	factories.ClassGroupFactory.MustCreateWithContextAndOption(s.dbCtx, nil)
	class := factories.ClassFactory.MustCreateWithContextAndOption(s.dbCtx, nil).(*entity.Class)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("GET", fmt.Sprintf("/api/classes/%d", class.ID), nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	expectedResponse := Must(json.Marshal(map[string]interface{}{"data": class})).([]byte)

	s.Equal(http.StatusOK, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *ClassesSuite) TestClassesShow_NotFound() {
	classGroup := factories.ClassGroupFactory.MustCreateWithContextAndOption(s.dbCtx, nil).(*factories.ClassGroup)
	size := len(classGroup.Classes)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("GET", fmt.Sprintf("/api/classes/%d", size+1), nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	expectedResponse := Must(json.Marshal(map[string]interface{}{"error": gorm.ErrRecordNotFound.Error()})).([]byte)

	s.Equal(http.StatusNotFound, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *ClassesSuite) TestClassesShow_BadRequest() {
	factories.ClassGroupFactory.MustCreateWithContextAndOption(s.dbCtx, nil)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("GET", "/api/classes/a", nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	err := strconv.NumError{
		Func: "ParseUint",
		Num:  "a",
		Err:  strconv.ErrSyntax}
	expectedResponse := Must(json.Marshal(map[string]interface{}{"error": err.Error()})).([]byte)

	s.Equal(http.StatusBadRequest, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *ClassesSuite) TestClassesCreation() {

	buildClass := factories.ClassFactory.MustCreate().(*entity.Class)
	expectedClass := PostClass{
		Name:      buildClass.Name,
		StartDate: buildClass.StartDate,
		EndDate:   buildClass.EndDate,
	}
	reqBody := Must(json.Marshal(expectedClass)).([]byte)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("POST", "/api/classes", bytes.NewBuffer(reqBody))).(*http.Request)
	s.routes.ServeHTTP(w, req)

	r := bytes.NewReader(w.Body.Bytes())
	decoder := json.NewDecoder(r)
	var dataClass DataClass
	if err := decoder.Decode(&dataClass); err != nil {
		panic(err)
	}

	s.Equal(http.StatusCreated, w.Code)
	s.Equal(expectedClass.Name, dataClass.Class.Name)
	s.Equal(expectedClass.StartDate, dataClass.Class.StartDate)
	s.Equal(expectedClass.EndDate, dataClass.Class.EndDate)
	s.NotZero(dataClass.Class.ID)
}
