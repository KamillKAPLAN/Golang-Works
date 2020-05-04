package libs

import (
	"github.com/dgrijalva/jwt-go"
	dbconfig "github.com/PROJECTS/user_application_go/dbConfig"
	"github.com/PROJECTS/user_application_go/models"
	"time"
)

func CreateJWT(user models.Users) (string, error) {
	// create the token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	// authorized = yetkili, izinli
	claims["id"] = user.ID
	claims["email"] = user.Email

	// verilen değere göre token değişecek
	claims["user"] = user.FirstName + " " + user.LastName

	// saat * 24 * ay güün olarak = 2 ay kullanılabilir bir token oluşturduk.

	claims["exp"] = time.Now().Add(dbconfig.TokenExpireDuration).Unix()
	claims["expd"] = time.Now().Add(dbconfig.TokenExpireDuration).Format("2006-01-02T15:04:05.000Z")
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(dbconfig.MySecretKey)

	return tokenString, err
}

// Parse with Token
func ParseToken(myToken string, ClaimsName string) interface{} {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(dbconfig.MySecretKey), nil
	})

	if err == nil && token.Valid {
		return token.Claims.(jwt.MapClaims)[ClaimsName]
	} else {
		return ""
	}
}
