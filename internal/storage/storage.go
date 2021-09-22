package storage

import (
	"encoding/gob"
	"flag"
	"os"
)

type Table map[string]string

type FileStore struct {
	FileName string
}

const (
	FileStorePathEnvName = "FILE_STORAGE_PATH"
)

var db Table
var fileStore *FileStore

func init() {
	fileStore = &FileStore{}
	flag.StringVar(&fileStore.FileName, "f", parseFileStoragePath(), "path to file store")
	db, _ = fileStore.LoadFromDisk()
}

func parseFileStoragePath() string {
	name, ok := os.LookupEnv(FileStorePathEnvName)
	if !ok {
		return "/tmp/godb" // default file storage path
	}
	return name
}

func Put(key, value string) {
	db[key] = value
	fileStore.SaveOnDisk(db)
}

func Get(key string) (string, bool) {
	v, ok := db[key]
	return v, ok
}

func (fs *FileStore) SaveOnDisk(db Table) error {
	file, err := os.OpenFile(fs.FileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(db)
}

func (fs *FileStore) LoadFromDisk() (Table, error) {
	db := make(Table)
	file, err := os.OpenFile(fs.FileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return db, err
	}
	defer file.Close()
	err = gob.NewDecoder(file).Decode(&db)
	if err != nil {
		return make(Table), err
	}
	return db, nil
}
