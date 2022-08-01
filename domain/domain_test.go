package domain_test

import (
	"foobar/domain"
	mock_domain "foobar/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserInfo_SetInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_domain.NewMockUserInterface(ctrl)
	mockClient.EXPECT().SetInfo(
		domain.UserInfo{
			Name: "foo",
			Age:  100,
		},
	).Return(nil)

	func(t *testing.T, client *mock_domain.MockUserInterface) {
		t.Helper()

		useCase := &domain.Usecase{}
		useCase.Client = client

		err := useCase.Client.SetInfo(domain.UserInfo{
			Name: "foo",
			Age:  100,
		})

		if err != nil {
			t.Error(err)
		}

	}(t, mockClient)

}
