package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/orewaee/group-manager/internal/entity"
	"github.com/orewaee/group-manager/internal/infra/postgres/db"
	"github.com/orewaee/group-manager/internal/usecase/group"
)

type postgresGroupRepo struct {
	conn    *pgx.Conn
	queries *db.Queries
}

func groupFromDatabase(group db.Group) *entity.Group {
	g := &entity.Group{
		Id:   group.Id,
		Name: group.Name,
	}

	if group.ParentId.Valid {
		g.ParentId = &group.ParentId.Int64
	}

	return g
}

func (p *postgresGroupRepo) FindById(ctx context.Context, id entity.Id) (*entity.Group, error) {
	group, err := p.queries.SelectGroupById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find group: %w", err)
	}

	return groupFromDatabase(group), nil
}

func (p *postgresGroupRepo) FindByName(ctx context.Context, name string) ([]*entity.Group, error) {
	groups, err := p.queries.SelectGroupsByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("find groups: %w", err)
	}

	processed := make([]*entity.Group, len(groups))
	for i := range processed {
		processed[i] = groupFromDatabase(groups[i])
	}

	return processed, nil
}

func (p *postgresGroupRepo) FindAll(ctx context.Context) ([]*entity.Group, error) {
	groups, err := p.queries.SelectAllGroups(ctx)
	if err != nil {
		return nil, fmt.Errorf("find all groups: %w", err)
	}

	processed := make([]*entity.Group, len(groups))
	for i := range processed {
		processed[i] = groupFromDatabase(groups[i])
	}

	return processed, nil
}

func (p *postgresGroupRepo) FindChildren(ctx context.Context, parentId entity.Id) ([]*entity.Group, error) {
	groups, err := p.queries.SelectChildGroups(ctx, pgtype.Int8{Int64: parentId, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("find children: %w", err)
	}

	processed := make([]*entity.Group, len(groups))
	for i := range processed {
		processed[i] = groupFromDatabase(groups[i])
	}

	return processed, nil
}

func (p *postgresGroupRepo) CountMembers(ctx context.Context, id entity.Id, deep bool) (int, error) {
	var count int64
	var err error

	if deep {
		count, err = p.queries.CountTotalMembers(ctx, id)
	} else {
		count, err = p.queries.CountDirectMembers(ctx, id)
	}

	if err != nil {
		return 0, fmt.Errorf("count members: %w", err)
	}

	return int(count), nil
}

func (p *postgresGroupRepo) Save(ctx context.Context, group *entity.Group) error {
	return withTx(ctx, p.conn, func(queries *db.Queries) error {
		if group.ParentId != nil {
			_, err := queries.SelectGroupById(ctx, *group.ParentId)
			if errors.Is(err, pgx.ErrNoRows) {
				return entity.ErrGroupNotFound
			}

			if err != nil {
				return fmt.Errorf("select parent group: %w", err)
			}
		}

		params := db.InsertGroupParams{
			Id:   group.Id,
			Name: group.Name,
		}

		if group.ParentId != nil {
			params.ParentId = pgtype.Int8{Int64: *group.ParentId, Valid: true}
		}

		_, err := queries.InsertGroup(ctx, params)
		if err != nil {
			return fmt.Errorf("insert group: %w", err)
		}

		return nil
	})
}

func (p *postgresGroupRepo) Update(ctx context.Context, group *entity.Group) error {
	return withTx(ctx, p.conn, func(queries *db.Queries) error {
		_, err := queries.SelectGroupById(ctx, group.Id)
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.ErrGroupNotFound
		}

		if err != nil {
			return fmt.Errorf("select group: %w", err)
		}

		if group.ParentId != nil {
			_, err := queries.SelectGroupById(ctx, *group.ParentId)
			if errors.Is(err, pgx.ErrNoRows) {
				return entity.ErrGroupNotFound
			}

			if err != nil {
				return fmt.Errorf("select parent group: %w", err)
			}
		}

		params := db.UpdateGroupParams{
			Name: group.Name,
			Id:   group.Id,
		}

		if group.ParentId != nil {
			params.ParentId = pgtype.Int8{Int64: *group.ParentId, Valid: true}
		}

		if err := queries.UpdateGroup(ctx, params); err != nil {
			return fmt.Errorf("update group: %w", err)
		}

		return nil
	})
}

func (p *postgresGroupRepo) Delete(ctx context.Context, id entity.Id) error {
	if err := p.queries.DeleteGroup(ctx, id); err != nil {
		return fmt.Errorf("delete group: %w", err)
	}

	return nil
}

func NewGroupRepo(conn *pgx.Conn) group.Repo {
	return &postgresGroupRepo{
		queries: db.New(conn),
		conn:    conn,
	}
}
