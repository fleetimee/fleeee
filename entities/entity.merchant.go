package entities

import (
	"github.com/fleetimee/flee/models"
	"github.com/fleetimee/flee/schemes"
)

type EntityMerchant interface {
	EntityCreate(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelMerchant, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeMerchant) (*models.ModelMerchant, schemes.SchemeDatabaseError)
}
