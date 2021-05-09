package taskrepository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ocsen-hoc-code/go-auj/models/config"
	"github.com/ocsen-hoc-code/go-auj/models/task"
)

type TaskRepository struct {
	dbConfig *config.DbConfig
}

func NewTaskRepository(dbConfigInject *config.DbConfig) *TaskRepository {
	return &TaskRepository{dbConfig: dbConfigInject}
}

// FindTasks returns tasks if match userID AND createDate.
func (rp TaskRepository) FindTasks(ctx context.Context, userID sql.NullString, createdDate time.Time) ([]*task.Task, error) {

	var rows *sql.Rows
	var err error
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND DATE(created_date) = $2`
	if createdDate.IsZero() {
		stmt = `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1`
		rows, err = rp.dbConfig.Database.QueryContext(ctx, stmt, userID)
	} else {
		rows, err = rp.dbConfig.Database.QueryContext(ctx, stmt, userID, createdDate.Format("2006-01-02"))
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*task.Task
	for rows.Next() {
		t := &task.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// CreateTask create a new task to DB
func (rp TaskRepository) CreateTask(ctx context.Context, t *task.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := rp.dbConfig.Database.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateTask returns boolend And max_todo if match userId AND date
func (rp TaskRepository) ValidateTask(ctx context.Context, userID string, now time.Time) (bool, int) {
	stmt := `SELECT count(id) FROM tasks WHERE user_id = $1 AND created_date = $2`
	row := rp.dbConfig.Database.QueryRowContext(ctx, stmt, userID, now.Format("2006-01-02"))
	// stmt := `SELECT count(id) FROM tasks WHERE user_id = $1`
	// row := rp.dbConfig.Database.QueryRowContext(ctx, stmt, userID)
	limitTask := 0
	err := row.Scan(&limitTask)
	if err != nil {
		return true, 0
	}

	maxToDo := 0
	stmt = `SELECT max_todo FROM users WHERE users.id = $1`
	row = rp.dbConfig.Database.QueryRowContext(ctx, stmt, userID)
	err = row.Scan(&maxToDo)
	if err != nil {
		return true, 0
	}

	return limitTask >= maxToDo, maxToDo
}
