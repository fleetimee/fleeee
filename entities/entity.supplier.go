package entities

import (
	"github.com/fleetimee/flee/models"
	"github.com/fleetimee/flee/schemes"
)

type EntitySupplier interface {
	EntityCreate(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelSupplier, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeSupplier) (*models.ModelSupplier, schemes.SchemeDatabaseError)
}
