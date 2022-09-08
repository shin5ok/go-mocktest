package domain

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"k8s.io/apimachinery/pkg/util/uuid"
)

type UserInterface interface {
	SetInfo(UserInfo) error
	GetInfo(string) []UserInfo
}

type UserInfo struct {
	Name string
	Age  int
}

type APIClient struct {
	Client UserInterface
}

type Usecase struct {
	APIClient
}

var projectId = os.Getenv("PROJECT_ID")

func (v *APIClient) SetInfo(u UserInfo) error {
	ctx := context.Background()
	//err := Client(ctx, u)
	client, _ := firestore.NewClient(ctx, projectId)
	id := uuid.NewUUID()
	_, err := client.Collection("user").Doc(string(id)).Set(ctx, map[string]interface{}{
		"name": u.Name,
		"age":  u.Age,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		return fmt.Errorf("an error has occurred: %w", err)
	}
	return nil
}

func (v *APIClient) GetInfo(name string) []UserInfo {
	ctx := context.Background()
	client, _ := firestore.NewClient(ctx, projectId)
	itr := client.Collection("user").Where("name", "==", name).Documents(ctx)
	var docData []UserInfo
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []UserInfo{}
		}
		data := doc.Data()
		// fmt.Printf("%#v\n", data)
		ageInt := data["age"].(int64)
		docData = append(docData,
			UserInfo{
				Name: data["name"].(string),
				Age:  int(ageInt),
			})
	}

	return docData

}
