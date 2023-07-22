package entities

import (
	"github.com/fleetimee/flee/models"
	"github.com/fleetimee/flee/schemes"
)

type EntityRole interface {
	EntityCreate(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelRole, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeRole) (*models.ModelRole, schemes.SchemeDatabaseError)
}
