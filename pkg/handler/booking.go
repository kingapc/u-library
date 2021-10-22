package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/pkg/dal"
	"github.com/rpinedafocus/u-library/pkg/model"
)

// create handles the user create request
func CreateBookingController(c *gin.Context) (*model.BookingRentEntity, string) {

	booking := &model.BookingRent{}

	if err := c.BindJSON(&booking); err != nil {
		return nil, err.Error() //utils.ErrorX(400)
	}

	entity, err := dal.CreateBooking(booking)
	if err != nil {
		return nil, err.Error() //utils.ErrorX(400)
	}

	return entity, ""
}

// fetchByID will return an user by its id
// func (h *Handler) fetchByID() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		id := vars[userID]
// 		entity, err := h.UserDAO.FetchByID(r.Context(), id)
// 		switch {
// 		case errors.Is(err, errorx.ErrNoUser):
// 			msg := &errorMessage{
// 				Message: fmt.Sprintf("user %s does not exist", id),
// 			}
// 			response.JSON(w, http.StatusNotFound, msg)
// 			return
// 		case errors.Is(err, errorx.ErrDeleteUser):
// 			msg := &errorMessage{
// 				Message: fmt.Sprintf("user %s has been deleted", id),
// 			}
// 			response.JSON(w, http.StatusGone, msg)
// 			return
// 		case err != nil:
// 			msg := &errorMessage{
// 				Error:   err.Error(),
// 				Message: "user datastore error",
// 			}
// 			response.JSON(w, http.StatusInternalServerError, msg)
// 			return
// 		default:
// 			response.JSON(w, http.StatusOK, entity)
// 		}

// 	}
// }

// // list will return all of the users
// func (h *Handler) list() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		entities, err := h.UserDAO.FetchAll(r.Context())
// 		switch {
// 		case errors.Is(err, errorx.ErrNoUser):
// 			msg := &errorMessage{
// 				Message: fmt.Sprintf("no users exist"),
// 			}
// 			response.JSON(w, http.StatusNotFound, msg)
// 			return
// 		case err != nil:
// 			msg := &errorMessage{
// 				Error:   err.Error(),
// 				Message: "user datastore error",
// 			}
// 			response.JSON(w, http.StatusInternalServerError, msg)
// 			return
// 		default:
// 			response.JSON(w, http.StatusOK, entities)
// 		}
// 	}
// }

// // update will return the updated user
// func (h *Handler) update() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		user := &model.User{}
// 		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
// 			msg := &errorMessage{
// 				Error:   err.Error(),
// 				Message: "user json decode error",
// 			}
// 			response.JSON(w, http.StatusBadRequest, msg)
// 			return
// 		}
// 		if len(user.FirstName) == 0 && len(user.LastName) == 0 {
// 			msg := &errorMessage{
// 				Message: "user must have fields to update",
// 			}
// 			response.JSON(w, http.StatusBadRequest, msg)
// 			return
// 		}

// 		vars := mux.Vars(r)
// 		id := vars[userID]
// 		entity, err := h.UserDAO.Update(r.Context(), id, user)
// 		switch {
// 		case errors.Is(err, errorx.ErrNoUser):
// 			msg := &errorMessage{
// 				Message: fmt.Sprintf("user %s does not exist", id),
// 			}
// 			response.JSON(w, http.StatusNotFound, msg)
// 			return
// 		case errors.Is(err, errorx.ErrDeleteUser):
// 			msg := &errorMessage{
// 				Message: fmt.Sprintf("user %s has been deleted", id),
// 			}
// 			response.JSON(w, http.StatusGone, msg)
// 			return
// 		case err != nil:
// 			msg := &errorMessage{
// 				Error:   err.Error(),
// 				Message: "user datastore error",
// 			}
// 			response.JSON(w, http.StatusInternalServerError, msg)
// 			return
// 		default:
// 			response.JSON(w, http.StatusOK, entity)
// 		}

// 	}
// }

// // delete will remove the user
// func (h *Handler) delete() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		id := vars[userID]
// 		err := h.UserDAO.Delete(r.Context(), id)
// 		switch {
// 		case errors.Is(err, errorx.ErrNoUser):
// 			msg := &errorMessage{
// 				Message: fmt.Sprintf("user %s does not exist", id),
// 			}
// 			response.JSON(w, http.StatusNotFound, msg)
// 			return
// 		case errors.Is(err, errorx.ErrDeleteUser):
// 			msg := &errorMessage{
// 				Message: fmt.Sprintf("user %s has been deleted", id),
// 			}
// 			response.JSON(w, http.StatusGone, msg)
// 			return
// 		case err != nil:
// 			msg := &errorMessage{
// 				Error:   err.Error(),
// 				Message: "user datastore error",
// 			}
// 			response.JSON(w, http.StatusInternalServerError, msg)
// 			return
// 		default:
// 			response.JSON(w, http.StatusNoContent, nil)
// 		}
// 	}

// }

// Add will configure the routes for user operations
// func (h *Handler) Add(router *mux.Router) {
// 	router.Methods(http.MethodPost).Path("/user").Handler(h.create()).Name("user-create")
// 	router.Methods(http.MethodGet).Path(fmt.Sprintf("/users/{%s}", userID)).Handler(h.fetchByID()).Name("user-fetch")
// 	router.Methods(http.MethodGet).Path("/users").Handler(h.list()).Name("user-fetch-all")
// 	router.Methods(http.MethodPatch).Path(fmt.Sprintf("/users/{%s}", userID)).Handler(h.update()).Name("user-update")
// 	router.Methods(http.MethodDelete).Path(fmt.Sprintf("/users/{%s}", userID)).Handler(h.delete()).Name("user-delete")
// }
