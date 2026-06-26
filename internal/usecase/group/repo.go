package group

import (
	"context"

	"github.com/orewaee/group-manager/internal/entity"
)

type Repo interface {
	FindById(context.Context, entity.Id) (*entity.Group, error)
	FindByName(context.Context, string) ([]*entity.Group, error)
	FindAll(context.Context) ([]*entity.Group, error)
	FindChildren(context.Context, entity.Id) ([]*entity.Group, error)
	CountMembers(context.Context, entity.Id, bool) (int, error)
	Save(context.Context, *entity.Group) error
	Update(context.Context, *entity.Group) error
	Delete(context.Context, entity.Id) error
}
