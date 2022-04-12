package reqid

import (
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
)

//CorrelationIDMiddleware adds correlationID if it's not specified in HTTP request
func CorrelationIDMiddleware() gin.HandlerFunc {
	return addCorrelationID
}

func GetCorralationID(c *gin.Context) string {
	corralationID := c.Request.Header.Get("CorrelationID")
	return corralationID
}

func addCorrelationID(c *gin.Context) {
	corralationID := GetCorralationID(c)

	if strings.TrimSpace(corralationID) == "" {
		id, _ := uuid.NewV4()
		corralationID = id.String()
		c.Request.Header.Add("CorrelationID", corralationID)
	}
	entry := log.WithFields(log.Fields{
		"correlationID": corralationID,
		"Method":        c.Request.Method,
		"Path":          c.FullPath(),
		"Handler":       c.HandlerName(),
	})
	c.Set("Logger", entry)
	c.Next()
}
