package service

import (
	"os"
	"path/filepath"
)

type Storage interface {
	WriteFile(filePath string, data []byte) error
	ReadFile(filePath string) ([]byte, error)
}

type localStorage struct {
}

func NewStorage() Storage {
	return &localStorage{}
}

func (s *localStorage) WriteFile(filePath string, data []byte) error {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, 0766)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (s *localStorage) ReadFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
