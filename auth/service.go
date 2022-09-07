package auth

import "github.com/dgrijalva/jwt-go"

type Servive interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

var SEKRET_KEY = []byte("CONTOH_SEKRET_KEY")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken, err := token.SignedString(SEKRET_KEY)

	if err != nil {
		return signToken, err
	}

	return signToken, nil
}
