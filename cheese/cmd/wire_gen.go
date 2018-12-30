// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cmd

import (
	"github.com/axle-h/cheese/config"
	"github.com/axle-h/cheese/organize"
	"github.com/axle-h/cheese/store"
)

// Injectors from wire.go:

func InitializeOrganize() (organize.Organize, error) {
	cheeseConfig, err := NewCheeseConfig()
	if err != nil {
		return organize.Organize{}, err
	}
	mongoConfig := config.NewMongoConfig()
	mongoContext, err := store.NewMongoContext(mongoConfig)
	if err != nil {
		return organize.Organize{}, err
	}
	photoRepository := store.NewPhotoRepository(mongoContext)
	organizeOrganize := organize.NewOrganize(cheeseConfig, photoRepository)
	return organizeOrganize, nil
}
