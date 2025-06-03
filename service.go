package main

import "log"

type UserService struct {
	db Database
}

func (s *UserService) CreateUser(name string) error {
	log.Printf("Creating user: %s", name)
	return s.db.Insert(name)
}
