package storage

import (
	"encoding/gob"
	"flag"
	"log"
	"os"
)

type Table map[string]string

type Store struct {
	Links Table
	Users map[string][]string
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
	fileStore = newFileStore()
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

func Put(shortURLPath, originalURL string) {
	fileStore.Links[shortURLPath] = originalURL
	fileStore.SaveOnDisk()
}

func Get(shortURLPath string) (string, bool) {
	v, ok := fileStore.Links[shortURLPath]
	return v, ok
}

func PutUser(user, shortURLPath string) {
	slice, ok := fileStore.Users[user]
	if !ok {
		fileStore.Users[user] = make([]string, 0)
	}
	fileStore.Users[user] = append(slice, shortURLPath)
}

func newFileStore() *FileStore {
	fs := new(FileStore)
	fs.Links = make(Table)
	fs.Users = make(map[string][]string)
	return fs
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
		log.Println("Error while reading filestore from file")
		return err
	}
	return nil
}
