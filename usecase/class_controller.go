package usecase

import (
	"net/http"
	"strconv"

	"github.com/dennys-bd/glow/entity"
	"github.com/dennys-bd/glow/repository"
	"github.com/gin-gonic/gin"
)

type ClassController struct {
	r repository.ClassRepoInf
}

func ProvideClassCtrl(r repository.ClassRepoInf) (ClassController, error) {
	if r == nil {
		return ClassController{}, errNoRepository("ProvideClassCtrl")
	}
	return ClassController{r: r}, nil
}

func (c *ClassController) Index(ctx *gin.Context) {
	classes, err := c.r.List(nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": classes})
}

func (c *ClassController) Show(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	class, err := c.r.Find(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": class})
}

func (c *ClassController) Create(ctx *gin.Context) {
	var class entity.Class
	if err := ctx.ShouldBindJSON(&class); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.r.Create(&class); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": class})
}
