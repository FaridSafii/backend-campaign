package auth

import "github.com/golang-jwt/jwt"

type Service interface {
	GenerateToken(userID int, userName string) (string, error)
}

type jwtService struct {
}

//membuat service (jwtService) dapat dikenali oleh service lain
func NewService() *jwtService {
	return &jwtService{}
}

//hanya contoh keep learn
var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")

func (s *jwtService) GenerateToken(userID int, userName string) (string, error) {
	//data yang disisipkan userID dan nilai yang diterima parameter
	claim := jwt.MapClaims{}
	//Payload data di json { "user_id" : 1 }
	claim["user_id"] = userID
	claim["user_name"] = userName

	//membuat token dengan HASH 256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//Generate token dengan tanda tangan secret key, yang saya sendiri masih bingung wkwkwkw
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
