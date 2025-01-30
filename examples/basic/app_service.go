package main

type AppService struct {}

func NewAppService() *AppService {
	return &AppService{}
}

func (s *AppService) GetHello() string {
	return "Hello, World!"
}
