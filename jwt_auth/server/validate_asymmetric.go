package server

import (
	"log/slog"
	"net/http"
	"test-jwt-auth/constants"
	"test-jwt-auth/entities"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *basicServer) validateAsymmetricToken(c *gin.Context) {
	ctx := c.Request.Context()
	headerToken := c.GetHeader("Token")
	if headerToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	tokenData := &entities.UserClaims{}
	userID := c.Request.URL.Query().Get("userId")

	// DB
	userPublicKeyM := &entities.UserPublicKey{}

	rs := s.db.Collection(constants.TokenPublicKeysColl).FindOne(ctx, bson.M{"user_id": userID})

	if err := rs.Decode(userPublicKeyM); err != nil {
		slog.Error("Decode error", slog.Any("err", err), slog.Any("userPublicKeyM", userPublicKeyM))
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(userPublicKeyM.PublicKey))
	if err != nil {
		slog.Error("ParseRSAPublicKeyFromPEM error", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	token, err := jwt.ParseWithClaims(headerToken, tokenData, func(t *jwt.Token) (interface{}, error) {
		return pubKey, nil
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
