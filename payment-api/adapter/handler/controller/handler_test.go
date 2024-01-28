package controller_test

import (
	"net/http/httptest"
	"testing"

	handler "github.com/fiap-postech-soat1-group21-stage4/payment-api/payment-api/adapter/handler/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {
	h := handler.NewHandler()
	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	h.RegisterRoutes(engine.Group("/api/v1"))

	routesInfo := engine.Routes()
	routesMethodAndPath := make([][]string, 0, len(routesInfo))
	for _, routeInfo := range routesInfo {
		routesMethodAndPath = append(routesMethodAndPath, []string{routeInfo.Method, routeInfo.Path})
	}

	expectedRoutesMethodAndPath := [][]string{
		{"POST", "/api/v1/payments/:id"},
	}

	assert.Equal(t, expectedRoutesMethodAndPath, routesMethodAndPath)
}
