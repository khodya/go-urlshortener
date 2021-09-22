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

var fileStore *FileStore

func init() {
	fileStore = &FileStore{}
	flag.StringVar(&fileStore.fileName, "f", parseFileStoragePath(), "path to file store")
	fileStore.LoadFromDisk()
}

func parseFileStoragePath() string {
	name, ok := os.LookupEnv(FileStorePathEnvName)
	if !ok {
		return "/tmp/godb" // default file storage path
	}
	return name
}

func Put(key, value string) {
	fileStore.Links[key] = value
	fileStore.SaveOnDisk()
}

func Get(key string) (string, bool) {
	v, ok := fileStore.Links[key]
	return v, ok
}

func (fs *FileStore) SaveOnDisk() error {
	file, err := os.OpenFile(fs.fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(*fs)
}

func (fs *FileStore) LoadFromDisk() error {
	file, err := os.OpenFile(fs.fileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	err = gob.NewDecoder(file).Decode(fs)
	if err != nil {
		return err
	}
	return nil
}
