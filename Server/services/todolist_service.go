package services

import (
	"Chat/Server/api"
	interfaces "Chat/Server/interfaces"
)

type TodoListService struct {
	database interfaces.TodoListServiceDataBase
}

func NewTodoListService(database interfaces.TodoListServiceDataBase) *TodoListService {
	return &TodoListService{
		database: database,
	}
}

func (todo *TodoListService) AddTask(_ api.ServerContext, userID string, listID string, task api.Task) api.ResponseStatus {
	todo.database.AddTask(userID, listID, task)
	return api.OK
}

func (todo *TodoListService) GetTasks(_ api.ServerContext, userID, listID string) []api.Task {
	return todo.database.GetListTasks(userID, listID)
}

func (todo *TodoListService) GetLists(_ api.ServerContext, userID string) []api.List {
	return todo.database.GetLists(userID)
}

func (todo *TodoListService) AddList(_ api.ServerContext, userID string, list api.List) api.ResponseStatus {
	todo.database.AddList(userID, list)
	return api.OK
}

func (todo *TodoListService) SetTaskState(api.ServerContext, string, string, string, bool) api.ResponseStatus {
	return api.OK
}
