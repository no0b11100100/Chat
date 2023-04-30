package services

import "Chat/Server/api"

type TodoListService struct{}

func NewTodoListService() *TodoListService {
	return &TodoListService{}
}

func (todo *TodoListService) AddTask(api.ServerContext, string, api.Task) api.ResponseStatus {
	return api.OK
}

func (todo *TodoListService) GetTasks(api.ServerContext, string) []api.Task {
	return make([]api.Task, 0)
}
