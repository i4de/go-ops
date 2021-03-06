package script

import "go-ops/internal/model"

type Script interface {
	Path() string
	Run() model.ResCmd
}

//counterfeiter:generate . CancellableScript

type CancellableScript interface {
	Script
	Cancel() error
}
