package file

var (
	storage SaveDeleter
)

func init() {
	storage = NewStorageFile()
}

type SaveDeleter interface {
	Save(filename string, content []byte) (path string, err error)
	Delete(path string) error
}

func Save(filename string, content []byte) (path string, err error) {
	return storage.Save(filename, content)
}

func Delete(path string) (err error) {
	return storage.Delete(path)
}
