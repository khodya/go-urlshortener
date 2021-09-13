package storage

type Db map[string]string

var db Db

func init() {
	db = make(Db)
}

func Put(key, value string) {
	db[key] = value
}

func Get(key string) (string, bool) {
	v, ok := db[key]
	return v, ok
}
