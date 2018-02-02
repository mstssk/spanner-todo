package api

import (
	"context"

	"github.com/favclip/ucon"
	"github.com/favclip/ucon/swagger"
)

func todoSetup(swPlugin *swagger.Plugin) {
	s := &todoService{}

	tag := swPlugin.AddTag(&swagger.Tag{Name: "Todo", Description: "TODOを扱う"})
	var info *swagger.HandlerInfo

	info = swagger.NewHandlerInfo(s.Insert)
	ucon.Handle("POST", "/api/todo", info)
	info.Summary, info.Tags = "TODOを登録する", []string{tag.Name}

	info = swagger.NewHandlerInfo(s.Update)
	ucon.Handle("PUT", "/api/todo/{id}", info)
	info.Summary, info.Tags = "TODOを更新する", []string{tag.Name}

	info = swagger.NewHandlerInfo(s.Get)
	ucon.Handle("GET", "/api/todo/{id}", info)
	info.Summary, info.Tags = "TODOを取得する", []string{tag.Name}

	info = swagger.NewHandlerInfo(s.List)
	ucon.Handle("GET", "/api/todo", info)
	info.Summary, info.Tags = "TODOを一覧する", []string{tag.Name}
}

type todoService struct{}

func (s *todoService) Insert(c context.Context, req *TodoJSON) (*TodoJSON, error) {
	todo, err := req.Convert()
	if err != nil {
		return nil, err
	}
	var store *TodoStore
	todo, err = store.Insert(c, todo)
	if err != nil {
		return nil, err
	}
	json, err := NewTodoJSONBuilder().Convert(todo)
	if err != nil {
		return nil, err
	}
	return json, nil
}

// TodoUpdateReq 更新用reqオブジェクト
type TodoUpdateReq struct {
	ID int64 `json:"id,string" swagger:",in=path"`
	TodoJSON
}

func (s *todoService) Update(c context.Context, req *TodoUpdateReq) (*TodoJSON, error) {
	todo, err := req.Convert()
	if err != nil {
		return nil, err
	}
	var store *TodoStore
	todo, err = store.Update(c, todo)
	if err != nil {
		return nil, err
	}
	json, err := NewTodoJSONBuilder().Convert(todo)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func (s *todoService) Get(c context.Context, req *Int64IDReq) (*TodoJSON, error) {
	var store *TodoStore
	todo, err := store.Get(c, req.ID)
	if err != nil {
		return nil, err
	}
	json, err := NewTodoJSONBuilder().Convert(todo)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func (s *todoService) List(c context.Context) ([]*TodoJSON, error) {
	var store *TodoStore
	todos, err := store.List(c)
	if err != nil {
		return nil, err
	}
	json, err := NewTodoJSONBuilder().ConvertList(todos)
	if err != nil {
		return nil, err
	}
	return json, nil
}
