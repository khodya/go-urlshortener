package storage

import (
	"database/sql"
	"encoding/gob"
	"flag"
	"log"
	"os"

	_ "github.com/lib/pq"
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
	DBConnectionString   = "postgresql://localhost"
)

var fileStore *FileStore
var databaseDSN string = ""

func init() {
	fileStore = newFileStore()
	flag.StringVar(&fileStore.fileName, "f", parseFileStoragePath(), "path to file store")
	flag.StringVar(&databaseDSN, "d", parseDatabaseDSN(), "postgress connection string")
	log.Println("connection string:", databaseDSN)
	fileStore.LoadFromDisk()
}

func parseFileStoragePath() string {
	name, ok := os.LookupEnv(FileStorePathEnvName)
	if !ok {
		return "/tmp/godb" // default file storage path
	}
	return name
}

func parseDatabaseDSN() string {
	value, ok := os.LookupEnv("DATABASE_DSN")
	if !ok {
		return DBConnectionString
	}
	return value
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
	fileStore.SaveOnDisk()
}

func GetUser(userID string) ([]string, bool) {
	links, ok := fileStore.Users[userID]
	return links, ok
}

func PingDB() bool {
	db, err := sql.Open("postgres", databaseDSN)
	if err != nil {
		log.Println("Failed DB ping")
	}
	defer db.Close()
	return err == nil
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
