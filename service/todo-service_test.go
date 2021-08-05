package service

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/palash287gupta/golang-mux-api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(todo *model.Todo) (*model.Todo, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.Todo), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]model.Todo, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Todo), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var identifier int64 = 1

	todo := model.Todo{ID: identifier, Title: "A", Text: "B"}
	// Setup expectations
	mockRepo.On("FindAll").Return([]model.Todo{todo}, nil)

	testService := NewTodoService(mockRepo)

	result, _ := testService.FindAll()
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	todo := model.Todo{Title: "A", Text: "B"}

	//Setup expectations
	mockRepo.On("Save").Return(&todo, nil)

	testService := NewTodoService(mockRepo)

	result, err := testService.Create(&todo)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)
}

func TestValidateEmptyTodo(t *testing.T) {
	testService := NewTodoService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "todo is empty", err.Error())
}

func TestValidateEmptyTodoTitle(t *testing.T) {
	todo := model.Todo{ID: 1, Title: "", Text: "B"}

	testService := NewTodoService(nil)

	err := testService.Validate(&todo)

	assert.NotNil(t, err)
	assert.Equal(t, "todo title is empty", err.Error())
}

func TestSaveTodo(t *testing.T) {

	tests := []struct {
		testName string
		title    string
		text     string
		expected *model.Todo
	}{
		{
			"success",
			"title 1",
			"text 1",
			&model.Todo{
				Title: "title 1",
				Text:  "text 1",
			},
		}, {
			"failure",
			"title2 ",
			"text 2",
			&model.Todo{
				Title: "title 2",
				Text:  "text 2",
			},
		},
	}

	mockRepo := new(MockRepository)
	testService := NewTodoService(mockRepo)
	var todo *model.Todo

	comparer := cmp.Comparer(func(x, y *model.Todo) bool {
		return x.Title == y.Title && x.Text == y.Text
	})

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			mockRepo.On("Save").Return(test.expected, nil)
			todo = &model.Todo{
				Title: test.title,
				Text:  test.text,
			}
			result, _ := testService.Create(todo)
			mockRepo.AssertExpectations(t)
			if diff := cmp.Diff(test.expected, result, comparer); diff != "" {
				t.Error(diff)
			}
		})
	}
}
