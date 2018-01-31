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

	info = swagger.NewHandlerInfo(s.Get)
	ucon.Handle("GET", "/api/todo/{id}", info)
	info.Summary, info.Tags = "TODOを取得する", []string{tag.Name}
}

type todoService struct{}

func (s *todoService) Get(c context.Context, req *StringIDReq) (*Todo, error) {
	var store *TodoStore
	todo, err := store.Get(c, req.ID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
