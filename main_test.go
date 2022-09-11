package main

import (
	"foobar/domain"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var name = "foo"

type mockAPIClient struct{}

func (v *mockAPIClient) SetInfo(u domain.UserInfo) error {
	_ = u
	return nil
}

func (v *mockAPIClient) GetInfo(name string) []domain.UserInfo {
	return []domain.UserInfo{{Name: name, Age: 10}}
}

func TestRun(t *testing.T) {

	useCase := &domain.Usecase{}
	useCase.Client = &mockAPIClient{}

	if err := useCase.Client.SetInfo(domain.UserInfo{Name: "name", Age: 10}); err != nil {
		t.Error(err)
	}
	userInfoSlice := useCase.Client.GetInfo(name)
	log.Printf("%+v", userInfoSlice)
	assert.Equal(t, userInfoSlice[0].Name, name)
}
