package logic

import (
	"errors"
	"fmt"
	"github.com/Averdenal/Dotation/Models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifPassword(creds *Models.Credentials) (Models.User, error) {
	var user Models.User
	err := user.FindByEmail(creds.Username)
	fmt.Println(user)
	if err != nil {
		fmt.Printf("erreur %v :", err)
		return Models.User{}, err
	}
	if user.CheckPasswordHash(creds.Password) {
		fmt.Println("check password")
		return Models.User{}, errors.New("Password not ok")
	}
	return user, nil
}

func Access() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				c.AbortWithStatusJSON(401, err)
				return
			}
			// For any other type of error, return a bad request status
			c.AbortWithStatusJSON(400, err)
			return
		}
		tknStr := cookie.Value
		claims := &Models.Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return Models.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Println("signature err")
				c.AbortWithStatusJSON(401, err)
				return
			}
			fmt.Println("signature err")
			c.AbortWithStatusJSON(400, err)
			return
		}
		if !tkn.Valid {
			fmt.Println("token not valid")

			c.AbortWithStatusJSON(401, err)
			return
		}
		c.Set("user", claims)
		c.Next()

	}

}
