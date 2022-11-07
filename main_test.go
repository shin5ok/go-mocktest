package main

import (
	"foobar/domain"
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

	err := useCase.Client.SetInfo(domain.UserInfo{Name: "name", Age: 10})
	assert.Nil(t, err)

	userInfoSlice := useCase.Client.GetInfo(name)
	assert.Equal(t, userInfoSlice[0].Name, name)
}
