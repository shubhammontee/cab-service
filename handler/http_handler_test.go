package handler

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/cab_booking_app/models"
	"github.com/cab_booking_app/utils/calculate_distance"
	"github.com/cab_booking_app/utils/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/suvamsingh/bookstore_users-api/utils/date_utils"
)

type MockRepository struct {
	mock.Mock
}

// BookCab(models.BookingRequest) (models.Booking, *errors.RestErr)
// 	GetAvailableCabs(float64, float64) ([]models.AvailableCabs, *errors.RestErr)
// 	GetBookingHistory(uEmail string) ([]*models.Booking, *errors.RestErr)

func (mock *MockRepository) BookCab(bookingReq models.BookingRequest) (models.Booking, *errors.RestErr) {
	args := mock.Called()
	result := args.Get(0)
	cabBooking := result.(models.BookingRequest)
	distance := calculate_distance.GetDistance(cabBooking.SrcLatitude, cabBooking.SrcLongitude, cabBooking.DestLatitude, cabBooking.DestLongitude)
	price := distance * cabBooking.RushHourIndex * 50
	response := models.Booking{
		SrcLatitude:   cabBooking.SrcLatitude,
		SrcLongitude:  cabBooking.SrcLongitude,
		UEmail:        cabBooking.UEmail,
		DestLatitude:  cabBooking.DestLatitude,
		DestLongitude: cabBooking.DestLongitude,
		Price:         price,
		DateOfBooking: date_utils.GetNowDBFormat(),
		CabId:         cabBooking.CabId,
		Status:        "BookingInProcess",
		Distance:      distance,
	}
	err := args.Error(1)
	return response, &errors.RestErr{Message: err.Error(), Status: 500, Error: "error occured calling Book cab"}
}

func (mock *MockRepository) GetAvailableCabs(lat, long float64) ([]models.AvailableCabs, *errors.RestErr) {
	// args := mock.Called()
	// result := args.Get(0)
	// fmt.P
	return nil, nil

}

func (mock *MockRepository) GetBookingHistory(uEmail string) ([]*models.Booking, *errors.RestErr) {
	return nil, nil
}

func TestBookCab(t *testing.T) {
	mockRepo := new(MockRepository)
	bookingRes := &models.Booking{
		BookingId:     1,
		SrcLatitude:   32.9697,
		SrcLongitude:  -96.80322,
		UEmail:        "user1@gmail.com",
		DestLatitude:  29.46786,
		DestLongitude: -98.53506,
		CabId:         "CAB1",
		Price:         3940.167,
		DateOfBooking: "2020-11-06 01:02:41",
		Status:        "BookingInProcess",
		Distance:      262.6778,
	}
	bookingReq := models.BookingRequest{
		SrcLatitude:   32.9697,
		SrcLongitude:  -96.80322,
		UEmail:        "user1@gmail.com",
		DestLatitude:  29.46786,
		DestLongitude: -98.53506,
		CabId:         "CAB1",
		RushHourIndex: 1.50,
	}

	mockRepo.On("BookCab", bookingReq).Return(bookingRes, nil)
	w := httptest.NewRecorder()
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(bookingReq)
	w.Body = reqBodyBytes

	c, _ := gin.CreateTestContext(w)

	//c.Request.

	entry := log.WithFields(log.Fields{
		"correlationID": "ABC",
	})
	c.Set("Logger", entry)

	testHandler := NewHandler(mockRepo)
	testHandler.BookCab(c)
	assert.Equal(t, 200, w.Code)

}
