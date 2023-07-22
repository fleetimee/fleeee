package repositories

import (
	"net/http"

	"github.com/fleetimee/flee/models"
	"github.com/fleetimee/flee/schemes"
	"gorm.io/gorm"
)

type repositoryContact struct {
	db *gorm.DB
}

func NewRepositoryContact(db *gorm.DB) *repositoryContact {
	return &repositoryContact{db: db}
}

func (r *repositoryContact) EntityCreate(
	input *schemes.SchemeContact,
) (*models.ModelContact, schemes.SchemeDatabaseError) {
	var contact models.ModelContact
	contact.Realname = input.Realname
	contact.Whatsapp = input.Whatsapp
	contact.Email = input.Email

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&contact)

	addContact := db.Debug().Create(&contact).Commit()

	if addContact.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_01",
		}
		return &contact, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &contact, <-err

}
