package entities

import (
	"github.com/fleetimee/flee/models"
	"github.com/fleetimee/flee/schemes"
)

type EntityProduct interface {
	EntityCreate(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelProduct, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeProduct) (*models.ModelProduct, schemes.SchemeDatabaseError)
	EntityProductByOutlet(input *schemes.SchemeProduct) (*[]models.ModelProduct, schemes.SchemeDatabaseError)
}
