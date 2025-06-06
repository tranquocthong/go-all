package utils

import (
	"bytes"
	"os/exec"
	"test-jwt-auth/constants"
	"test-jwt-auth/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(constants.ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func GetJwtTokenPair(data jwt.Claims, id string) (string, string, error) {

	otps := []jwt.TokenOption{}

	jwtTk := jwt.NewWithClaims(jwt.SigningMethodHS256, data, otps...)

	accessToken, err := jwtTk.SignedString([]byte(constants.HSKey))
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := entities.RefreshTokenClaims{
		RefreshToken: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)},
			ID:        id,
		},
	}

	jwtRf := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims, otps...)

	rf, err := jwtRf.SignedString([]byte(constants.HSKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, rf, nil
}
