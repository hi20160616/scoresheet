package scoresheet

import (
	"log"

	"github.com/hi20160616/scoresheet/configs"
)

type Opts struct {
	Cfg *configs.Config
	Log *log.Logger
}

func NewOpts(cfg *configs.Config, log *log.Logger) *Opts {
	return &Opts{cfg, log}
}
