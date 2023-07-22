package entities

import (
	"github.com/fleetimee/flee/models"
	"github.com/fleetimee/flee/schemes"
)

type EntityTransaction interface {
	EntityCreate(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelTransaction, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeTransaction) (*models.ModelTransaction, schemes.SchemeDatabaseError)
}
