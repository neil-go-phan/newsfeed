package repository

import (
	"server/entities"
	"server/helpers"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(userInput *entities.User) (*entities.User, error)
	Get(username string) (u *entities.User, err error)

	List(page int, pageSize int) ([]entities.User, error)
	Delete(id uint) error
	ChangeRole(id uint, role string) error
	UserUpgrateRole(username string) error
	Count() (int, error)

	Update(userInput *entities.User) error
	GetWithEmail(email string) (*entities.User, error)
	FindOrCreateWithEmail(*entities.User) (*entities.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

const PREMIUM_TIER = "Premium tier user"

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

func (repo *UserRepo) List(page int, pageSize int) ([]entities.User, error) {
	users := make([]entities.User, 10)

	err := repo.DB.Scopes(helpers.Paginate(page, pageSize)).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepo) ChangeRole(id uint, role string) error {
	err := repo.DB.
		Model(&entities.User{}).
		Where("id = ?", id).
		Update("role_name", role).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepo) UserUpgrateRole(username string) error {
	err := repo.DB.
		Model(&entities.User{}).
		Where("username = ?", username).
		Update("role_name", PREMIUM_TIER).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepo) Count() (int, error) {
	var count int64
	err := repo.DB.Table("users").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
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

func (repo *UserRepo) Delete(id uint) error {
	err := repo.DB.Unscoped().Delete(&entities.User{}, id).Error
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
