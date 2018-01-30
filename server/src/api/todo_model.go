package api

import (
	"context"
	"errors"
	"strings"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/appengine/log"
)

// "projects/your-project-id/instances/your-instance-id/databases/your-database-id"
const databaseName = "projects/sandbox-mstssk/instances/test-todo/databases/test-todo"

// Todo is todo
type Todo struct {
	TodoID  string
	Title   string
	DueDate spanner.NullDate // YYYY-MM-DD
	Done    bool
}

/*
CREATE TABLE Todo (
	TodoID STRING(MAX) NOT NULL,
	Done BOOL,
	DueDate DATE,
	Title STRING(MAX) NOT NULL,
) PRIMARY KEY (TodoID)
*/

// TodoStore manages Todo CRUD operation.
type TodoStore struct{}

// Insert Todo
func (s TodoStore) Insert(c context.Context, todo *Todo) (*Todo, error) {

	if todo.TodoID != "" {
		return nil, errors.New("Shoud not set TodoID")
	}

	// ナノ秒までのタイムスタンプを逆転させた文字列をIDにする
	todo.TodoID = Reverse(strings.Replace(time.Now().Format("20060102150405.000000000"), ".", "", 1))

	start := time.Now()

	m, err := spanner.InsertStruct("Todo", todo)
	if err != nil {
		return nil, err
	}
	client, err := spanner.NewClient(c, databaseName)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	ts, err := client.Apply(c, []*spanner.Mutation{m})
	if err != nil {
		return nil, err
	}
	log.Debugf(c, "commitTimestamp: %v", ts)
	log.Debugf(c, "%.3fs", time.Now().Sub(start).Seconds())

	return todo, nil
}

// Get Todo
func (s TodoStore) Get(c context.Context, id int64) (Todo, error) {
	return Todo{}, nil
}

// List Todo
func (s TodoStore) List(c context.Context) ([]Todo, error) {
	return nil, nil
}
