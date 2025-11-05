package repositories

import "github.com/sizzlorox/sols-cms/pkg/providers/database"

type IUserRepository interface {
	GetUserByID(id int64) (interface{}, error)
	GetUsers() ([]interface{}, error)
	CreateUser(data interface{}) (interface{}, error)
	UpdateUserByID(id int64, data interface{}) (interface{}, error)
	DeleteUserByID(id int64) error
}

type userRepository struct {
	db database.IRepository
}

func NewUserRepository(db database.IRepository) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetUserByID(id int64) (interface{}, error) {
	return nil, nil
}

func (ur *userRepository) GetUsers() ([]interface{}, error) {
	return nil, nil
}

func (ur *userRepository) CreateUser(data interface{}) (interface{}, error) {
	return nil, nil
}

func (ur *userRepository) UpdateUserByID(id int64, data interface{}) (interface{}, error) {
	return nil, nil
}

func (ur *userRepository) DeleteUserByID(id int64) error {
	return nil
}
