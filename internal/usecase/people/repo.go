package people

import (
	"context"

	"github.com/orewaee/group-manager/internal/entity"
)

type Repo interface {
	FindById(context.Context, entity.Id) (*entity.Person, error)
	FindByGroupId(context.Context, entity.Id, bool) ([]*entity.Person, error)
	Save(context.Context, *entity.Person) error
	Update(context.Context, *entity.Person) error
	Delete(context.Context, entity.Id) error
}
