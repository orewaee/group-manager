package group

import (
	"context"

	"github.com/orewaee/group-manager/internal/entity"
)

type service struct {
}

func (s *service) Create(context.Context, CreateCmd) (*entity.Group, error) {
	panic("unimplemented")
}

func (s *service) Delete(context.Context, DeleteCmd) error {
	panic("unimplemented")
}

func (s *service) Update(context.Context, UpdateCmd) (*entity.Group, error) {
	panic("unimplemented")
}

func (s *service) Count(context.Context, CountCmd) (int, error) {
	panic("unimplemented")
}

func (s *service) Members(context.Context, MembersCmd) ([]*entity.Person, error) {
	panic("unimplemented")
}
