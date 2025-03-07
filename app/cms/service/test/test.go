package test

import "gorm.io/gorm"

type TestService interface {
	Test() string
}

type testService struct {
	db *gorm.DB
}

func (t testService) Test() string {
	return "hello"
}

func NewTestService(db *gorm.DB) TestService {
	return &testService{db: db}
}
