// +build wireinject

package wire

import (
	"github.com/dennys-bd/glow/repository"
	"github.com/dennys-bd/glow/usecase"
	"github.com/google/wire"
)

var classRepoSet = wire.NewSet(
	repository.ProvideClassRepo,
	wire.Bind(new(repository.ClassRepoInf), new(repository.ClassRepository)),
)

func InjectClassController() (usecase.ClassController, error) {
	panic(wire.Build(
		classRepoSet,
		usecase.ProvideClassCtrl,
	))
}
