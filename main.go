package main

import (
	"flag"
	"foobar/domain"
	"strconv"

	"log"
)

func run(apiClient *domain.APIClient) {
	flag.Parse()
	name := flag.Arg(0)
	age := flag.Arg(1)
	ageInt, _ := strconv.Atoi(age)

	useCase := &domain.Usecase{}
	useCase.Client = apiClient
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
	run(&domain.APIClient{})
}
