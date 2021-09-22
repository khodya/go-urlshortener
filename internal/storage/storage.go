package storage

import (
	"encoding/gob"
	"flag"
	"os"
)

type Table map[string]string

type Store struct {
	Links Table
}

type FileStore struct {
	Store
	fileName string
}

const (
	FileStorePathEnvName = "FILE_STORAGE_PATH"
)

var db Table
var fileStore *FileStore

func init() {
	fileName := flag.String("f", parseFileStoragePath(), "path to file store")
	fileStore = &FileStore{fileName: *fileName}
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
	file, err := os.OpenFile(fs.fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(db)
}

func (fs *FileStore) LoadFromDisk() (Table, error) {
	db := make(Table)
	file, err := os.OpenFile(fs.fileName, os.O_RDONLY|os.O_CREATE, 0777)
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
