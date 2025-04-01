package commands

import (
	"github.com/warrco/gator/internal/config"
	"github.com/warrco/gator/internal/database"
)

type State struct {
	Config *config.Config
	Db     *database.Queries
}
