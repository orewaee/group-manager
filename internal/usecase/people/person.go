package people

import (
	"context"

	"github.com/orewaee/group-manager/internal/entity"
	"github.com/orewaee/group-manager/internal/usecase/id"
)

type People interface {
	Create(context.Context, CreateCmd) (*entity.Person, error)
	Update(context.Context, UpdateCmd) (*entity.Person, error)
	Delete(context.Context, DeleteCmd) error
}

func New(idProvider id.Provider, peopleRepo Repo) People {
	return &service{
		idProvider: idProvider,
		peopleRepo: peopleRepo,
	}
}
