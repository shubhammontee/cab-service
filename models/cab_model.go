package models

type BookingRequest struct {
	SrcLatitude   float64 `json:"src_latitude,omitempty"`
	SrcLongitude  float64 `json:"src_longitude,omitempty"`
	UEmail        string  `json:"u_email,omitempty"`
	DestLatitude  float64 `json:"dest_latitude,omitempty"`
	DestLongitude float64 `json:"dest_longitude,omitempty"`
	CabId         string  `json:"cab_id,omitempty"`
	RushHourIndex float64 `json:"rush_hour_index,omitempty"`
}

type Booking struct {
	BookingId     int64   `json:"booking_id,omitempty"`
	SrcLatitude   float64 `json:"src_latitude,omitempty"`
	SrcLongitude  float64 `json:"src_longitude,omitempty"`
	UEmail        string  `json:"u_email,omitempty"`
	DestLatitude  float64 `json:"dest_latitude,omitempty"`
	DestLongitude float64 `json:"dest_longitude,omitempty"`
	Price         float64 `json:"price,omitempty"`
	DateOfBooking string  `json:"date_of_booking,omitempty"`
	CabId         string  `json:"cab_id,omitempty"`
	Status        string  `json:"status,omitempty"`
	Distance      float64 `json:"distance,omitempty"`
}

type AvailableCabs struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	CabId     string  `json:"cab_id,omitempty"`
}
