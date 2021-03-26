package usecase

import (
	"net/http"
	"strconv"

	"github.com/dennys-bd/glow/entity"
	"github.com/dennys-bd/glow/repository"
	"github.com/gin-gonic/gin"
)

type BookingController struct {
	r  repository.BookingRepoInf
	cr repository.ClassRepoInf
}

func ProvideBookingCtrl(r repository.BookingRepoInf, cr repository.ClassRepoInf) (BookingController, error) {
	if r == nil {
		return BookingController{}, errNoRepository("ProvideBookingCtrl")
	}

	if cr == nil {
		return BookingController{}, errNoRepository("ProvideBookingCtrl")
	}

	return BookingController{r: r, cr: cr}, nil
}

func (c *BookingController) Index(ctx *gin.Context) {
	bookings, err := c.r.List(nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": bookings})
}

func (c *BookingController) Show(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	booking, err := c.r.Find(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": booking})
}

func (c *BookingController) Create(ctx *gin.Context) {
	var booking entity.Booking
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	class, err := c.cr.Find(booking.ClassID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	booking.Class = class

	if err := c.r.Create(&booking); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": booking})
}

func (c BookingController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := c.r.Find(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	if err := c.r.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
