package group

import (
	"context"

	"github.com/orewaee/group-manager/internal/entity"
	"github.com/orewaee/group-manager/internal/usecase/id"
	"github.com/orewaee/group-manager/internal/usecase/people"
)

type Group interface {
	Create(context.Context, CreateCmd) (*entity.Group, error)
	Update(context.Context, UpdateCmd) (*entity.Group, error)
	Delete(context.Context, DeleteCmd) error
	GetAll(context.Context) ([]*entity.Group, error)
	Count(context.Context, CountCmd) (int, error)
	Members(context.Context, MembersCmd) ([]*entity.Person, error)
}

func New(idProvider id.Provider, groupRepo Repo, peopleRepo people.Repo) Group {
	return &service{
		idProvider: idProvider,
		groupRepo:  groupRepo,
		peopleRepo: peopleRepo,
	}
}
