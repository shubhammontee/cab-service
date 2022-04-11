package service

import (
	"github.com/cab_booking_app/models"
	"github.com/cab_booking_app/utils/calculate_distance"
	"github.com/cab_booking_app/utils/errors"
	"github.com/suvamsingh/bookstore_users-api/utils/date_utils"

	"github.com/cab_booking_app/repository"
)

const (
	PER_KM_COST          = 10.00
	SPEED_IN_KM_PER_HOUR = 30.00
	SEARCH_RADIUS_IN_KM  = 500000.00
)

//Service ...
type Service interface {
	BookCab(models.BookingRequest) (models.Booking, *errors.RestErr)
	GetAvailableCabs(float64, float64) ([]models.AvailableCabs, *errors.RestErr)
	GetBookingHistory(uEmail string) ([]*models.Booking, *errors.RestErr)
}

type service struct {
	repository repository.CabRepositoryInterface
}

//NewService ...
func NewService(repo repository.CabRepositoryInterface) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) BookCab(cabBooking models.BookingRequest) (models.Booking, *errors.RestErr) {
	distance := calculate_distance.GetDistance(cabBooking.SrcLatitude, cabBooking.SrcLongitude, cabBooking.DestLatitude, cabBooking.DestLongitude)
	price := distance * cabBooking.RushHourIndex * PER_KM_COST
	repoRequest := &models.Booking{
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

	err := s.repository.BookCab(repoRequest)
	if err != nil {
		return models.Booking{}, err
	}
	return *repoRequest, nil
}

func (s *service) GetAvailableCabs(latitude, longitude float64) ([]models.AvailableCabs, *errors.RestErr) {
	var reqAvailableCabs []*models.AvailableCabs
	if err := s.repository.GetAvailableCabs(&reqAvailableCabs); err != nil {
		return nil, err
	}
	var cabUnderSearchRadius []models.AvailableCabs
	for i := 0; i < len(reqAvailableCabs); i++ {
		if calculate_distance.GetDistance(latitude, longitude, reqAvailableCabs[i].Latitude, reqAvailableCabs[i].Longitude) < SEARCH_RADIUS_IN_KM {
			cabUnderSearchRadius = append(cabUnderSearchRadius, *reqAvailableCabs[i])
		}
	}
	return cabUnderSearchRadius, nil
}

func (s *service) GetBookingHistory(uEmail string) ([]*models.Booking, *errors.RestErr) {
	var userBookings []*models.Booking
	var booking models.Booking
	booking.UEmail = uEmail
	userBookings = append(userBookings, &booking)
	if err := s.repository.GetBookingHistory(&userBookings); err != nil {
		return nil, err
	}
	return userBookings, nil
}

// func (s *service) CreateAlbum(album models.Album) (models.Album, *errors.RestErr) {
// 	album.ID = primitive.NewObjectID()
// 	album.Info.CreatedDate = date_utils.GetNowDBFormat()
// 	if err := s.repository.CreateAlbum(&album); err != nil {
// 		return models.Album{}, err
// 	}
// 	err := sendNotification(fmt.Sprintf("new album with album name %s created", album.AlbumName))
// 	if err != nil {
// 		return models.Album{}, nil
// 	}
// 	return album, nil
// }

// func (s *service) DeleteAlbum(id string) (models.Album, *errors.RestErr) {
// 	var album models.Album
// 	oId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return models.Album{}, errors.NewInternalServerError(err.Error())
// 	}
// 	album.ID = oId
// 	if err := s.repository.DeleteAlbum(&album); err != nil {
// 		return models.Album{}, err
// 	}
// 	restErr := sendNotification(fmt.Sprintf("album with album name %s deleted", album.AlbumName))
// 	if restErr != nil {
// 		return models.Album{}, nil
// 	}
// 	return album, nil
// }

// func (s *service) DeleteImage(albumId, imageId string) (models.Image, *errors.RestErr) {
// 	var image models.Image
// 	oId, err := primitive.ObjectIDFromHex(imageId)
// 	if err != nil {
// 		return models.Image{}, errors.NewInternalServerError(err.Error())
// 	}
// 	image.ID = oId
// 	image.AlbumId = albumId
// 	if err := s.repository.DeleteImage(&image); err != nil {
// 		return models.Image{}, err
// 	}
// 	restErr := sendNotification(fmt.Sprintf("image with name %s and id of %s deleted", image.ImageName, image.FileId))
// 	if restErr != nil {
// 		return models.Image{}, nil
// 	}
// 	return image, nil
// }

// func sendNotification(message string) *errors.RestErr {
// 	req := &struct {
// 		Message string `form:"message" json:"message"`
// 	}{Message: message}

// 	log.Printf("Write file to DB was successful. before post File size:")
// 	jsonReq, restErr := json.Marshal(req)
// 	if restErr != nil {
// 		return errors.NewInternalServerError(restErr.Error())
// 	}

// 	_, restErr = http.Post("http://host.docker.internal:9000/sendnotification", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
// 	if restErr != nil {
// 		return errors.NewInternalServerError(restErr.Error())
// 	}
// 	log.Printf("after post")
// 	return nil

// }
