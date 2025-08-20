package repository

import (
	"split-bill/backend/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Db: db,
	}
}

// Delete implements UserRepository.
func (r *UserRepositoryImpl) Delete(id string) error {
	var user model.User
	if err := r.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	return r.Db.Delete(&user).Error
}

// FindAll implements UserRepository.
func (r *UserRepositoryImpl) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FindByUsername implements UserRepository.
func (r *UserRepositoryImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID implements UserRepository.
func (r *UserRepositoryImpl) FindByID(id string) (*model.User, error) {
	var user model.User
	if err := r.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	if user == (model.User{}) {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

// Create implements UserRepository.
func (r *UserRepositoryImpl) Create(user *model.User) (*model.User, error) {
	// Generate a new UUID
	user.ID = uuid.New()

	// Set the created time and updated time
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()
	return user, r.Db.Create(user).Error
}

// Update implements UserRepository.
func (r *UserRepositoryImpl) Update(user *model.User) (*model.User, error) {
	// Set the updated time
	user.UpdatedAt = time.Now()
	if err := r.Db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
