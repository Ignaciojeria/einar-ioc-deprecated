package usecase

import (
	"example/kvdatabase"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type updateAKeyUsecase struct {
	kvDB          kvdatabase.IKVDB
	getCurrentKey IgetCurrentKey
}

type IUpdateAKey func(value string)

var UpdateAKey = ioc.InjectUseCase[IUpdateAKey](func() (IUpdateAKey, error) {
	usecase := updateAKeyUsecase{
		kvDB:          kvdatabase.KVDB.Dependency,
		getCurrentKey: GetCurrentKey.Dependency,
	}
	return usecase.execute, nil
})

func (u updateAKeyUsecase) execute(value string) {
	u.kvDB.UpdateA(value)
}
