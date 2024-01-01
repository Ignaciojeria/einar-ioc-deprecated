package usecase

import (
	"example/kvdatabase"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type getCurrentKeyUsecase struct {
	kvDB kvdatabase.IKVDB
}

type IgetCurrentKey func(key string) string

var GetCurrentKey = ioc.InjectUseCase[IgetCurrentKey](func() (IgetCurrentKey, error) {
	usecase := getCurrentKeyUsecase{
		kvDB: kvdatabase.KVDB.Dependency,
	}
	return usecase.execute, nil
})

func (u getCurrentKeyUsecase) execute(key string) string {
	return u.kvDB.Get(key)
}
