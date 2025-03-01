package user

import "demo/pkg/db"

type UserRepository struct {
	database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) FindByEmail(emai string) (*User, error) {
	var user User
	result := repo.database.DB.First(&user, "email = ?", emai)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
