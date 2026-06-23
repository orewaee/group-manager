package id

import "github.com/orewaee/group-manager/internal/entity"

type Provider interface {
	Generate() entity.Id
}
