package Models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (r *User) FindByEmail(email string) error {
	return database.Where("email = ?", email).First(r).Error
}

func (r *User) FindById(id string) error {
	return database.First(r, id).Error
}
func (u *Users) FindAllUser() error {
	return database.Find(&u.Users).Error
}
func (r *User) FindByName(n string) error {
	return database.Where("name = ?", n).First(&r).Error
}

func (r *User) Saver() error { return database.Save(r).Error }

func (r *User) Delete() error { return database.Delete(&r).Error }

var JwtKey = []byte("my_secret_key")

func (r *User) BcryptPassWord() {
	pwd, _ := bcrypt.GenerateFromPassword([]byte(r.Pwd), 14)
	r.Pwd = string(pwd)
}

func (r *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(r.Pwd), []byte(password))
	return err != nil
}
func (r *User) CreatedToken(creds Credentials) (string, error, time.Time) {

	expirationTime := time.Now().Add(60 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		IdUser:   r.ID,
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
