package handlers

import (
	"net/http"

	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

//swagger:parameters AuthUser
type AuthUserParams struct {
	// Required: true
	// in: body
	Body structs.AuthUser
}

// swagger:route POST /auth user AuthUser
// Authorization user method
//		Responses:
//			default: ErrorResponse
//			200:
func (hs *Handlers) Auth(w http.ResponseWriter, r *http.Request) {
	var auth structs.AuthUser
	err := ReadBody(r, &auth)

	if err != nil {
		e := errors.NewError(errors.BAD_REQUEST_ERROR)
		if fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, e); fatal != nil {
			// log FATAL_ERROR
		}
		return
	}

	passHash := HashPassword(auth.Password)

	user, err := hs.database.AuthUser(auth.Login, passHash)
	if err != nil {
		e := &errors.ErrorResponse{
			Code:    errors.NOT_AUTHORIZED_ERROR,
			Message: err.Error(),
		}

		if fatal := errors.WriteHttpErrorMessage(w, http.StatusUnauthorized, e); fatal != nil {
			// log FATAL_ERROR
		}
		return
	}

	hs.secure.SetSession(w, user)
}
