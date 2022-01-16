package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	//Jika nilai balik lebih dari 1
	//GenerateToken(userID int, userName string) (string, error)
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

//membuat service (jwtService) dapat dikenali oleh service lain
func NewService() *jwtService {
	return &jwtService{}
}

//hanya contoh keep learn
var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")

//Bisa seperti ini juga loh
//func (s *jwtService) GenerateToken(userID int, userName string) (string, error) {
func (s *jwtService) GenerateToken(userID int) (string, error) {
	//data yang disisipkan userID dan nilai yang diterima parameter
	claim := jwt.MapClaims{}
	//Payload data di json { "user_id" : 1 }
	claim["user_id"] = userID
	//untuk data lebih dari 1
	//claim["user_name"] = userName

	//membuat token dengan HASH 256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//Generate token dengan tanda tangan secret key, yang saya sendiri masih bingung wkwkwkw
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	// parameter encoded token akan diterjemahkan di jwt.Parse, kemudian divalidasi oleh function dengan nilai balik interface dan errot
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		//token akan dicek apakah sudah menggunakan metode yang sama atau tidak
		//method hmac sama dengan hs256
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		//kemudian akan dikembalikan dengan encoded secret key kedalam function kosong
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
