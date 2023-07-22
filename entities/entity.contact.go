package entities

import (
	"github.com/fleetimee/flee/models"
	"github.com/fleetimee/flee/schemes"
)

type EntityContact interface {
	EntityCreate(input *schemes.SchemeContact) (*models.ModelContact, schemes.SchemeDatabaseError)
}
