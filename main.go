package main

import (
	"flag"
	"fmt"
	"foobar/domain"
	"strconv"

	"log"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	age := flag.Arg(1)
	ageInt, _ := strconv.Atoi(age)

	useCase := &domain.Usecase{}
	useCase.Client = &domain.APIClient{}
	userInfo := domain.UserInfo{
		Name: name,
		Age:  ageInt,
	}
	if err := useCase.Client.SetInfo(userInfo); err != nil {
		log.Printf("%v", err)
	}
	userInfoSlice := useCase.Client.GetInfo(name)
	fmt.Printf("%+v\n", userInfoSlice)
}
