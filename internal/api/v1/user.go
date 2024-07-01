package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go.opentelemetry.io/otel"
)

type Users struct {
}

func NewUsers() *Users {
	return &Users{}
}

func (u *Users) Path() string {
	return "/users"
}

func (u *Users) Find(c *gin.Context) {
	// Start a new child span
	_, span := otel.Tracer("api.find").Start(c.Request.Context(), "Find")
	defer span.End()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
