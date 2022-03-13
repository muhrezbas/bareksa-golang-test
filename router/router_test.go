package router

import (
	// "errors"
	// "fmt"
	// "github.com/stretchr/testify/assert"
	// "net/http"
	// "net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	// r := LoadRoutes()
	// r.Get("/users").Name("users")
	// assert.NotNil(t, r.Route("users"))
	// assert.Nil(t, r.Route("users2"))

	// routers := gin.Default()
	// routers.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "UPDATE"},
	// 	AllowHeaders:     []string{"Access-Control-Allow-Origin", "Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	// routers.Use(gzip.Gzip(gzip.DefaultCompression))

	// router := LoadRoutes

	// w := httptest.NewRecorder()
	// req, _ := http.NewRequest("GET", "/ping", nil)
	// fmt.Println("sebelum router", req, w)
	// routers.ServeHTTP(w, req)
	// fmt.Println(routers)

	// assert.Equal(t, 200, w.Code)
}
