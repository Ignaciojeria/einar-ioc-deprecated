package kvdatabase

import ioc "github.com/Ignaciojeria/einar-ioc"

type IKVDB map[string]string

var KVDB = ioc.InjectOutBoundAdapter[IKVDB](func() (IKVDB, error) {
	return map[string]string{}, nil
})

func (kvdb IKVDB) UpdateA(value string) {
	kvdb["a"] = value
}

func (kvdb IKVDB) UpdateB(value string) {
	kvdb["b"] = value
}

func (kvdb IKVDB) Get(key string) string {
	return kvdb[key]
}
