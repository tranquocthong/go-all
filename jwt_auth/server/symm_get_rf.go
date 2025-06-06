package server

import (
	"log/slog"
	"net/http"
	"test-jwt-auth/constants"
	"test-jwt-auth/entities"
	"test-jwt-auth/utils"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func (s *basicServer) symmetricGetRf(c *gin.Context) {
	rf := c.Request.Header.Get("Token")

	t := entities.RefreshTokenClaims{}

	token, err := jwt.ParseWithClaims(rf, &t, func(t *jwt.Token) (interface{}, error) {
		return []byte(constants.HSKey), nil
	})
	if err != nil {
		slog.Error("ParseWithClaims error", slog.Any("err", err), slog.String("rf", rf))
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	} else if claims, ok := token.Claims.(*entities.RefreshTokenClaims); ok && claims.RefreshToken {
		user := getUserById(claims.ID)

		accessToken, rf, err := utils.GetJwtTokenPair(user, claims.ID)
		if err != nil {
			slog.Error("utils.GetJwtTokenPair error", slog.Any("err", err), slog.String("claims.ID", claims.ID))
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken":  accessToken,
			"refreshToken": rf,
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errMsg": "invalid type of token",
		})
		return
	}
}

func getUserById(_ string) *entities.UserClaims {
	return &entities.UserClaims{
		Name:  "Thong",
		Id:    "12321bei21ghh4jkh12jkrh",
		Email: "thong@email.com",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-symmetric",
			Subject:   "12321bei21ghh4jkh12jkrh",
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour)},
			NotBefore: &jwt.NumericDate{Time: time.Now()},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ID:        "12321bei21ghh4jkh12jkrh",
		},
	}
}
