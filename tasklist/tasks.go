package tasklist

import (
	"database/sql"
	"errors"
)

type Task struct {
	Id        int
	Name      string
	Status    string
	Priority  int
	CreatedAt string
	CreatedBy string
	DueDate   string
}

type TaskList struct {
	Tasks *sql.DB
}

func (tl *TaskList) Create(t Task) error {
	_, err := tl.Tasks.Exec(`insert into task_list (name, status, created_at, created_by, due_date)
	values
	($1, $2, $3, $4, $5);`, t.Id, t.Name, t.Status, t.Priority, t.CreatedAt, t.CreatedBy, t.DueDate)
	if err != nil {
		return err
	}
	return nil
}

func (tl *TaskList) Update(t Task) error {
	_, err := tl.Tasks.Exec(`update task_list where id = $1
	set name = $2, status = $3, created_at = $4, created_by = $5, 
	due_date = $6;`, t.Id, t.Name, t.Status, t.Priority, t.CreatedAt, t.CreatedBy, t.DueDate)
	if err != nil {
		return err
	}
	return nil
}

func (tl *TaskList) Get(id int) (Task, error) {
	row := tl.Tasks.QueryRow(`select * from task_list where id = $1;`, id)

	var tempTask, eTask Task
	row.Scan(&tempTask.Id, &tempTask.Name, &tempTask.Status, &tempTask.Priority, &tempTask.CreatedAt, &tempTask.CreatedBy, &tempTask.DueDate)
	if tempTask == eTask {
		return eTask, errors.New("error with such id does not exist")
	}
	return tempTask, nil
}

func (tl *TaskList) GetAll() ([]Task, error) {
	rows, err := tl.Tasks.Query(`select * from task_list;`)
	if err != nil {
		return make([]Task, 0), err
	}

	var (
		tempTask Task
		answer   []Task
	)

	for rows.Next() {
		rows.Scan(&tempTask.Id, &tempTask.Name, &tempTask.Status, &tempTask.Priority, &tempTask.CreatedAt, &tempTask.CreatedBy, &tempTask.DueDate)

		answer = append(answer, tempTask)
	}

	return answer, nil
}

func (tl *TaskList) Delete(id int) error {
	_, err := tl.Tasks.Exec(`delete from task_list where id = $1;`)
	return err
}
