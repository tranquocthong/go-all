package server

import (
	"log/slog"
	"net/http"
	"test-jwt-auth/utils"

	"github.com/gin-gonic/gin"
)

func (s *basicServer) symmetricLogin(c *gin.Context) {
	// Validate password
	userID := ""

	user := getUserById(userID)

	accessToken, rf, err := utils.GetJwtTokenPair(user, userID)
	if err != nil {
		slog.Error("SignedString rf error", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": rf,
	})
}
