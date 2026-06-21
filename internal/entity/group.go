package entity

type Group struct {
	Id       Id     `json:"id"`
	ParentId Id     `json:"parent_id"`
	Name     string `json:"name"`
}
