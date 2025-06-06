package entities

import jwt "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	RefreshToken bool
	jwt.RegisteredClaims
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserPublicKey struct {
	UserID    string `bson:"user_id"`
	PublicKey string `bson:"public_key"`
}
