package usecase

import (
	"example/kvdatabase"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type updateBKeyUsecase struct {
	kvDB          kvdatabase.IKVDB
	getCurrentKey IgetCurrentKey
}

type IUpdateBKey func(value string)

var UpdateBKey = ioc.InjectUseCase[IUpdateBKey](func() (IUpdateBKey, error) {
	usecase := updateBKeyUsecase{
		kvDB:          kvdatabase.KVDB.Dependency,
		getCurrentKey: GetCurrentKey.Dependency,
	}
	return usecase.execute, nil
})

func (u updateBKeyUsecase) execute(value string) {
	u.kvDB.UpdateB(value)
}
