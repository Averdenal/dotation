package Controller

import (
	"encoding/json"
	"fmt"
	"github.com/Averdenal/Dotation/Models"
	"github.com/Averdenal/Dotation/logic"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user Models.User

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(400, err)
		return
	}
	user.BcryptPassWord()
	errSaver := user.Saver()
	if errSaver != nil {
		c.JSON(503, errSaver)
		return
	}
	c.JSON(202, user)
}
func Login(c *gin.Context) {
	var creds Models.Credentials
	err := json.NewDecoder(c.Request.Body).Decode(&creds)
	fmt.Println(creds)
	if err != nil {
		c.JSON(403, gin.H{
			"err": "err creds",
		})
		return
	}
	u, errPass := logic.VerifPassword(&creds)
	if errPass != nil {
		c.JSON(401, gin.H{
			"err": "err verifpassword",
		})
		return
	}
	token, errToken, _ := u.CreatedToken(creds)
	if errToken != nil {
		c.JSON(503, gin.H{
			"err": "error token",
		})
		return
	}

	c.JSON(202, gin.H{
		"user":  u,
		"token": token,
	})
	u = Models.User{}
	return
}
func Welcome(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			c.JSON(401, err)
			return
		}
		// For any other type of error, return a bad request status
		c.JSON(400, err)
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
			c.JSON(401, err)
			return
		}
		fmt.Println("signature err")
		c.JSON(400, err)
		return
	}
	if !tkn.Valid {
		fmt.Println("token not valid")

		c.JSON(401, err)
		return
	}
	c.JSON(200, claims.Username)

}
