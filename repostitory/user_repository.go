package repostitory

import (
	"bankSystem/model"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(dataBase *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: dataBase,
	}
}

func (r *PostgresUserRepository) GetUser(login string) (*model.UserEntity, error) {
	var user model.UserEntity
	if err := r.db.Where("login = ?", login).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) SaveUser(user *model.UserEntity) error {
	return r.db.Save(&user).Error
}

func (r *PostgresUserRepository) DeleteUser(user *model.UserEntity) error {
	return r.db.Delete(&user).Error
}
