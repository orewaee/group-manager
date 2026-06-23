package group

import (
	"context"

	"github.com/orewaee/group-manager/internal/entity"
)

type Group interface {
	Create(context.Context, CreateCmd) (*entity.Group, error)
	Update(context.Context, UpdateCmd) (*entity.Group, error)
	Delete(context.Context, DeleteCmd) error
	Count(context.Context, CountCmd) (int, error)
	Members(context.Context, MembersCmd) ([]*entity.Person, error)
}

func New() Group {
	return &service{}
}
