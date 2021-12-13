package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var users = map[string]string{
	"username1": "password1",
	"username2": "password2",
}

var jwtkey = []byte("secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json: "password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	expectedPassword, ok := users[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationtime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationtime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(jwtkey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenstring,
			Expires: expirationtime,
		})

}
func Home(w http.ResponseWriter, r *http.Request) {

	cookie, err:= r.Cookie("token")

	if err != nil{
		{
		if err ==http.ErrNoCookie{
			w.WriteHeader(http.StatusUnauthorized)
}}
	}
	tokenstring := cookie.Value
	
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenstring, claims,func(t *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})

	if !tkn.Valid{
		w.WriteHeader(http.StatusUnauthorized)
	}

}
