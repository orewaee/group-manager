package group

import (
	"context"
	"fmt"

	"github.com/orewaee/group-manager/internal/entity"
	"github.com/orewaee/group-manager/internal/usecase/id"
	"github.com/orewaee/group-manager/internal/usecase/people"
)

type service struct {
	idProvider id.Provider
	groupRepo  Repo
	peopleRepo people.Repo
}

func (s *service) Create(ctx context.Context, cmd CreateCmd) (*entity.Group, error) {
	group := &entity.Group{
		Id:       s.idProvider.Generate(),
		ParentId: cmd.ParentId,
		Name:     cmd.Name,
	}

	if err := s.groupRepo.Save(ctx, group); err != nil {
		return nil, fmt.Errorf("create group: %w", err)
	}

	return group, nil
}

func (s *service) Update(ctx context.Context, cmd UpdateCmd) (*entity.Group, error) {
	group := &entity.Group{
		Id:       cmd.Id,
		ParentId: cmd.ParentId,
		Name:     cmd.Name,
	}

	if err := s.groupRepo.Update(ctx, group); err != nil {
		return nil, fmt.Errorf("update group: %w", err)
	}

	return group, nil
}

func (s *service) Delete(ctx context.Context, cmd DeleteCmd) error {
	if err := s.groupRepo.Delete(ctx, cmd.Id); err != nil {
		return fmt.Errorf("delete group: %w", err)
	}

	return nil
}

func (s *service) Count(ctx context.Context, cmd CountCmd) (int, error) {
	count, err := s.groupRepo.CountMembers(ctx, cmd.Id, cmd.Deep)
	if err != nil {
		return 0, fmt.Errorf("count members: %w", err)
	}

	return count, nil
}

func (s *service) GetAll(ctx context.Context) ([]*entity.Group, error) {
	groups, err := s.groupRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all groups: %w", err)
	}

	return groups, nil
}

func (s *service) Members(ctx context.Context, cmd MembersCmd) ([]*entity.Person, error) {
	people, err := s.peopleRepo.FindByGroupId(ctx, cmd.Id, cmd.Deep)
	if err != nil {
		return nil, fmt.Errorf("find members: %w", err)
	}

	return people, nil
}
