package people

import (
	"time"

	"github.com/orewaee/group-manager/internal/entity"
)

type CreateCmd struct {
	GroupId   entity.Id `json:"group_id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Birthday  time.Time `json:"birthday"`
}

type UpdateCmd struct {
	Id        entity.Id `json:"id"`
	GroupId   entity.Id `json:"group_id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Birthday  time.Time `json:"birthday"`
}

type DeleteCmd struct {
	Id entity.Id `json:"id"`
}
