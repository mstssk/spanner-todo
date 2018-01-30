package api

import (
	"context"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/appengine/log"
)

// "projects/your-project-id/instances/your-instance-id/databases/your-database-id"
const databaseName = "projects/sandbox-mstssk/instances/test-todo/databases/test-todo"

// Todo is todo
type Todo struct {
	TodoID  int64
	Title   string
	Done    bool
	DueDate time.Time
}

// TodoStore manages Todo CRUD operation.
type TodoStore struct{}

// Insert Todo
func (s TodoStore) Insert(c context.Context, todo Todo) (Todo, error) {

	client, err := spanner.NewClient(c, databaseName)
	if err != nil {
		return Todo{}, err
	}
	defer client.Close()

	stmt := spanner.Statement{SQL: "SELECT 1"}
	iter := client.Single().Query(c, stmt)
	defer iter.Stop()

	row, err := iter.Next()
	if err != nil {
		log.Errorf(c, "Query failed with %v", err.Error())
		return Todo{}, err
	}

	var i int64
	if row.Columns(&i) != nil {
		log.Errorf(c, "Failed to parse row %v", err.Error())
		return Todo{}, err
	}
	log.Infof(c, "Got value %v\n", i)

	return Todo{}, nil
}

// Get Todo
func (s TodoStore) Get(c context.Context, id int64) (Todo, error) {
	return Todo{}, nil
}

// List Todo
func (s TodoStore) List(c context.Context) ([]Todo, error) {
	return nil, nil
}
