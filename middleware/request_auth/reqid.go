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

// // ContextKey is
// //type ContextKey string

// const (
// 	// ContextKeyReqID is the context key for RequestID
// 	// ContextKeyReqID ContextKey = "requestID"
// 	ContextKeyReqID = "requestID"

// 	// HTTPHeaderNameRequestID has the name of the header for request ID
// 	HTTPHeaderNameRequestID = "X-Request-ID"
// )

// // GetReqID will get reqID from a http request and return it as a string
// // func GetReqID(ctx context.Context) string {

// // 	reqID := ctx.Value(ContextKeyReqID)

// // 	if ret, ok := reqID.(string); ok {
// // 		return ret
// // 	}

// // 	return ""
// // }

// // // AttachReqID will attach a brand new request ID to a http request
// // func AttachReqID(ctx context.Context) context.Context {
// // 	var reqID string
// // 	if GetReqID(ctx) == "" {
// // 		reqID = uuid.New().String()
// // 	} else {
// // 		reqID = GetReqID(ctx)
// // 	}
// // 	return context.WithValue(ctx, ContextKeyReqID, reqID)
// // }

// //Middleware will attach the reqID to the http.Request and add reqID to http header in the response
// // func Middleware(next http.Handler) http.Handler {
// // 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// // 		ctx := AttachReqID(r.Context())
// // 		r = r.WithContext(ctx)
// // 		next.ServeHTTP(w, r)
// // 		h := w.Header()
// // 		h.Add(HTTPHeaderNameRequestID, GetReqID(ctx))
// // 	})
// // }

// func GetReqID(ctx *gin.Context) string {
// 	ret := ctx.GetString(ContextKeyReqID)
// 	return ret

// }

// func AttachReqID(ctx *gin.Context) {
// 	var reqID string
// 	if GetReqID(ctx) == "" {
// 		reqID = uuid.New().String()
// 	} else {
// 		reqID = GetReqID(ctx)
// 	}
// 	ctx.Set(ContextKeyReqID, reqID)

// }

// func Middleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		AttachReqID(c)
// 		fmt.Println("added req id : ", GetReqID(c))
// 		c.Next()
// 		return

// 	}
// }
