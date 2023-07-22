package services

import (
	"github.com/fleetimee/flee/entities"
	"github.com/fleetimee/flee/models"
	"github.com/fleetimee/flee/schemes"
)

type serviceContact struct {
	contact entities.EntityContact
}

func NewServiceContact(contact entities.EntityContact) *serviceContact {
	return &serviceContact{contact: contact}
}

func (s *serviceContact) EntityCreate(
	input *schemes.SchemeContact,
) (*models.ModelContact, schemes.SchemeDatabaseError) {
	var contact schemes.SchemeContact

	contact.Realname = input.Realname
	contact.Email = input.Email
	contact.Whatsapp = input.Whatsapp

	res, err := s.contact.EntityCreate(&contact)
	return res, err
}
