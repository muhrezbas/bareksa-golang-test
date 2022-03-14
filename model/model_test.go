package model

import (
	// "bareksa-api/model"
	"reflect"

	"fmt"
	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type TopicRepositoryMock struct {
	mock.Mock
}

func (r TopicRepositoryMock) FindTopic() ([]Topic, error) {
	args := r.Called()
	topics := []Topic{
		{Name: "test"},
	}
	return topics, args.Error(1)
}
func (r TopicRepositoryMock) FindTopicByID(id string) (Topic, error) {
	args := r.Called()
	fmt.Println(id)
	topics := Topic{Name: "test"}
	return topics, args.Error(1)
}
func (r TopicRepositoryMock) CreateTopic(body Topic) (Topic, error) {
	args := r.Called()
	topics := Topic{Name: "test"}
	return topics, args.Error(1)
}
func (r TopicRepositoryMock) UpdateTopic(body Topic, id string) (Topic, error) {
	args := r.Called()
	topics := Topic{}
	return topics, args.Error(1)
}

func TestService_FindTopic(t *testing.T) {
	repository := &TopicRepositoryMock{}
	repository.On("FindTopic").Return([]Topic{}, nil)
	service := TopicService{repository}
	users, _ := service.FindTopic()
	for i := range users {
		expectation := true
		test := false
		zt := reflect.TypeOf(users[i].Name).Kind()
		if zt == reflect.String {
			test = true
		}
		if test != expectation {
			t.Errorf("Expected %v but got %v", expectation, test)
		}
	}
}
func TestService_FindTopicByID(t *testing.T) {
	repository := &TopicRepositoryMock{}
	repository.On("FindTopicByID").Return([]Topic{}, nil)
	service := TopicService{repository}
	users, _ := service.FindTopicByID("1")

	expectation := true
	test := false
	zt := reflect.TypeOf(users.Name).Kind()
	if zt == reflect.String {
		test = true
	}
	if test != expectation {
		t.Errorf("Expected %v but got %v", expectation, test)
	}

}
