package http

import (
	"net/http"

	"github.com/orewaee/group-manager/internal/usecase/group"
	"github.com/orewaee/group-manager/internal/usecase/people"
)

type Handler struct {
	peopleApi people.People
	groupApi  group.Group
}

func NewHander(peopleApi people.People, groupApi group.Group) *Handler {
	return &Handler{
		peopleApi: peopleApi,
		groupApi:  groupApi,
	}
}

/*
  - CreatePerson)
    r.Put("/people/:id", handler.UpdatePerson)
    r.Delete("/people/:id", handler.DeletePerson)
*/
func (h *Handler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	data, err := read[*CreatePersonRequest](r)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, &ErrorResponse{
			Message: "failed to read request",
		})

		return
	}

	person, err := h.peopleApi.Create(r.Context(), people.CreateCmd{
		GroupId:   data.GroupId,
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Birthday:  data.Birthday,
	})

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, &ErrorResponse{
			Message: "failed to create person",
		})

		return
	}

	writeJSON(w, http.StatusCreated, &CreatePersonResponse{
		GroupId:   person.GroupId,
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Birthday:  person.Birthday,
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
	})
}
