package app

import (
	"github.com/cab_booking_app/middleware/cors_middleware"
	reqid "github.com/cab_booking_app/middleware/request_auth"

	"github.com/cab_booking_app/configuration"
	dbconnection "github.com/cab_booking_app/datasource/mysql"
	"github.com/cab_booking_app/handler"
	"github.com/cab_booking_app/repository"
	"github.com/cab_booking_app/service"

	"github.com/gin-gonic/gin"
)

func StartApplication() {

	router := gin.New()

	router.Use(gin.Recovery(), gin.Logger(), reqid.CorrelationIDMiddleware(), cors_middleware.CORSMiddleware())

	dbConn := dbconnection.NewDataBaseConnection().GetDatabaseConnection(configuration.ViperEnvVariable("DB_USER"),
		configuration.ViperEnvVariable("DB_PASS"), configuration.ViperEnvVariable("DB_HOST"), configuration.ViperEnvVariable("DATABASE"))

	repo := repository.NewCabRepo(dbConn)
	srv := service.NewService(repo)
	handler := handler.NewHandler(srv)
	router.POST("/bookCab", handler.BookCab)
	router.GET("/searchCab", handler.SearchCab)
	router.GET("/getBookingHistory", handler.GetBookingHistory)
	router.Run(":8000")
}
