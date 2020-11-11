package Models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (r *User) FindByEmail(email string) error {
	err := database.Where("email = ?", email).First(r).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *User) FindById(id string) error {
	err := database.First(r, id).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *Users) FindAllUser() error {
	err := database.Find(&u.Users).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *User) FindByName(n string) error {
	database.Where("name = ?", n).First(&r)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *User) Saver() error {
	database.Save(r)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *User) Delete() error {
	database.Delete(&r)
	err := database.Error
	if err != nil {
		return err
	}
	return nil
}

var JwtKey = []byte("my_secret_key")

func (u *User) BcryptPassWord() {
	pwd, _ := bcrypt.GenerateFromPassword([]byte(u.Pwd), 14)
	u.Pwd = string(pwd)
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Pwd), []byte(password))
	return err != nil
}
func (u *User) CreatedToken(creds Credentials) (string, error, time.Time) {

	expirationTime := time.Now().Add(60 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		IdUser:   u.ID,
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err, expirationTime
	}
	return tokenString, nil, expirationTime
}
