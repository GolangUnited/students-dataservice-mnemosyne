package file

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/NEKETSKY/mnemosyne/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type StorageFile struct {
	uploadFolder string
	mutex        *sync.RWMutex
}

func NewStorageFile() *StorageFile {
	rel := ""
	if err := godotenv.Load(".env"); err != nil {
		rel = "./../../"
		if err = godotenv.Load(rel + ".env"); err != nil {
			logger.Fatalf("error loading env variables: %s", err.Error())
		}
	}

	return &StorageFile{
		uploadFolder: rel + os.Getenv("UPLOAD_FOLDER") + "/",
		mutex:        &sync.RWMutex{},
	}
}

func (sf *StorageFile) Save(filename string, content []byte) (path string, err error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// calc hash from content
	h := md5.New()
	_, err = h.Write(content)
	if err != nil {
		return "", errors.Wrap(err, "save file md5")
	}
	hashRunes := []rune(hex.EncodeToString(h.Sum(nil)))
	hashFilename := string(hashRunes[0:2]) + "/"

	// create hash sub folder
	err = os.MkdirAll(sf.uploadFolder+hashFilename, os.ModeDir)
	if err != nil {
		return "", errors.Wrap(err, "save file mkdir")
	}

	// find files with same name
	pathUploadFile := sf.uploadFolder + hashFilename + filename
	_, err = os.Stat(pathUploadFile)
	for prefix := 1; !errors.Is(err, os.ErrNotExist); prefix++ {
		fileExt := filepath.Ext(filename)
		fileBaseName := strings.TrimSuffix(filename, fileExt)
		pathUploadFile = sf.uploadFolder + hashFilename + fileBaseName + "_" + strconv.Itoa(prefix) + fileExt
		_, err = os.Stat(pathUploadFile)
	}

	// save file
	err = os.WriteFile(pathUploadFile, content, os.ModeDevice)
	if err != nil {
		return "", errors.Wrap(err, "save file write")
	}

	return pathUploadFile, nil
}

func (sf *StorageFile) Delete(path string) (err error) {
	sf.mutex.Lock()

	err = os.Remove(path)
	if err != nil {
		sf.mutex.Unlock()
		return errors.Wrap(err, "delete file")
	}

	sf.mutex.Unlock()
	go sf.cleanDir(path)

	return nil
}

func (sf *StorageFile) cleanDir(path string) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	dir := filepath.Dir(path)
	entries, err := os.ReadDir(dir)
	if err != nil {
		logger.Infof("clean dir read error:%s", err.Error())
	}

	if len(entries) == 0 {
		err = os.Remove(dir)
		if err != nil {
			logger.Infof("clean dir remove error:%s", err.Error())
		}
	}
}
