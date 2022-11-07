package main

import (
	"flag"
	"foobar/domain"
	"strconv"

	"log"
)

type APIClienter interface {
	Get() *domain.APIClient
}

type Config struct {
	api *domain.APIClient
}

func (c Config) Get() *domain.APIClient {
	return c.api
}

type RunConfig struct {
	api APIClienter
}

func (v RunConfig) run() {
	flag.Parse()
	name := flag.Arg(0)
	age := flag.Arg(1)
	ageInt, _ := strconv.Atoi(age)

	useCase := &domain.Usecase{}
	useCase.Client = v.api.Get()
	userInfo := domain.UserInfo{
		Name: name,
		Age:  ageInt,
	}
	if err := useCase.Client.SetInfo(userInfo); err != nil {
		log.Printf("%v", err)
	}
	userInfoSlice := useCase.Client.GetInfo(name)
	log.Printf("%+v\n", userInfoSlice)
}

func main() {
	config := Config{&domain.APIClient{}}
	r := RunConfig{api: config}

	r.run()
}
