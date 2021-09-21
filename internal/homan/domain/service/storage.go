package service

import (
	"os"
	"path/filepath"
)

type Storage interface {
	WriteFile(filePath string, data []byte) error
	ReadFile(filePath string) ([]byte, error)
	Mkdir(path string) error
}

type localStorage struct {
}

func NewStorage() Storage {
	return &localStorage{}
}

func (s *localStorage) WriteFile(filePath string, data []byte) error {
	dir := filepath.Dir(filePath)
	err := s.Mkdir(dir)
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

func (s *localStorage) Mkdir(path string) error {
	err := os.MkdirAll(path, 0766)
	if err != nil {
		return err
	}
	return nil
}
