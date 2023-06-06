package userservice

import (
	// "backend/entities"
	"backend/entities"
	"backend/helper"
	"backend/services"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

const ACCESS_TOKEN_LIFE = 1 * time.Hour
const REFRESH_TOKEN_LIFE = 24 * time.Hour
const RANDOM_TOKEN_STRING_SIZE = 8
var ARGON2_MEMORY uint32 = 64 * 1024
var ARGON2_TIME uint32 = 3
var ARGON2_THREADS uint8 = 2
var ARGON2_KEYLENGTH uint32 = 32

// func NewEntitiesUser(userInput *services.UserFromFrontend) *entities.User {
// 	user := &entities.User{
// 		Email: userInput.Email,
// 		Username:  userInput.Username,
// 		Password:  userInput.Password,
// 		Salt:      userInput.Salt,

// 	}
// 	return user
// }

func generateRandomSalt(saltSize uint8) ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		return nil, err
	}

	return salt, nil
}

func hashPassword(password string, salt []byte) (string, error) {
	hash := argon2.IDKey([]byte(password), salt, ARGON2_TIME, ARGON2_MEMORY, ARGON2_THREADS, ARGON2_KEYLENGTH)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, ARGON2_MEMORY, ARGON2_TIME, ARGON2_THREADS, b64Salt, b64Hash)

	return encodedHash, nil
}

func verifyPassword(plain, hash string) (bool, error) {
	hashParts := strings.Split(hash, "$")
	var argon2Const struct {
		memory uint32
		time uint32
		threads uint8
	} 
	_, err := fmt.Sscanf(hashParts[3], "m=%d,t=%d,p=%d", &argon2Const.memory, &argon2Const.time, &argon2Const.threads)
	if err != nil {
			return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(hashParts[4])
	if err != nil {
			return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(hashParts[5])
	if err != nil {
			return false, err
	}

	hashToCompare := argon2.IDKey([]byte(plain), salt, argon2Const.time, argon2Const.memory, argon2Const.threads, uint32(len(decodedHash)))

	return subtle.ConstantTimeCompare(decodedHash, hashToCompare) == 1, nil
}

func validateRegisterUser(user *services.RegisterUserInput) (error) {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}

	match := checkRegexp(user.Password)
	if !match {
		return fmt.Errorf("password must not contain special character")
	}

	match = checkRegexp(user.Username)
	if !match {
		return fmt.Errorf("username must not contain special character")
	}

	err = validEmail(user.Email)
	if err != nil{
		log.Error(err)
		return fmt.Errorf("email address is incorrect format")
	}

	if user.Password != user.PasswordConfirmation {
		return fmt.Errorf("password confirm not match")
	}

	return nil
}

func validateUserLoginWithEmail(user *services.LoginUserInput) error {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}
	match := checkRegexp(user.Password)
	if !match {
		return fmt.Errorf("password must not contain special character")
	}

	err = validEmail(user.Username)
	if err != nil{
		log.Error(err)
		return fmt.Errorf("email address is incorrect format")
	}
	return nil
}

func validateUserLoginWithUsername(user *services.LoginUserInput) error {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}
	match := checkRegexp(user.Username)
	if !match {
		return fmt.Errorf("username must not contain special character")
	}

	match = checkRegexp(user.Password)
	if !match {
		return fmt.Errorf("password must not contain special character")
	}

	return nil
}


func validEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func checkRegexp(checkedString string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", checkedString)
	return match
}

func checkIsUserLoginWithEmailCorrect(userLoginInput *services.LoginUserInput, userFromDB entities.User) error {
	if userLoginInput.Username != userFromDB.Email {
		return fmt.Errorf("username is incorrect")
	}
	
	isPasswordCorrect, err := verifyPassword(userLoginInput.Password, userFromDB.Password)
	if err != nil {
		return fmt.Errorf("input invalid")
	}

	if !isPasswordCorrect {
		return fmt.Errorf("password is incorrect")
	}
	return nil
}

func checkIsUserLoginWithUsernameCorrect(userLoginInput *services.LoginUserInput, userFromDB entities.User) error {
	if userLoginInput.Username != userFromDB.Username {
		return fmt.Errorf("username is incorrect")
	}
	
	isPasswordCorrect, err := verifyPassword(userLoginInput.Password, userFromDB.Password)
	if err != nil {
		return fmt.Errorf("input invalid")
	}
	
	if !isPasswordCorrect {
		return fmt.Errorf("password is incorrect")
	}
	return nil
}

func generateToken(username, role string) (accessToken string, refreshToken string, err error) {
	accessToken, err = GenerateAccessToken(username, role)
	if err != nil {
		return accessToken, "", err
	}

	refreshToken, err = GenerateRefreshToken(username, role)
	if err != nil {
		return accessToken, refreshToken, err
	}

	return accessToken, refreshToken, nil 
}

func GenerateAccessToken(username, role string) (string, error) {
	expirationTime := time.Now().Add(ACCESS_TOKEN_LIFE)

	claims := &entities.JWTClaim{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(helper.TOKEN_SERECT_KEY)
}

func GenerateRefreshToken(username, role string) (string, error) {
	randomString, err := generateRandomTokenString()
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(REFRESH_TOKEN_LIFE)
	claims := &entities.JWTClaim{
		Username:     username,
		Role:         role,
		RandomString: randomString,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(helper.TOKEN_SERECT_KEY)
}

func generateRandomTokenString() ([]byte, error) {
	var randomString = make([]byte, RANDOM_TOKEN_STRING_SIZE)

	_, err := rand.Read(randomString[:])

	if err != nil {
		return nil, err
	}

	return randomString, nil
}

func createUserFromGoogleOAuth(googleUser *services.GoogleUserResult) (*entities.User, error){
	randomByte, err := generateRandomTokenString()
	password := string(randomByte[:])
	if err != nil {
		password = "defaul"
	}

	salt, err := generateRandomSalt(SALT_SIZE)
	if err != nil {
		return nil, fmt.Errorf("error when generate salt")
	}

	hashedPassword, err := hashPassword(password, salt)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Username: strings.ToLower(googleUser.Email),
		Email: strings.ToLower(googleUser.Email),
		Password: hashedPassword,
		RoleName: DEFAULT_ROLE,
		Salt:     base64.RawStdEncoding.EncodeToString(salt),
	}
	return user, nil
}