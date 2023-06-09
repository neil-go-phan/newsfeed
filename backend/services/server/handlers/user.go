package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"server/helpers"
	"server/services"
	userservice "server/services/user"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	service services.UserServices
}

//go:generate mockery --name UserHandlerInterface
type UserHandlerInterface interface {
	CheckAuth(c *gin.Context)
	Token(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	GoogleOAuth(c *gin.Context)
	AccessAdminPage(c *gin.Context)
	ChangeRole(c *gin.Context)
	List(c *gin.Context)
	Delete(c *gin.Context)
	Total(c *gin.Context)
	UserUpgrateRole(c *gin.Context)
}

const GOOGLE_OATH_TOKEN_ROOT_URl = "https://oauth2.googleapis.com/token"

type deleteUserPayload struct {
	ID uint `json:"id"`
}

type changeRoleUserPayload struct {
	ID      uint   `json:"id"`
	NewRole string `json:"new_role"`
}

func NewUserHandler(service services.UserServices) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CheckAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Granted permission"})
}

func (h *UserHandler) AccessAdminPage(c *gin.Context) {
	role, exsit := c.Get("role")
	if !exsit {
		log.Error("Not found role in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	err := h.service.AccessAdminPage(role.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Granted permission"})
}

func (h *UserHandler) List(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}
	users, err := h.service.List(page, pageSize)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "users": users})
}

func (h *UserHandler) Delete(c *gin.Context) {
	var payload deleteUserPayload
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	role, exsit := c.Get("role")
	if !exsit {
		log.Error("Not found role in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	err = h.service.Delete(role.(string), payload.ID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete success"})
}

func (h *UserHandler) ChangeRole(c *gin.Context) {
	var payload changeRoleUserPayload
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	role, exsit := c.Get("role")
	if !exsit {
		log.Error("Not found role in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	err = h.service.ChangeRole(role.(string), payload.ID, payload.NewRole)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete success"})
}

func (h *UserHandler) UserUpgrateRole(c *gin.Context) {
	role, exsit := c.Get("role")
	if !exsit {
		log.Error("Not found role in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	accessToken, refreshToken, err := h.service.UserUpgrateRole(role.(string), username.(string))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "upgrate tier success", "access_token": accessToken, "refresh_token": refreshToken})
}

func (h *UserHandler) Total(c *gin.Context) {

	total, err := h.service.Count()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "total": total})
}

func (h *UserHandler) Token(c *gin.Context) {
	username, _ := c.Get("username")
	role, _ := c.Get("role")
	accessToken, err := userservice.GenerateAccessToken(username.(string), role.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Bad request: fail to generate access token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Successful token reissue", "access_token": accessToken})
}

func (h *UserHandler) Register(c *gin.Context) {
	var inputUser services.RegisterUserInput
	err := c.BindJSON(&inputUser)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	_, err = h.service.CreateUser(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "register success"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var inputUser services.LoginUserInput
	err := c.BindJSON(&inputUser)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "internal server error"})
		return
	}
	var accessToken, refreshToken string
	if !helpers.CheckIsEmail(inputUser.Username) {
		log.Println("login with username")
		accessToken, refreshToken, err = h.service.LoginWithUsername(&inputUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}
	} else {
		log.Println("login with email")
		accessToken, refreshToken, err = h.service.LoginWithEmail(&inputUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "access_token": accessToken, "refresh_token": refreshToken})
}

func (h *UserHandler) GoogleOAuth(c *gin.Context) {
	code := c.Query("code")
	var pathUrl string = "/"

	if c.Query("state") != "" {
		pathUrl = c.Query("state")
	}

	if code == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Authorization code not provided!"})
		return
	}

	tokenRes, err := GetGoogleOauthToken(code)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"success": false, "message": err.Error()})
		return
	}

	google_user, err := GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"success": false, "message": err.Error()})
		return
	}

	config, _ := helpers.LoadEnv(".")

	accessToken, refreshToken, err := h.service.GoogleOAuth(google_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s%s?access_token=%s&refresh_token=%s", config.FrontEndOrigin, pathUrl, accessToken, refreshToken))
}

func GetGoogleOauthToken(code string) (*services.GoogleOauthToken, error) {
	resBody, err := sendRequestGetGoogleOauthToken(code)
	if err != nil {
		return nil, err
	}

	var GoogleOauthTokenRes map[string]interface{}
	if err := json.Unmarshal(resBody.Bytes(), &GoogleOauthTokenRes); err != nil {
		return nil, err
	}

	tokenBody := &services.GoogleOauthToken{
		Access_token: GoogleOauthTokenRes["access_token"].(string),
		Id_token:     GoogleOauthTokenRes["id_token"].(string),
	}

	return tokenBody, nil
}

func sendRequestGetGoogleOauthToken(code string) (bytes.Buffer, error) {
	var resBody bytes.Buffer
	query := createQueryGetGoogleOauthToken(code)

	req, err := http.NewRequest("POST", GOOGLE_OATH_TOKEN_ROOT_URl, bytes.NewBufferString(query))
	if err != nil {
		return resBody, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return resBody, err
	}
	defer res.Body.Close()

	log.Println("res", res)
	log.Println("err", err)

	if res.StatusCode != http.StatusOK {
		return resBody, fmt.Errorf("could not retrieve token")
	}

	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return resBody, err
	}
	return resBody, nil
}

func createQueryGetGoogleOauthToken(code string) string {
	config, err := helpers.LoadEnv(".")
	if err != nil {
		log.Error(err)
	}
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", config.GoogleClientID)
	values.Add("client_secret", config.GoogleClientSecret)
	values.Add("redirect_uri", config.GoogleOAuthRedirectUrl)

	query := values.Encode()
	return query
}

func GetGoogleUser(access_token string, id_token string) (*services.GoogleUserResult, error) {
	resBody, err := sendRequestGetGoogleUser(access_token, id_token)
	if err != nil {
		return nil, err
	}

	var GoogleUserRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		return nil, err
	}

	userBody := &services.GoogleUserResult{
		Id:             GoogleUserRes["id"].(string),
		Email:          GoogleUserRes["email"].(string),
		Verified_email: GoogleUserRes["verified_email"].(bool),
		Name:           GoogleUserRes["name"].(string),
		Given_name:     GoogleUserRes["given_name"].(string),
		Picture:        GoogleUserRes["picture"].(string),
		Locale:         GoogleUserRes["locale"].(string),
	}

	return userBody, nil
}

func sendRequestGetGoogleUser(access_token string, id_token string) (bytes.Buffer, error) {
	rootUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", access_token)
	var resBody bytes.Buffer
	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return resBody, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", id_token))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return resBody, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return resBody, fmt.Errorf("could not retrieve user")
	}

	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return resBody, err
	}
	return resBody, nil
}
