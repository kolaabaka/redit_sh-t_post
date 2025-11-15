package controller

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"goSiteProject/internal/app"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func TestA(t *testing.T) {
	//todo: use .env files
	godotenv.Load()
	fmt.Println(os.Getenv("TEMPLATE_PATH"))
}

// current test doesn`t work without makeFile and set TEMPLATE_PATH, doesn`t bother me =*)
func TestLoginPage(t *testing.T) {
	router := gin.Default()
	app.SetupRoutes(router)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/login", nil))

	if recorder.Code != 200 {
		t.Errorf("Login page should return code 200, but code is %d", recorder.Code)
	}
}
