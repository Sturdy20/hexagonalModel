package services

import (
	"HexagonalModel/modules/project1/models"
	"HexagonalModel/modules/project1/repositories"

	"log"
)

type IService interface {
	PotsRegisterSer(register models.RequestRegister) error
}

type service struct {
	r repositories.IRepositorie
}

func NewService(r repositories.IRepositorie) IService {
	return &service{r: r}
}

func (s *service) PotsRegisterSer(register models.RequestRegister) error {
	err := s.r.PotsRegister(register)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
