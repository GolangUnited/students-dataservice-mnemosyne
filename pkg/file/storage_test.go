//go:build integration

package file

import (
	"os"
	"path/filepath"
	"testing"
)

var sFile *StorageFile

func TestMain(m *testing.M) {
	sFile = NewStorageFile()
	code := m.Run()
	os.Exit(code)
}

func TestStorageFile_Save(t *testing.T) {
	testContent := "test content"

	tests := []struct {
		name     string
		content  []byte
		filename string
		wantPath string
		wantErr  bool
	}{
		{
			name:     "save txt file",
			content:  []byte(testContent),
			filename: "test.txt",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := sFile.Save(tt.filename, tt.content)
			defer os.Remove(filepath.Dir(path))
			defer os.Remove(path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestStorageFile_Delete(t *testing.T) {
	testContent := "test content"

	tests := []struct {
		name     string
		content  []byte
		filename string
		wantErr  bool
	}{
		{
			name:     "delete txt file",
			content:  []byte(testContent),
			filename: "test.txt",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, _ := sFile.Save(tt.filename, tt.content)
			err := sFile.Delete(path)
			defer os.Remove(filepath.Dir(path))
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
