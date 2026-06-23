package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/orewaee/group-manager/internal/entity"
	"github.com/orewaee/group-manager/internal/infra/postgres/db"
	"github.com/orewaee/group-manager/internal/usecase/people"
)

type postgresPeopleRepo struct {
	queries *db.Queries
}

func (p *postgresPeopleRepo) FindById(ctx context.Context, id entity.Id) (*entity.Person, error) {
	person, err := p.queries.SelectPersonById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find person: %w", err)
	}

	return &entity.Person{
		Id:        person.Id,
		GroupId:   person.GroupId,
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Birthday:  person.Birthday.Time,
		CreatedAt: person.CreatedAt.Time,
		UpdatedAt: person.UpdatedAt.Time,
	}, nil
}

func (p *postgresPeopleRepo) FindByGroupId(ctx context.Context, groupId entity.Id, deep bool) ([]*entity.Person, error) {
	people, err := p.queries.SelectPersonByGroupId(ctx, groupId)
	if err != nil {
		return nil, fmt.Errorf("find people: %w", err)
	}

	processed := make([]*entity.Person, len(people))
	for i := range processed {
		processed[i] = &entity.Person{
			Id:        people[i].Id,
			GroupId:   people[i].GroupId,
			Firstname: people[i].Firstname,
			Lastname:  people[i].Lastname,
			Birthday:  people[i].Birthday.Time,
			CreatedAt: people[i].CreatedAt.Time,
			UpdatedAt: people[i].UpdatedAt.Time,
		}
	}

	return processed, nil
}

func (p *postgresPeopleRepo) Save(ctx context.Context, person *entity.Person) error {
	params := db.InsertPersonParams{
		Id:        person.Id,
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Birthday: pgtype.Date{
			Time: person.Birthday,
		},
		GroupId: person.GroupId,
	}

	if err := p.queries.InsertPerson(ctx, params); err != nil {
		return fmt.Errorf("save person: %w", err)
	}

	return nil
}

func (p *postgresPeopleRepo) Update(ctx context.Context, person *entity.Person) error {
	params := db.UpdatePersonParams{
		Id:        person.Id,
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Birthday: pgtype.Date{
			Time: person.Birthday,
		},
		GroupId: person.GroupId,
	}

	if err := p.queries.UpdatePerson(ctx, params); err != nil {
		return fmt.Errorf("update person: %w", err)
	}

	return nil
}

func (p *postgresPeopleRepo) Delete(ctx context.Context, id entity.Id) error {
	if err := p.queries.DeletePerson(ctx, id); err != nil {
		return fmt.Errorf("delete person: %w", err)
	}

	return nil
}

func NewPeopleRepo(conn db.DBTX) people.Repo {
	return &postgresPeopleRepo{
		queries: db.New(conn),
	}
}
