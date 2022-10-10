package test_v2

import (
	"path/filepath"
	"runtime"
)

type TrashScanner struct{}

func (TrashScanner) Scan(interface{}) error {
	return nil
}

func currentDir() string {
	_, pwd, _, _ := runtime.Caller(0)
	return filepath.Dir(pwd)
}
