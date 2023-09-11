package user

import (
	"github.com/eXoterr/StrimaNET_Backend/internal/store/models"
	"github.com/jmoiron/sqlx"
)

type UserStore interface {
	Create(user models.User) (models.User, error)
	GetById(id uint) (models.User, error)
	Delete(id uint) error
	Update(user models.User) error
}

type store struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) UserStore {
	return &store{DB: db}
}

func (s *store) Create(user models.User) (models.User, error) {
	result, err := s.DB.NamedQuery("INSERT INTO Users(login, password, email) VALUES (:login, :password, :email) RETURNING id", models.User{
		Login:    user.Login,
		Password: user.Password,
		Email:    user.Email,
	})

	if err != nil {
		return models.User{}, err
	}

	resultModel := models.User{}

	sqlx.StructScan(result, resultModel)

	return resultModel, nil
}

func (s *store) GetById(id uint) (models.User, error) {
	return models.User{}, nil
}

func (s *store) Delete(id uint) error {
	return nil
}

func (s *store) Update(user models.User) error {
	return nil
}
