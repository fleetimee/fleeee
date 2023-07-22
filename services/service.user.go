package services

import (
	"github.com/fleetimee/flee/entities"
	"github.com/fleetimee/flee/models"
	"github.com/fleetimee/flee/schemes"
)

type serviceUser struct {
	user entities.EntityUser
}

func NewServiceUser(user entities.EntityUser) *serviceUser {
	return &serviceUser{user: user}
}

func (s *serviceUser) EntityRegister(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError) {
	var schema schemes.SchemeUser
	schema.FirstName = input.FirstName
	schema.LastName = input.LastName
	schema.Email = input.Email
	schema.Password = input.Password
	schema.Role = input.Role

	res, err := s.user.EntityRegister(&schema)
	return res, err
}

func (s *serviceUser) EntityLogin(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError) {
	var schema schemes.SchemeUser
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.user.EntityLogin(&schema)
	return res, err
}
