package main

import (
	"flag"
	"fmt"
	"mock_test/domain"
	"strconv"

	"log"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	age := flag.Arg(1)
	ageInt, _ := strconv.Atoi(age)
	u := domain.UserInfo{Name: name, Age: int16(ageInt)}
	if err := u.SetInfo(); err != nil {
		log.Printf("%v", err)
	}
	userInfo := u.GetInfo(name)
	fmt.Printf("%+v", userInfo)
}
