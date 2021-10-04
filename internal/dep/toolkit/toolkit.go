package toolkit

import (
	//"back"
	"github.com/go-conflict/toolkit"
	"recruiting_system/internal/dep/log"
)

func init() {
	toolkit.Init(log.ProvideLogger())
}
