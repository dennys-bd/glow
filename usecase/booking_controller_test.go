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
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
	"gorm.io/gorm"
)

type BookingsSuite struct {
	suite.Suite
	routes *gin.Engine
	dbCtx  context.Context
}

func (s *BookingsSuite) SetupSuite() {
	sqlite := engine.NewSqliteEngine(projectpath.Root() + "/test.db")
	Cleaner.SetEngine(sqlite)

	db := repository.GetDB()
	ctx := context.WithValue(context.Background(), factories.DBContextKey, db)
	s.dbCtx = ctx

	s.routes = gin.Default()
	router.SetApiRouter(s.routes)
}

func (s *BookingsSuite) SetupTest() {
	Cleaner.Acquire("bookings")
	Cleaner.Acquire("classes")
}

func (s *BookingsSuite) TearDownTest() {
	Cleaner.Clean("bookings")
	Cleaner.Clean("classes")
}

func TestRunBookingsSuite(t *testing.T) {
	suite.Run(t, new(BookingsSuite))
}

func (s *BookingsSuite) TestBookingsIndex() {
	bookingsGroup := factories.BookingGroupFactory.MustCreateWithContextAndOption(s.dbCtx, nil).(*factories.BookingGroup)
	bookings := bookingsGroup.Bookings
	sort.Slice(bookings, func(i, j int) bool {
		return bookings[i].ID < bookings[j].ID
	})

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("GET", "/api/bookings", nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	expectedResponse := Must(json.Marshal(map[string]interface{}{"data": bookings})).([]byte)

	println(w.Body.String())
	s.Equal(http.StatusOK, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *BookingsSuite) TestBookingsShow() {
	factories.BookingGroupFactory.MustCreateWithContextAndOption(s.dbCtx, nil)
	booking := factories.BookingFactory.MustCreateWithContextAndOption(s.dbCtx, nil).(*entity.Booking)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("GET", fmt.Sprintf("/api/bookings/%d", booking.ID), nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	expectedResponse := Must(json.Marshal(map[string]interface{}{"data": booking})).([]byte)

	s.Equal(http.StatusOK, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *BookingsSuite) TestBookingsShow_NotFound() {
	classGroup := factories.BookingGroupFactory.MustCreateWithContextAndOption(s.dbCtx, nil).(*factories.BookingGroup)
	size := len(classGroup.Bookings)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("GET", fmt.Sprintf("/api/bookings/%d", size+1), nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	expectedResponse := Must(json.Marshal(map[string]interface{}{"error": gorm.ErrRecordNotFound.Error()})).([]byte)

	s.Equal(http.StatusNotFound, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *BookingsSuite) TestBookingsShow_BadRequest() {
	factories.BookingGroupFactory.MustCreateWithContextAndOption(s.dbCtx, nil)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("GET", "/api/bookings/a", nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	err := strconv.NumError{
		Func: "ParseUint",
		Num:  "a",
		Err:  strconv.ErrSyntax}
	expectedResponse := Must(json.Marshal(map[string]interface{}{"error": err.Error()})).([]byte)

	s.Equal(http.StatusBadRequest, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *BookingsSuite) TestBookingsCreation() {

	buildBooking := factories.BookingFactory.MustCreate().(*entity.Booking)
	class := factories.ClassFactory.MustCreateWithContextAndOption(s.dbCtx, nil).(*entity.Class)
	expectedBooking := PostBooking{
		Name:    buildBooking.Name,
		Date:    class.StartDate,
		ClassID: class.ID,
	}
	reqBody := Must(json.Marshal(expectedBooking)).([]byte)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("POST", "/api/bookings", bytes.NewBuffer(reqBody))).(*http.Request)
	s.routes.ServeHTTP(w, req)

	r := bytes.NewReader(w.Body.Bytes())
	decoder := json.NewDecoder(r)
	var data DataBooking
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}

	println(w.Body.String())
	s.Equal(http.StatusCreated, w.Code)
	s.Equal(expectedBooking.Name, data.Booking.Name)
	s.Equal(expectedBooking.Date, data.Booking.Date)
	s.Equal(expectedBooking.ClassID, data.Booking.ClassID)
	s.NotZero(data.Booking.ID)
}

func (s *BookingsSuite) TestBookingsCreation_BadParameters() {

	postBooking := PostBooking{}
	reqBody := Must(json.Marshal(postBooking)).([]byte)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("POST", "/api/bookings", bytes.NewBuffer(reqBody))).(*http.Request)
	s.routes.ServeHTTP(w, req)

	r := bytes.NewReader(w.Body.Bytes())
	decoder := json.NewDecoder(r)
	var data DataBooking
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}

	err := validator.New().Struct(Booking{})
	expectedResponse := Must(json.Marshal(map[string]interface{}{"error": err.Error()})).([]byte)

	s.Equal(http.StatusBadRequest, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *BookingsSuite) TestBookingsCreation_NonExistentClass() {

	buildBooking := factories.BookingFactory.MustCreate().(*entity.Booking)
	postBooking := PostBooking{
		Name:    buildBooking.Name,
		Date:    buildBooking.Date,
		ClassID: 1,
	}
	reqBody := Must(json.Marshal(postBooking)).([]byte)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("POST", "/api/bookings", bytes.NewBuffer(reqBody))).(*http.Request)
	s.routes.ServeHTTP(w, req)

	r := bytes.NewReader(w.Body.Bytes())
	decoder := json.NewDecoder(r)
	var data DataBooking
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}

	expectedResponse := Must(json.Marshal(map[string]interface{}{"error": gorm.ErrRecordNotFound.Error()})).([]byte)

	s.Equal(http.StatusBadRequest, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *BookingsSuite) TestBookingsCreation_Date_GT_ClassEndDate() {

	buildBooking := factories.BookingFactory.MustCreate().(*entity.Booking)
	class := factories.ClassFactory.MustCreateWithContextAndOption(s.dbCtx, nil).(*entity.Class)
	postBooking := PostBooking{
		Name:    buildBooking.Name,
		Date:    class.EndDate.Add(1),
		ClassID: class.ID,
	}
	reqBody := Must(json.Marshal(postBooking)).([]byte)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("POST", "/api/bookings", bytes.NewBuffer(reqBody))).(*http.Request)
	s.routes.ServeHTTP(w, req)

	r := bytes.NewReader(w.Body.Bytes())
	decoder := json.NewDecoder(r)
	var data DataBooking
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}

	booking := Booking{
		Name:    buildBooking.Name,
		Date:    postBooking.Date,
		ClassID: class.ID,
		Class: &Class{
			StartDate: class.StartDate,
			EndDate:   class.EndDate,
		},
	}
	err := validator.New().Struct(booking)
	expectedResponse := Must(json.Marshal(map[string]interface{}{"error": err.Error()})).([]byte)

	s.Equal(http.StatusUnprocessableEntity, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *BookingsSuite) TestBookingsCreation_Date_LT_ClassStartDate() {

	buildBooking := factories.BookingFactory.MustCreate().(*entity.Booking)
	class := factories.ClassFactory.MustCreateWithContextAndOption(s.dbCtx, nil).(*entity.Class)
	postBooking := PostBooking{
		Name:    buildBooking.Name,
		Date:    class.StartDate.Add(-1),
		ClassID: class.ID,
	}
	reqBody := Must(json.Marshal(postBooking)).([]byte)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("POST", "/api/bookings", bytes.NewBuffer(reqBody))).(*http.Request)
	s.routes.ServeHTTP(w, req)

	r := bytes.NewReader(w.Body.Bytes())
	decoder := json.NewDecoder(r)
	var data DataBooking
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}

	booking := Booking{
		Name:    buildBooking.Name,
		Date:    postBooking.Date,
		ClassID: class.ID,
		Class: &Class{
			StartDate: class.StartDate,
			EndDate:   class.EndDate,
		},
	}
	err := validator.New().Struct(booking)
	expectedResponse := Must(json.Marshal(map[string]interface{}{"error": err.Error()})).([]byte)

	s.Equal(http.StatusUnprocessableEntity, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *BookingsSuite) TestBookingsDeletion() {
	booking := factories.BookingFactory.MustCreateWithContextAndOption(s.dbCtx, nil).(*entity.Booking)

	w := httptest.NewRecorder()
	req := Must(http.NewRequest("DELETE", fmt.Sprintf("/api/bookings/%d", booking.ID), nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	s.Equal(http.StatusNoContent, w.Code)
	s.Equal("", w.Body.String())

	w = httptest.NewRecorder()
	req = Must(http.NewRequest("GET", fmt.Sprintf("/api/bookings/%d", booking.ID), nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	expectedResponse := Must(json.Marshal(map[string]interface{}{"error": gorm.ErrRecordNotFound.Error()})).([]byte)

	s.Equal(http.StatusNotFound, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *BookingsSuite) TestBookingsDeletion_NotFound() {
	w := httptest.NewRecorder()
	req := Must(http.NewRequest("DELETE", "/api/bookings/1", nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	expectedResponse := Must(json.Marshal(map[string]interface{}{"error": gorm.ErrRecordNotFound.Error()})).([]byte)

	s.Equal(http.StatusNotFound, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}

func (s *BookingsSuite) TestBookingsDeletion_BadRequest() {
	w := httptest.NewRecorder()
	req := Must(http.NewRequest("DELETE", "/api/bookings/a", nil)).(*http.Request)
	s.routes.ServeHTTP(w, req)

	err := strconv.NumError{
		Func: "ParseUint",
		Num:  "a",
		Err:  strconv.ErrSyntax}
	expectedResponse := Must(json.Marshal(map[string]interface{}{"error": err.Error()})).([]byte)

	s.Equal(http.StatusBadRequest, w.Code)
	s.Equal(string(expectedResponse), w.Body.String())
}
