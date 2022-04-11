package repository

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	//r "github.com/cab_booking_app/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cab_booking_app/models"
	"github.com/cab_booking_app/utils/errors"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestBookCab(t *testing.T) {
	db, mock := NewMock()
	repo := &cabRepository{db}
	defer func() {
		repo.dbClient.Close()
	}()
	query := "INSERT INTO booking_list(from_latitude,from_longitude,to_latitude,to_longitude,u_email,cab_id,status,date_of_booking,price,distance) VALUES(?,?,?,?,?,?,?,?,?,?);"

	bookingReq := &models.Booking{
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
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(bookingReq.SrcLatitude, bookingReq.SrcLongitude, bookingReq.DestLatitude,
		bookingReq.DestLongitude, bookingReq.UEmail, bookingReq.CabId, bookingReq.Status, bookingReq.DateOfBooking,
		bookingReq.Price, bookingReq.Distance).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.BookCab(bookingReq)
	var expectedResult *errors.RestErr
	assert.Equal(t, expectedResult, err)
}

func TestBookCabError(t *testing.T) {
	db, mock := NewMock()
	repo := &cabRepository{db}
	defer func() {
		repo.dbClient.Close()
	}()

	query := "INSERT INTO booking_list(from_latitude,from_longitude,to_latitude,to_longitude,u_email,cab_id,status,date_of_booking,price,distance) VALUES(?,?,?,?,?,?,?,?,?,?);"

	bookingReq := &models.Booking{
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
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(bookingReq.SrcLatitude, bookingReq.SrcLongitude, bookingReq.DestLatitude,
		bookingReq.DestLongitude, bookingReq.UEmail, bookingReq.CabId, bookingReq.Status, bookingReq.DateOfBooking,
		bookingReq.Price, bookingReq.Distance, bookingReq.CabId).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.BookCab(bookingReq)
	var expectedResult *errors.RestErr
	assert.NotEqual(t, expectedResult, err)
}

func TestGetBookingHistory(t *testing.T) {
	db, mock := NewMock()
	repo := &cabRepository{db}
	defer func() {
		repo.dbClient.Close()
	}()

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

	query := "select * from booking_list where u_email=?;"

	rows := sqlmock.NewRows([]string{"booking_id", "src_latitude", "src_longitude", "dest_latitude", "dest_longitude", "u_email",
		"cab_id", "status", "date_of_booking", "price", "distance"}).
		AddRow(bookingRes.BookingId, bookingRes.SrcLatitude, bookingRes.SrcLongitude, bookingRes.DestLatitude,
			bookingRes.DestLongitude, bookingRes.UEmail, bookingRes.CabId, bookingRes.Status, bookingRes.DateOfBooking,
			bookingRes.Price, bookingRes.Distance)

	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectQuery().WithArgs(bookingRes.UEmail).WillReturnRows(rows)
	var expectedReQuest []*models.Booking
	expectedReQuest = append(expectedReQuest, bookingRes)
	err := repo.GetBookingHistory(&expectedReQuest)
	var expectedResult *errors.RestErr
	assert.NotNil(t, expectedReQuest)
	assert.Equal(t, expectedResult, err)
}

func TestGetBookingHistoryError(t *testing.T) {
	db, mock := NewMock()
	repo := &cabRepository{db}
	defer func() {
		repo.dbClient.Close()
	}()

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

	query := "select * from booking_list where u_email=?;"

	rows := sqlmock.NewRows([]string{"booking_id", "src_latitude", "src_longitude", "u_email", "dest_latitude", "dest_longitude",
		"price", "date_of_booking", "cab_id", "status", "distance"})

	mock.ExpectQuery(query).WithArgs(bookingRes.UEmail).WillReturnRows(rows)
	var expectedReQuest []*models.Booking
	expectedReQuest = append(expectedReQuest, bookingRes)
	err := repo.GetBookingHistory(&expectedReQuest)
	var expectedResult *errors.RestErr
	assert.NotNil(t, expectedReQuest)
	assert.NotEqual(t, expectedResult, err)
}

func TestGetAvailableCabs(t *testing.T) {
	db, mock := NewMock()
	repo := &cabRepository{db}
	defer func() {
		repo.dbClient.Close()
	}()

	availableCabsRes := &models.AvailableCabs{
		Latitude:  12.12,
		Longitude: 76.68,
		CabId:     "CAB2",
	}

	query := "select * from live_location_with_status where status=?;"

	rows := sqlmock.NewRows([]string{"latitude", "longitude", "cab_id", "status"}).
		AddRow(availableCabsRes.Latitude, availableCabsRes.Longitude, availableCabsRes.CabId, "WAITING")

	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectQuery().WithArgs("WAITING").WillReturnRows(rows)
	var expectedReQuest []*models.AvailableCabs
	err := repo.GetAvailableCabs(&expectedReQuest)
	var expectedResult *errors.RestErr
	assert.NotNil(t, expectedReQuest)
	assert.Equal(t, expectedResult, err)
}

func TestGetAvailableCabsError(t *testing.T) {
	db, mock := NewMock()
	repo := &cabRepository{db}
	defer func() {
		repo.dbClient.Close()
	}()

	availableCabsRes := &models.AvailableCabs{
		Latitude:  12.12,
		Longitude: 76.68,
		CabId:     "CAB2",
	}

	query := "select * from live_location_with_status where status=?;"

	rows := sqlmock.NewRows([]string{"latitude", "longitude", "cab_id"}).
		AddRow(availableCabsRes.Latitude, availableCabsRes.Longitude, availableCabsRes.CabId)

	mock.ExpectQuery(query).WithArgs("WAITING").WillReturnRows(rows)
	var expectedReQuest []*models.AvailableCabs
	err := repo.GetAvailableCabs(&expectedReQuest)
	var expectedResult *errors.RestErr
	assert.Nil(t, expectedReQuest)
	assert.NotEqual(t, expectedResult, err)
}
