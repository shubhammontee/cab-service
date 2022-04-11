package repository

import (
	"database/sql"
	"fmt"

	"github.com/cab_booking_app/models"
	"github.com/cab_booking_app/utils/errors"
)

//CabRepositoryInterface ...
type CabRepositoryInterface interface {
	BookCab(*models.Booking) *errors.RestErr
	GetAvailableCabs(*[]*models.AvailableCabs) *errors.RestErr
	GetBookingHistory(*[]*models.Booking) *errors.RestErr
}

type cabRepository struct {
	dbClient *sql.DB
}

//NewCabRepo ...
func NewCabRepo(dbClient *sql.DB) CabRepositoryInterface {
	return &cabRepository{
		dbClient: dbClient,
	}
}

const (
	queryInsertCabBookingTable = "INSERT INTO booking_list(from_latitude,from_longitude,to_latitude,to_longitude,u_email,cab_id,status,date_of_booking,price,distance) VALUES(?,?,?,?,?,?,?,?,?,?);"
	queryGetWaitingCab         = "select * from live_location_with_status where status=?;"
	queryGetUserBookingList    = "select * from booking_list where u_email=?;"
)

func (cr cabRepository) BookCab(cabBooking *models.Booking) *errors.RestErr {

	stmt, err := cr.dbClient.Prepare(queryInsertCabBookingTable)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	cabBooking.Status = "Booking Complete"
	defer stmt.Close()
	insertResult, err := stmt.Exec(cabBooking.SrcLatitude, cabBooking.SrcLongitude, cabBooking.DestLatitude,
		cabBooking.DestLongitude, cabBooking.UEmail, cabBooking.CabId, cabBooking.Status, cabBooking.DateOfBooking,
		cabBooking.Price, cabBooking.Distance)

	if err != nil {
		cabBooking.Status = "Booking In Progress"
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to book cab : %s", err.Error()))
	}
	cabBooking.BookingId, err = insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save booking : %s", err.Error()))
	}
	return nil
}

func (cr cabRepository) GetAvailableCabs(availableCabs *[]*models.AvailableCabs) *errors.RestErr {
	stmt, err := cr.dbClient.Prepare(queryGetWaitingCab)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query("WAITING")
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var availableCab models.AvailableCabs
		var status string
		if err := rows.Scan(&availableCab.Latitude, &availableCab.Longitude, &availableCab.CabId, &status); err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		*availableCabs = append(*availableCabs, &availableCab)
	}
	return nil
}

func (cr cabRepository) GetBookingHistory(bookingHistory *[]*models.Booking) *errors.RestErr {
	stmt, err := cr.dbClient.Prepare(queryGetUserBookingList)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	bookingL := *bookingHistory
	rows, err := stmt.Query(bookingL[0].UEmail)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	*bookingHistory = []*models.Booking{}
	defer rows.Close()
	for rows.Next() {
		var bookings models.Booking
		if err := rows.Scan(&bookings.BookingId, &bookings.SrcLatitude, &bookings.SrcLongitude, &bookings.DestLatitude,
			&bookings.DestLongitude, &bookings.UEmail, &bookings.CabId, &bookings.Status, &bookings.DateOfBooking,
			&bookings.Price, &bookings.Distance); err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		*bookingHistory = append(*bookingHistory, &bookings)
	}
	return nil
}
