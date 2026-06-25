package http

import (
	"errors"
	"net/http"

	"github.com/orewaee/group-manager/internal/entity"
	"github.com/orewaee/group-manager/internal/usecase/group"
	"github.com/orewaee/group-manager/internal/usecase/people"
	"github.com/rs/zerolog/log"
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
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to read request")
		return
	}

	person, err := h.peopleApi.Create(r.Context(), people.CreateCmd{
		GroupId:   data.GroupId,
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Birthday:  data.Birthday,
	})

	if errors.Is(err, entity.ErrGroupNotFound) {
		log.Error().Err(err).Send()
		writeError(w, http.StatusBadRequest, "group not found")
		return
	}

	if err != nil {
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to create person")
		return
	}

	log.Debug().
		Int64("id", person.Id).
		Msg("person created")

	writeJson(w, http.StatusCreated, &Person{
		Id:        person.Id,
		GroupId:   person.GroupId,
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Birthday:  person.Birthday,
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
	})
}

func (h *Handler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	data, err := read[*UpdatePersonRequest](r)
	if err != nil {
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to read request")
		return
	}

	person, err := h.peopleApi.Update(r.Context(), people.UpdateCmd{
		Id:        data.Id,
		GroupId:   data.GroupId,
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Birthday:  data.Birthday,
	})

	switch {
	case errors.Is(err, entity.ErrPersonNotFound):
		log.Error().Err(err).Send()
		writeError(w, http.StatusBadRequest, "person not found")
		return
	case errors.Is(err, entity.ErrGroupNotFound):
		log.Error().Err(err).Send()
		writeError(w, http.StatusBadRequest, "group not found")
		return
	case err != nil:
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to update person")
		return
	}

	log.Debug().
		Int64("id", person.Id).
		Msg("person updated")

	writeJson(w, http.StatusCreated, &Person{
		Id:        person.Id,
		GroupId:   person.GroupId,
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Birthday:  person.Birthday,
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
	})
}
