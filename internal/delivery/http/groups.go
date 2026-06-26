package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/orewaee/group-manager/internal/entity"
	"github.com/orewaee/group-manager/internal/usecase/group"
	"github.com/rs/zerolog/log"
)

func (h *Handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	data, err := read[*CreateGroupRequest](r)
	if err != nil {
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to read request")

		return
	}

	grp, err := h.groupApi.Create(r.Context(), group.CreateCmd{
		ParentId: data.ParentId,
		Name:     data.Name,
	})

	if errors.Is(err, entity.ErrGroupNotFound) {
		log.Error().Err(err).Send()
		writeError(w, http.StatusBadRequest, "parent group not found")

		return
	}

	if err != nil {
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to create group")

		return
	}

	log.Debug().
		Int64("id", grp.Id).
		Msg("group created")

	writeJson(w, http.StatusCreated, &Group{
		Id:       grp.Id,
		ParentId: grp.ParentId,
		Name:     grp.Name,
	})
}

func (h *Handler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	data, err := read[*UpdateGroupRequest](r)
	if err != nil {
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to read request")

		return
	}

	grp, err := h.groupApi.Update(r.Context(), group.UpdateCmd{
		Id:       data.Id,
		ParentId: data.ParentId,
		Name:     data.Name,
	})

	if errors.Is(err, entity.ErrGroupNotFound) {
		log.Error().Err(err).Send()
		writeError(w, http.StatusBadRequest, "group not found")

		return
	}

	if err != nil {
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to update group")

		return
	}

	log.Debug().
		Int64("id", grp.Id).
		Msg("group updated")

	writeJson(w, http.StatusOK, &Group{
		Id:       grp.Id,
		ParentId: grp.ParentId,
		Name:     grp.Name,
	})
}

func (h *Handler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")

		return
	}

	err = h.groupApi.Delete(r.Context(), group.DeleteCmd{Id: id})
	if err != nil {
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to delete group")

		return
	}

	log.Debug().
		Int64("id", id).
		Msg("group deleted")

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ListGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := h.groupApi.GetAll(r.Context())
	if err != nil {
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to list groups")

		return
	}

	response := make([]*GroupWithCount, len(groups))
	for i, g := range groups {
		directCount, err := h.groupApi.Count(r.Context(), group.CountCmd{Id: g.Id, Deep: false})
		if err != nil {
			log.Error().Err(err).Send()
			writeError(w, http.StatusInternalServerError, "failed to count members")

			return
		}

		totalCount, err := h.groupApi.Count(r.Context(), group.CountCmd{Id: g.Id, Deep: true})
		if err != nil {
			log.Error().Err(err).Send()
			writeError(w, http.StatusInternalServerError, "failed to count members")

			return
		}

		response[i] = &GroupWithCount{
			Id:          g.Id,
			ParentId:    g.ParentId,
			Name:        g.Name,
			DirectCount: directCount,
			TotalCount:  totalCount,
		}
	}

	writeJson(w, http.StatusOK, response)
}

func (h *Handler) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")

		return
	}

	deep := r.URL.Query().Get("deep") == "true"

	members, err := h.groupApi.Members(r.Context(), group.MembersCmd{Id: id, Deep: deep})
	if err != nil {
		log.Error().Err(err).Send()
		writeError(w, http.StatusInternalServerError, "failed to get members")

		return
	}

	response := make([]*Person, len(members))
	for i, m := range members {
		response[i] = &Person{
			Id:        m.Id,
			GroupId:   m.GroupId,
			Firstname: m.Firstname,
			Lastname:  m.Lastname,
			Birthday:  m.Birthday,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		}
	}

	writeJson(w, http.StatusOK, response)
}
