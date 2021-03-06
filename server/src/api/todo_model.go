package api

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/appengine/log"
)

// "projects/your-project-id/instances/your-instance-id/databases/your-database-id"
const databaseName = "projects/sandbox-mstssk/instances/test-todo/databases/test-todo"

// Todo is todo
// +jwg
type Todo struct {
	TodoID  int64
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
func (s *TodoStore) Insert(c context.Context, todo *Todo) (*Todo, error) {
	if todo.TodoID != 0 {
		return nil, errors.New("Shoud not set TodoID")
	}
	todo.TodoID = s.generateID(time.Now())

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

func (s *TodoStore) generateID(seed time.Time) int64 {
	// https://cloud.google.com/spanner/docs/whitepapers/optimizing-schema-design#anti-pattern_sequences
	return rand.NewSource(seed.UnixNano()).Int63()
}

// Update Todo
func (s *TodoStore) Update(c context.Context, todo *Todo) (*Todo, error) {
	if todo.TodoID == 0 {
		return nil, errors.New("Shoud set TodoID")
	}

	start := time.Now()
	m, err := spanner.UpdateStruct("Todo", todo)
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
func (s *TodoStore) Get(c context.Context, id int64) (*Todo, error) {
	client, err := spanner.NewClient(c, databaseName)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	name, fields := GetStructFieldNames(Todo{})
	row, err := client.Single().ReadRow(c, name, spanner.Key{id}, fields)
	if err != nil {
		return nil, err
	}
	todo := &Todo{}
	err = row.ToStruct(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// List Todo
func (s *TodoStore) List(c context.Context) ([]*Todo, error) {
	client, err := spanner.NewClient(c, databaseName)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// FIXME 雑実装

	name, _ := GetStructFieldNames(Todo{})
	stmt := spanner.NewStatement(fmt.Sprintf("SELECT * FROM %s", name))
	iter := client.ReadOnlyTransaction().Query(c, stmt)
	todos := make([]*Todo, 0)
	err = iter.Do(func(row *spanner.Row) error {
		todo := &Todo{}
		err := row.ToStruct(todo)
		if err != nil {
			return err
		}
		todos = append(todos, todo)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return todos, nil
}
