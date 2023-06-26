package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(userInput *entities.User) (*entities.User, error)
	Get(username string) (u *entities.User, err error)
	List() (user *[]entities.User, err error)
	Delete(username string) error
	Update(userInput *entities.User) error
	GetWithEmail(email string) (*entities.User, error)
	FindOrCreateWithEmail(*entities.User) (*entities.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) Create(userInput *entities.User) (*entities.User, error) {
	err := repo.DB.Create(userInput).Error
	if err != nil {
		return nil, err
	}
	return userInput, nil
}

func (repo *UserRepo) Get(username string) (u *entities.User, err error) {
	return getUser(username, repo)
}

func getUser(username string, repo *UserRepo) (u *entities.User, err error) {
	user := new(entities.User)
	err = repo.DB.Select("role_name", "email", "username", "password", "salt").Where(map[string]interface{}{"username": username}).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) GetWithEmail(email string) (*entities.User, error) {
	admin := new(entities.User)
	err := repo.DB.Where(map[string]interface{}{"email": email}).Find(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo *UserRepo) List() (user *[]entities.User, err error) {
	users := make([]entities.User, 10)
	err = repo.DB.Select("role_name", "email", "username").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (repo *UserRepo) Delete(username string) error {
	user, err := getUser(username, repo)
	if err != nil {
		return err
	}
	err = repo.DB.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) Update(userInput *entities.User) error {
	// err := repo.DB.Model(&userInput).Where("username = ?", userInput.Username).Updates(map[string]interface{}{"full_name": userInput.FullName, "role": userInput.Role}).Error
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (repo *UserRepo) FindOrCreateWithEmail(user *entities.User) (*entities.User, error) {
	err := repo.DB.
		Where(entities.User{Username: user.Username}).
		FirstOrCreate(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
