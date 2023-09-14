package snowflake

import (
	"errors"

	"github.com/bwmarrin/snowflake"

	"time"

	id_facade "open.chat/app/service/idgen/facade"
	"open.chat/pkg/log"
)

type SnowflakeUUIDGen struct {
	idgen *snowflake.Node
}

func New() id_facade.UUIDGen {
	c := new(SnowflakeUUIDGen)
	c.idgen, _ = snowflake.NewNode(time.Now().UnixNano() % 1024)
	return c
}

func (id *SnowflakeUUIDGen) GetUUID() (int64, error) {
	var err error
	if id.idgen == nil {
		err = errors.New("idgen not init")
		log.Error(err.Error())
		return 0, err
	}
	return id.idgen.Generate().Int64(), nil
}

func init() {
	id_facade.UUIDGenRegister("snowflake", New)
}
