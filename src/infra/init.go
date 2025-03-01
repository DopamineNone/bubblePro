package infra

import (
	"github.com/DopamineNone/bubblePro/src/infra/db"
	idGen "github.com/DopamineNone/bubblePro/src/utils/id_generator"
)

func Init() {
	db.Init()
	idGen.Init()
}
