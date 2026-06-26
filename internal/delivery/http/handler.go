package http

import (
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
