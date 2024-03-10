package repositories

import (
	"HexagonalModel/modules/project1/models"
	"database/sql"
)

type IRepositorie interface {
	PotsRegister(register models.RequestRegister) error
}

type repositorie struct {
	db *sql.DB
}

func NewRepositorie(db *sql.DB) IRepositorie {
	return &repositorie{db: db}
}

func (r *repositorie) PotsRegister(register models.RequestRegister) error {

	return nil
}
