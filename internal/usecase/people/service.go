package people

import (
	"context"
	"fmt"
	"time"

	"github.com/orewaee/group-manager/internal/entity"
	"github.com/orewaee/group-manager/internal/usecase/id"
)

type service struct {
	idProvider id.Provider
	peopleRepo Repo
}

func (s *service) Create(ctx context.Context, cmd CreateCmd) (*entity.Person, error) {
	now := time.Now()
	person := &entity.Person{
		Id:        s.idProvider.Generate(),
		GroupId:   cmd.GroupId,
		Firstname: cmd.Firstname,
		Lastname:  cmd.Lastname,
		Birthday:  cmd.Birthday,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.peopleRepo.Save(ctx, person); err != nil {
		return nil, fmt.Errorf("create person: %w", err)
	}

	return person, nil
}

func (s *service) Update(ctx context.Context, cmd UpdateCmd) (*entity.Person, error) {
	person := &entity.Person{
		Id:        cmd.Id,
		GroupId:   cmd.GroupId,
		Firstname: cmd.Firstname,
		Lastname:  cmd.Lastname,
		Birthday:  cmd.Birthday,
		UpdatedAt: time.Now(),
	}

	if err := s.peopleRepo.Update(ctx, person); err != nil {
		return nil, fmt.Errorf("update person: %w", err)
	}

	return person, nil
}

func (s *service) Delete(ctx context.Context, cmd DeleteCmd) error {
	if err := s.peopleRepo.Delete(ctx, cmd.Id); err != nil {
		return fmt.Errorf("delete person: %w", err)
	}

	return nil
}
