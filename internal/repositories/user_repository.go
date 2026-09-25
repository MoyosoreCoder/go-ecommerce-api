package repositories

import (
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	

)


type UserRepository struct {
	DB *pgxpool.Pool

}

func (r UserRepository) CreateUser(user models.User) error {
	r.DB

}