package service

import (
	"fmt"
	"testing"

	"github.com/cab_booking_app/models"
	"github.com/cab_booking_app/utils/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) BookCab(bookingReq *models.Booking) *errors.RestErr {
	bookingReq.BookingId = 1
	bookingReq.Status = "Booking Complete"
	return nil
}

func (mock *MockRepository) GetAvailableCabs(availableCabs *[]*models.AvailableCabs) *errors.RestErr {
	cabAt := &models.AvailableCabs{
		Latitude: 12.12,
	}
	*availableCabs = append(*availableCabs, cabAt)
	return nil
}

func (mock *MockRepository) GetBookingHistory(bookings *[]*models.Booking) *errors.RestErr {
	bookingRes := &models.Booking{
		BookingId:     1,
		SrcLatitude:   32.9697,
		SrcLongitude:  -96.80322,
		UEmail:        "user1@gmail.com",
		DestLatitude:  29.46786,
		DestLongitude: -98.53506,
		CabId:         "CAB1",
		Price:         100.50,
		DateOfBooking: "22 10 2020",
		Status:        "Booking Complete",
		Distance:      50.4,
	}
	*bookings = append(*bookings, bookingRes)
	return nil
}

func TestBookCab(t *testing.T) {
	mockRepo := new(MockRepository)
	bookingReq1 := &models.Booking{
		SrcLatitude:   32.9697,
		SrcLongitude:  -96.80322,
		UEmail:        "user1@gmail.com",
		DestLatitude:  29.46786,
		DestLongitude: -98.53506,
		CabId:         "CAB1",
		Price:         3940.167,
		DateOfBooking: "2020-11-06 01:02:41",
		Status:        "Booking Complete",
		Distance:      262.6778,
	}

	mockRepo.On("BookCab", bookingReq1).Return(nil)
	bookingReq := models.BookingRequest{
		SrcLatitude:   32.9697,
		SrcLongitude:  -96.80322,
		UEmail:        "user1@gmail.com",
		DestLatitude:  29.46786,
		DestLongitude: -98.53506,
		CabId:         "CAB1",
		RushHourIndex: 1.50,
	}
	testService := NewService(mockRepo)
	res, _ := testService.BookCab(bookingReq)

	//mockRepo.AssertExpectations(t)
	assert.Equal(t, int64(1), res.BookingId)
	assert.Equal(t, "Booking Complete", res.Status)

}

func TestGetAvailableCabs(t *testing.T) {
	mockRepo := new(MockRepository)
	lat := 12.12
	long := 76.68
	bookingRes := &models.AvailableCabs{
		Latitude:  13.12,
		Longitude: 76.68,
		CabId:     "CAB2",
	}
	var expextedMockReq []*models.AvailableCabs
	expextedMockReq = append(expextedMockReq, bookingRes)
	mockRepo.On("GetAvailableCabs", &expextedMockReq).Return(nil)

	testService := NewService(mockRepo)
	res, _ := testService.GetAvailableCabs(lat, long)

	assert.Equal(t, 12.12, res[0].Latitude)

}

func TestGetBookingHistory(t *testing.T) {
	mockRepo := new(MockRepository)
	u_email := "user1@gmail.com"

	bookingRes := &models.Booking{
		BookingId:     1,
		SrcLatitude:   32.9697,
		SrcLongitude:  -96.80322,
		UEmail:        "user1@gmail.com",
		DestLatitude:  29.46786,
		DestLongitude: -98.53506,
		CabId:         "CAB1",
		Price:         100.50,
		DateOfBooking: "22 10 2020",
		Status:        "Booking Complete",
		Distance:      50.4,
	}
	var expextedMockReq []*models.Booking
	expextedMockReq = append(expextedMockReq, bookingRes)
	mockRepo.On("GetBookingHistory", &expextedMockReq).Return(nil)

	testService := NewService(mockRepo)
	res, _ := testService.GetBookingHistory(u_email)
	fmt.Println(res)
	assert.Equal(t, int64(1), res[1].BookingId)

}
