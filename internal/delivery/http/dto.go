package http

import (
	"time"

	"github.com/orewaee/group-manager/internal/entity"
)

type CreatePersonRequest struct {
	GroupId   entity.Id `json:"group_id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Birthday  time.Time `json:"birthday"`
}

type UpdatePersonRequest struct {
	Id        entity.Id `json:"id"`
	GroupId   entity.Id `json:"group_id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Birthday  time.Time `json:"birthday"`
}

type Person struct {
	Id        entity.Id `json:"id"`
	GroupId   entity.Id `json:"group_id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateGroupRequest struct {
	ParentId *entity.Id `json:"parent_id"`
	Name     string     `json:"name"`
}

type UpdateGroupRequest struct {
	Id       entity.Id  `json:"id"`
	ParentId *entity.Id `json:"parent_id"`
	Name     string     `json:"name"`
}

type Group struct {
	Id       entity.Id  `json:"id"`
	ParentId *entity.Id `json:"parent_id"`
	Name     string     `json:"name"`
}

type GroupWithCount struct {
	Id          entity.Id  `json:"id"`
	ParentId    *entity.Id `json:"parent_id"`
	Name        string     `json:"name"`
	DirectCount int        `json:"direct_count"`
	TotalCount  int        `json:"total_count"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
