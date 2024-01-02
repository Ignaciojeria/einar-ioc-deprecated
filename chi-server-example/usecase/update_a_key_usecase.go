package usecase

import (
	"example/kvdatabase"
	"fmt"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type updateAKeyUsecase struct {
	kvDB          kvdatabase.IKVDB
	getCurrentKey IgetCurrentKey
}

type IUpdateAKey func(value string)

var UpdateAKey = ioc.UseCase[IUpdateAKey](func() (IUpdateAKey, error) {
	usecase := updateAKeyUsecase{
		kvDB:          kvdatabase.KVDB.Dependency,
		getCurrentKey: GetCurrentKey.Dependency,
	}
	return usecase.execute, nil
})

func (u updateAKeyUsecase) execute(value string) {
	currentAValue := u.kvDB.Get("a")
	fmt.Println("current a value : " + currentAValue)
	u.kvDB.UpdateA(value)
}
