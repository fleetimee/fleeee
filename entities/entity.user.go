package entities

import (
	"github.com/fleetimee/flee/models"
	"github.com/fleetimee/flee/schemes"
)

type EntityUser interface {
	EntityRegister(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError)
	EntityLogin(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError)
}
