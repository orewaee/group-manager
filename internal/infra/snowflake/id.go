package snowflake

import (
	"github.com/bytentropy/snowflake"
	"github.com/orewaee/group-manager/internal/entity"
	"github.com/orewaee/group-manager/internal/usecase/id"
)

type snowflakeIdProvider struct {
	nodeId    int64
	generator *snowflake.Generator
}

func (s *snowflakeIdProvider) Generate() entity.Id {
	return s.generator.Next()
}

func NewIdProvider(nodeId int64) id.Provider {
	return &snowflakeIdProvider{
		nodeId:    nodeId,
		generator: snowflake.NewGenerator(nodeId),
	}
}
