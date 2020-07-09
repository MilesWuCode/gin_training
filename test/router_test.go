package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"mygo/database"
	"mygo/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("../")
	viper.ReadInConfig()

	database.Init()

	r := router.Router()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
