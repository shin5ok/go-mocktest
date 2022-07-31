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
	SetInfo() error
	GetInfo(string) []UserInfo
}

type UserInfo struct {
	Name string
	Age  int16
}

var projectId = os.Getenv("PROJECT_ID")

func (u UserInfo) SetInfo() error {
	ctx := context.Background()
	client, _ := firestore.NewClient(ctx, projectId)
	id := uuid.NewUUID()
	_, err := client.Collection("user").Doc(string(id)).Set(ctx, map[string]interface{}{
		"name": u.Name,
		"age":  u.Age,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		return fmt.Errorf("An error has occurred: %w", err)
	}
	return nil
}

func (u UserInfo) GetInfo(name string) []UserInfo {
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
		// docData = append(docData, UserInfo{Name: data["name"].(string), Age: data["age"]})
		fmt.Printf("%#v\n", data)
		ageInt := int16(data["age"].(int64))
		docData = append(docData,
			UserInfo{
				Name: data["name"].(string),
				Age:  ageInt,
			})
	}

	return docData

}
