package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("sdfghuytre236hgfertoijh012345r0")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
