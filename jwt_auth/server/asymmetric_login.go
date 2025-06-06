package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"test-jwt-auth/constants"
	"test-jwt-auth/entities"
	"test-jwt-auth/utils"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *basicServer) asymmetricLogin(c *gin.Context) {
	ctx := c.Request.Context()
	loginInfo := &entities.Login{}

	err := c.Bind(loginInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	userClaims := entities.UserClaims{
		Name:  "Thong",
		Id:    "12321bei21ghh4jkh12jkrh",
		Email: "thong@email.com",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-asymmetric",
			Subject:   "12321bei21ghh4jkh12jkrh",
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour)},
			NotBefore: &jwt.NumericDate{Time: time.Now()},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ID:        "12321bei21ghh4jkh12jkrh",
		},
	}

	prvFileName := fmt.Sprintf("%s_priv.pem", userClaims.ID)
	publicFileName := fmt.Sprintf("%s_pub.pem", userClaims.ID)

	defer func() {
		_ = os.Remove("./rsa/" + prvFileName)
		_ = os.Remove("./rsa/" + publicFileName)
	}()

	_, _, err = utils.Shellout(fmt.Sprintf("openssl genrsa -out ./rsa/%s 2048", prvFileName))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	_, _, err = utils.Shellout(fmt.Sprintf("openssl rsa -pubout -in ./rsa/%s -out ./rsa/%s", prvFileName, publicFileName))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	// Load file
	prvKeyPEM, err := os.ReadFile("./rsa/" + prvFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	pubKeyPEM, err := os.ReadFile("./rsa/" + publicFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	prvKey, err := jwt.ParseRSAPrivateKeyFromPEM(prvKeyPEM)
	if err != nil {
		slog.Error("ParseRSAPrivateKeyFromPEM error", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	otps := []jwt.TokenOption{}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, userClaims, otps...)

	signedToken, err := token.SignedString(prvKey)
	if err != nil {
		slog.Error("SignedString error", slog.Any("err", err), slog.String("signedToken", signedToken))
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// DB
	filter := bson.M{
		"user_id": userClaims.ID,
	}

	update := bson.M{
		"$set": bson.M{
			"public_key": string(pubKeyPEM),
		},
	}

	upsert := true

	updateOtps := &options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err = s.db.Collection(constants.TokenPublicKeysColl).UpdateOne(ctx, filter, update, updateOtps)
	if err != nil {
		slog.Error("UpdateOne error", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": signedToken,
	})
}
