package storage

type DB map[string]string

var db DB

func init() {
	db = make(DB)
}

func Put(key, value string) {
	db[key] = value
}

func Get(key string) (string, bool) {
	v, ok := db[key]
	return v, ok
}
