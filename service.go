package main

type UserService struct {
	db Database
}

func (s *UserService) CreateUser(name string) error {
	return s.db.Insert(name)
}
