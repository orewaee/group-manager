package group

import "github.com/orewaee/group-manager/internal/entity"

type CreateCmd struct {
	ParentId *entity.Id `json:"parent_id"`
	Name     string     `json:"name"`
}

type UpdateCmd struct {
	Id       entity.Id  `json:"id"`
	ParentId *entity.Id `json:"parent_id"`
	Name     string     `json:"name"`
}

type DeleteCmd struct {
	Id entity.Id `json:"id"`
}

type CountCmd struct {
	Id   entity.Id `json:"Id"`
	Deep bool      `json:"deep"`
}

type MembersCmd struct {
	Id   entity.Id `json:"id"`
	Deep bool      `json:"deep"`
}
