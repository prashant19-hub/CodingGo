package store

import "fmt"

type Store struct{}

func (s Store) SaveRecord(record interface{}) {
	//fmt.Println("Record =", record)
	s.saverecord(record)
}
func (s Store) saverecord(record interface{}) {
	fmt.Println("Record =", record)
}