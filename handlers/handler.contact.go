package handlers

import (
	"net/http"

	"github.com/fleetimee/flee/entities"
	"github.com/fleetimee/flee/helpers"
	"github.com/fleetimee/flee/pkg"
	"github.com/fleetimee/flee/schemes"
	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handlerContact struct {
	contact entities.EntityContact
}

func NewHandlerContact(contact entities.EntityContact) *handlerContact {
	return &handlerContact{contact: contact}
}

func (h *handlerContact) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Contact", http.StatusOK, nil)
}

func (h *handlerContact) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeContact
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorContact(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.contact.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Contact name already exist", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new contact success", http.StatusCreated, nil)
}

func ValidatorContact(
	ctx *gin.Context,
	input schemes.SchemeContact,
	Type string,
) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "Realname",
					Message: "Realname is required on body",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
