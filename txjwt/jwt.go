package txjwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/txvier/base/txerror"
	"github.com/txvier/base/txlogger"
	"net/http"
	"time"
)

var (
	logger    = txlogger.GetLogger()
	SecretMap = map[string]string{
		"PROJECT_NAME": "welcome to dmp",
	}

	IgnoreValidateRoute = map[string]bool{
		"/dpm/api/v1.0/users/register": true,
		"/dpm/api/v1.0/users/login":    true,
	}
)

const (
	CURRENT_USER = "CUSR"
	TOKEN_KEY    = "Authorization"
)

type UserClaims struct {
	*jwt.StandardClaims
	TokenType string
	*UserToken
}

type UserToken struct {
	Id     string
	Name   string
	Pwd    string
	Issuer string
}

func CreateToken(user *UserToken) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("HS256"))
	now := time.Now()
	issuedAt := time.Date(2018, time.May, 30, 0, 0, 0, 0, now.Location())
	// set our claims
	t.Claims = &UserClaims{
		&jwt.StandardClaims{
			// set the expire time
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: now.Add(time.Hour * 24).Unix(), //for dev
			// ExpiresAt: now.Add(time.Minute * 10).Unix(),
			Issuer:   user.Issuer,
			IssuedAt: issuedAt.Unix(),
		},
		"level1",
		user,
	}
	// Creat token string
	return t.SignedString([]byte(SecretMap[user.Issuer]))
}

func ValidateTokenHandlerFunc(issuer string) gin.HandlerFunc {
	return func(cxt *gin.Context) {
		if !IgnoreValidateRoute[cxt.Request.URL.Path] && cxt.Request.Method != "OPTIONS" {
			// only accessible with a valid token
			// Get token from request
			token, err := request.ParseFromRequestWithClaims(cxt.Request, request.AuthorizationHeaderExtractor, &UserClaims{},
				func(token *jwt.Token) (interface{}, error) {
					return []byte(SecretMap[issuer]), nil
				})
			logger.Info("token:", token)
			// If the token is missing or invalid, return error
			if err != nil {
				error := txerror.ErrTraceCode(http.StatusUnauthorized, err)
				panic(error)
			}

			// Token is valid
			// fmt.Fprintln(w, "Welcome,", token.Claims.(*UserClaims).Name)
			u := token.Claims.(*UserClaims).UserToken
			cxt.Set(CURRENT_USER, u)
			// Got the value like this : context.Get(r,"cusr").(*UserToken)
			logger.Infof("username is:[%s],and pwd is:[%s]", u.Name, "*********")
		}

		cxt.Next()
	}
}
