package services

import "Chat/Server/api"

type TodoListService struct{}

func NewTodoListService() *TodoListService {
	return &TodoListService{}
}

func (todo *TodoListService) AddTask(api.ServerContext, string, string, api.Task) api.ResponseStatus {
	return api.OK
}

func (todo *TodoListService) GetTasks(api.ServerContext, string, string) []api.Task {
	return make([]api.Task, 0)
}

func (todo *TodoListService) GetLists(api.ServerContext, string) []api.List {
	return make([]api.List, 0)
}

func (todo *TodoListService) AddList(api.ServerContext, string, api.List) api.ResponseStatus {
	return api.OK
}

func (todo *TodoListService) SetTaskState(api.ServerContext, string, string, string, bool) api.ResponseStatus {
	return api.OK
}
