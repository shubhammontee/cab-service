// Package classification CabBookingApp.
//
// Documentation of our CabBooking API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:8000
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package docs

import (
	"github.com/cab_booking_app/models"
)

// swagger:route POST /bookCab bookCab bookingId
// bookCab books a cab from souce to destination location
// responses:
//   200: booking

// swagger:response booking
type bookingWrapper struct {
	// in:body
	Body models.Booking
}

// swagger:parameters bookingId
type foobarRequestWrapper struct {
	// in:body
	Body models.BookingRequest
}

// swagger:route GET /getBookingHistory getBookingHistory listOfBooking
// getBookingHistory gets the booking history of particular user
// responses:
//   200: booking

// swagger:response bookingHistory
type bookingHistoryResponseWrapper struct {
	// in:body
	Body []models.Booking
}

// swagger:parameters listOfBooking
type bookingHistoryRequestWrapper struct {
	// in:query
	Email string `json:"email,omitempty"`
}

// swagger:route GET /searchCab searchCab listOfAvailableCabs
// searchCab gets the list of cabs available within the defined search radius perimeter
// responses:
//   200: availableCabs

// swagger:response availableCabs
type searchCabsResponseWrapper struct {
	// in:body
	Body []models.AvailableCabs
}

// swagger:parameters listOfAvailableCabs
type searchCabsRequestWrapper struct {
	// in:query
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}
