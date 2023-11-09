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

func (todo *TodoListService) AddTask(_ api.ServerContext, userID string, listID string, task api.Task) api.TodoListReponse {
	task.Id = userID + listID + task.Description
	todo.database.AddTask(userID, listID, task)
	return api.TodoListReponse{
		Status: api.OK,
		Id:     userID + listID + task.Description, //TODO: make id generation
	}
}

func (todo *TodoListService) GetTasks(_ api.ServerContext, userID, listID string) []api.Task {
	return todo.database.GetListTasks(userID, listID)
}

func (todo *TodoListService) GetLists(_ api.ServerContext, userID string) []api.List {
	return todo.database.GetLists(userID)
}

func (todo *TodoListService) AddList(_ api.ServerContext, userID string, list api.List) api.TodoListReponse {
	list.Id = userID + list.Title
	todo.database.AddList(userID, list)
	return api.TodoListReponse{
		Status: api.OK,
		Id:     userID + list.Title,
	}
}

func (todo *TodoListService) SetTaskState(_ api.ServerContext, userID, listID, taskID string, state bool) api.ResponseStatus {
	todo.database.SetTaskState(userID, listID, taskID, state)
	return api.OK
}
