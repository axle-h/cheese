//+build wireinject

package cmd

import (
	"github.com/axle-h/cheese/config"
	"github.com/axle-h/cheese/organize"
	"github.com/axle-h/cheese/store"
	"github.com/google/wire"
)

func InitializeOrganize() (organize.Organize, error) {
	wire.Build(NewCheeseConfig, config.NewMongoConfig, store.NewMongoContext, store.NewPhotoRepository, organize.NewOrganize)
	return organize.Organize{}, nil
}
