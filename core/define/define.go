package define

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"
var MailPassword = os.Getenv("MailPassword")
