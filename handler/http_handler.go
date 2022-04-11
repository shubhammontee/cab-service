package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cab_booking_app/models"
	"github.com/cab_booking_app/service"
	"github.com/cab_booking_app/utils/errors"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type CabBookingHandlerInterface interface {
	BookCab(c *gin.Context)
	SearchCab(c *gin.Context)
	GetBookingHistory(c *gin.Context)
}
type cabBookingHandler struct {
	service service.Service
}

func NewHandler(service service.Service) CabBookingHandlerInterface {
	return &cabBookingHandler{
		service: service,
	}
}

func (handler cabBookingHandler) BookCab(c *gin.Context) {
	lg, _ := c.Get("Logger")
	logger := lg.(*log.Entry)

	var cabBooking models.BookingRequest
	if err := c.ShouldBindJSON(&cabBooking); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
	}
	logger.Info(cabBooking)
	res, err := handler.service.BookCab(cabBooking)
	if err != nil {
		logger.Fatal(err)
		c.JSON(err.Status, err)
		return
	}
	fmt.Println(c.Request.Header.Get("CorrelationID"))
	c.JSON(http.StatusOK, res)

}

func (handler cabBookingHandler) SearchCab(c *gin.Context) {
	lg, _ := c.Get("Logger")
	logger := lg.(*log.Entry)
	lat := c.Query("latitude")
	long := c.Query("longitude")

	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		restErr := errors.NewBadRequestError(fmt.Sprintf("Invalid Json Body: %s", err.Error()))
		logger.Fatal(restErr)
		c.JSON(restErr.Status, restErr)
	}

	longitude, err := strconv.ParseFloat(long, 64)
	if err != nil {
		restErr := errors.NewBadRequestError(fmt.Sprintf("Invalid Json Body: %s", err.Error()))
		logger.Fatal(restErr)
		c.JSON(restErr.Status, restErr)
	}
	logger.Info(fmt.Sprintf("Latitude : %f Longitude: %f", latitude, longitude))
	res, rErr := handler.service.GetAvailableCabs(latitude, longitude)
	if rErr != nil {
		logger.Fatal(rErr)
		c.JSON(rErr.Status, rErr)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (handler cabBookingHandler) GetBookingHistory(c *gin.Context) {
	lg, _ := c.Get("Logger")
	logger := lg.(*log.Entry)
	uEmail := c.Query("email")
	logger.Info(fmt.Sprintf("user email : %s", uEmail))

	res, rErr := handler.service.GetBookingHistory(uEmail)
	if rErr != nil {
		logger.Fatal(rErr)
		c.JSON(rErr.Status, rErr)
		return
	}
	c.JSON(http.StatusOK, res)

}
