package helper

import (
	"fmt"
	"gofiber_restapi/model/domain"
	"gofiber_restapi/model/web"
)

func ToTodoResponse(todo domain.Todo) web.TodoResponse {
	return web.TodoResponse{
		Id:     todo.Id,
		Title:  todo.Title,
		Todos:  todo.Todos,
		Status: fmt.Sprintf("%d", todo.Status),
	}
}

func ToTodoResponses(categories []domain.Todo) []web.TodoResponse {
	var todoResponses []web.TodoResponse
	for _, todo := range categories {
		todoResponses = append(todoResponses, ToTodoResponse(todo))
	}
	return todoResponses
}
