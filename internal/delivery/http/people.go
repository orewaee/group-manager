package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/orewaee/group-manager/internal/entity"
	"github.com/orewaee/group-manager/internal/usecase/people"
	"github.com/rs/zerolog/log"
)

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

	if errors.Is(err, entity.ErrPersonNotFound) {
		log.Error().Err(err).Send()
		writeError(w, http.StatusBadRequest, "person not found")

		return
	}

	if errors.Is(err, entity.ErrGroupNotFound) {
		log.Error().Err(err).Send()
		writeError(w, http.StatusBadRequest, "group not found")

		return
	}

	if err != nil {
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to update person")

		return
	}

	log.Debug().
		Int64("id", person.Id).
		Msg("person updated")

	writeJson(w, http.StatusOK, &Person{
		Id:        person.Id,
		GroupId:   person.GroupId,
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Birthday:  person.Birthday,
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
	})
}

func (h *Handler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")

		return
	}

	err = h.peopleApi.Delete(r.Context(), people.DeleteCmd{Id: id})
	if err != nil {
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to delete person")

		return
	}

	log.Debug().
		Int64("id", id).
		Msg("person deleted")

	w.WriteHeader(http.StatusNoContent)
}
