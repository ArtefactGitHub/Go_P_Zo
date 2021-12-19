package test

import (
	"log"
)

const TimeLayout = "2006-01-02"

type TrashScanner struct{}

func (TrashScanner) Scan(interface{}) error {
	return nil
}

func Failuer(err error) {
	log.Panicf("failuer: %v", err)
}
