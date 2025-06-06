package server

import (
	"log/slog"
	"net/http"
	"test-jwt-auth/constants"
	"test-jwt-auth/entities"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func (s *basicServer) validateSymmetricToken(c *gin.Context) {
	headerToken := c.GetHeader("Token")
	if headerToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	tokenData := &entities.UserClaims{}

	token, err := jwt.ParseWithClaims(headerToken, tokenData, func(t *jwt.Token) (interface{}, error) {
		return []byte(constants.HSKey), nil
	})
	if err != nil {
		slog.Error("ParseWithClaims error", slog.Any("err", err), slog.String("headerToken", headerToken))
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	} else if userClaims, ok := token.Claims.(*entities.UserClaims); ok {
		c.JSON(http.StatusOK, userClaims)
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
}
