package entities

import (
	"github.com/fleetimee/flee/models"
	"github.com/fleetimee/flee/schemes"
)

type EntityOutlet interface {
	EntityCreate(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelOutlet, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeOutlet) (*models.ModelOutlet, schemes.SchemeDatabaseError)
}
