package toolkit

import (
	"conflict-template/internal/service/log"
	"github.com/go-conflict/toolkit"
)

func Init() {
	toolkit.Init(log.ProvideLogger())
}
