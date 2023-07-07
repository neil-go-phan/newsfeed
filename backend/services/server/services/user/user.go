package userservice

import (
	"encoding/base64"
	"server/entities"
	"server/repository"
	"server/services"
	"strings"

	"fmt"

	log "github.com/sirupsen/logrus"
)

const SALT_SIZE = 8 // 8 byte
const DEFAULT_ROLE = "Free tier user"
const PREMIUM_TIER = "Premium tier user"
const GORM_DUPLICATE_USERNAME_ERR_CONTAIN = "users_username_key"
const GORM_DUPLICATE_EMAIL_ERR_CONTAIN = "users_email_key"

const USER_ROLE_ENTITY = "USER"
const USER_ROLE_DELETE_METHOD = "DELETE"

const ADMIN_ROLE_ENTITY = "ADMIN PAGE"
const ADMIN_ROLE_ACCESS_PAGE_METHOD = "ACCESS"



type UserService struct {
	repo        repository.UserRepository
	roleService services.RoleServices
}

func NewUserService(repo repository.UserRepository, roleService services.RoleServices) *UserService {
	return &UserService{
		repo:        repo,
		roleService: roleService,
	}
}

func (s *UserService) AccessAdminPage(role string) error {
	isAllowed := s.roleService.GrantPermission(role, ADMIN_ROLE_ENTITY, ADMIN_ROLE_ACCESS_PAGE_METHOD)
	if !isAllowed {
		return fmt.Errorf("unauthorized")
	}
	return nil
}

func (s *UserService) GetUser(username string) (*entities.User, error) {
	return s.repo.Get(username)
}

func (s *UserService) List(page int, pageSize int) ([]services.UserResponse, error) {
	userResponse := make([]services.UserResponse, 0)
	users, err := s.repo.List(page, pageSize)
	if err != nil {
		return userResponse, err
	}
	for _, user := range users {
		userResponse = append(userResponse, castUserToResponse(user))
	}
	return userResponse, nil
}

func (s *UserService) Delete(role string, id uint) error {
	isAllowed := s.roleService.GrantPermission(role, USER_ROLE_ENTITY, USER_ROLE_DELETE_METHOD)
	if !isAllowed {
		return fmt.Errorf("unauthorized")
	}

	return s.repo.Delete(id)
}

func (s *UserService) Count() (int, error) {
	return s.repo.Count()
}

func (s *UserService) ChangeRole(role string, id uint, newRole string) error {
	return s.repo.ChangeRole(id, newRole)
}

func (s *UserService) UserUpgrateRole(role string, username string) (string, string, error) {
	if role == PREMIUM_TIER {
		return "", "", fmt.Errorf("you already a premium user")
	}
	err := s.repo.UserUpgrateRole(username)
	if err != nil {
		log.Errorf("error occrus when a user try to upgrate: %s\n", err)
		return "", "", fmt.Errorf("internal server error")
	}

	accessToken, refreshToken, err := generateToken(username, PREMIUM_TIER)
	if err != nil {
		log.Errorf("error occrus when a user try to upgrate: %s\n", err)
		return "", "", fmt.Errorf("internal server error")
	}
	return accessToken, refreshToken, nil
}

func (s *UserService) CreateUser(registerUserInput *services.RegisterUserInput) (*entities.User, error) {
	err := validateRegisterUser(registerUserInput)
	if err != nil {
		return nil, err
	}

	salt, err := generateRandomSalt(SALT_SIZE)
	if err != nil {
		return nil, fmt.Errorf("error when generate salt")
	}

	hashedPassword, err := hashPassword(registerUserInput.Password, salt)
	if err != nil {
		return nil, err
	}

	entitiesUser := &entities.User{
		Username: registerUserInput.Username,
		Password: hashedPassword,
		Email:    registerUserInput.Email,
		RoleName: DEFAULT_ROLE,
		Salt:     base64.RawStdEncoding.EncodeToString(salt),
	}

	user, err := s.repo.Create(entitiesUser)
	if err != nil {
		if strings.Contains(err.Error(), GORM_DUPLICATE_USERNAME_ERR_CONTAIN) {
			return nil, fmt.Errorf("username already exist")
		}
		if strings.Contains(err.Error(), GORM_DUPLICATE_EMAIL_ERR_CONTAIN) {
			return nil, fmt.Errorf("email already exist")
		}
	}

	return user, nil
}

func (s *UserService) LoginWithEmail(inputUser *services.LoginUserInput) (accessToken string, refreshToken string, err error) {
	err = validateUserLoginWithEmail(inputUser)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login: %s\n", err)
		return "", "", fmt.Errorf("input invalid")
	}

	userFromDB, err := s.repo.GetWithEmail(inputUser.Username)
	if err != nil {
		return "", "", err
	}

	err = checkIsUserLoginWithEmailCorrect(inputUser, *userFromDB)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", "", fmt.Errorf("username or password is incorrect")
	}

	accessToken, refreshToken, err = generateToken(userFromDB.Username, userFromDB.RoleName)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", "", fmt.Errorf("internal server error")
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) LoginWithUsername(inputUser *services.LoginUserInput) (accessToken string, refreshToken string, err error) {
	err = validateUserLoginWithUsername(inputUser)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login: %s\n", err)
		return "", "", fmt.Errorf("input invalid")
	}

	userFromDB, err := s.repo.Get(inputUser.Username)
	if err != nil {
		return "", "", err
	}

	err = checkIsUserLoginWithUsernameCorrect(inputUser, *userFromDB)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login: %s\n", err)
		return "", "", fmt.Errorf("username or password is incorrect")
	}

	accessToken, refreshToken, err = generateToken(userFromDB.Username, userFromDB.RoleName)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login: %s\n", err)
		return "", "", fmt.Errorf("internal server error")
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) GoogleOAuth(googleUser *services.GoogleUserResult) (accessToken string, refreshToken string, err error) {
	user, err := createUserFromGoogleOAuth(googleUser)
	if err != nil {
		return "", "", err
	}

	user, err = s.repo.FindOrCreateWithEmail(user)
	if err != nil {
		return "", "", fmt.Errorf("user not found")
	}

	accessToken, refreshToken, err = generateToken(user.Username, user.RoleName)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login: %s\n", err)
		return "", "", fmt.Errorf("internal server error")
	}

	return accessToken, refreshToken, nil
}

// func (s *UserService) VerifyUser(username string, userFromFrontend services.UserFromFrontend) (bool, error){
// 	if username != userFromFrontend.Username {
// 		return false, errors.New("username is incorrect")
// 	}
// 	userFromDB, err := s.GetUser(username)
// 	if err != nil {
// 		return false, err
// 	}
// 	if userFromDB.Username == "" {
// 		return false, errors.New("username is incorrect")
// 	}
// 	return verifyPassword(userFromFrontend.Password, userFromDB.Password)
// }

// func (s *UserService)UpdateUser(userFromFrontend *services.RegisterUserInput) error {
// 	// user := NewEntitiesUser(userFromFrontend)
// 	err := s.roleService.Validate(user.RoleName)
// 	if err != nil {
// 		return err
// 	}
// 	return s.repo.Update(user)
// }
