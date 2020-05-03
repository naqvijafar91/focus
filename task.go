package focus

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const layout = "2006-01-02"

// Time is a wrapper over time.Time so that this field can be parsed using custom logic
type Time struct {
	time.Time
}

// Value tells the database how to store this type
func (t *Time) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	var tempTime time.Time
	tempTime = *&t.Time
	return tempTime.Local(), nil
}

// func (t *Time) Value() (driver.Value, error) {
// 	if t == nil {
// 		return nil, nil
// 	}
// 	i := t.Time.UnixNano()
// 	return json.Marshal(i)
// }

// Scan is used by the database to initialize a struct from the value stored in DB
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	xType := fmt.Sprintf("%T", value)
	switch value.(type) {
	case time.Time:
		*&t.Time = value.(time.Time).Local()
	default:
		fmt.Println("Invalid type found ", xType)
	}

	return nil
}

// func (t *Time) Scan(value interface{}) error {
// 	if value == nil {
// 		return nil
// 	}
// 	var i int64
// 	if err := json.Unmarshal(value.([]byte), &i); err != nil {
// 		return err
// 	}
// 	t.Time = time.Unix(0, i)
// 	return nil
// }

// MarshalJSON converts struct into json string
func (t Time) MarshalJSON() ([]byte, error) {
	// ""2020-05-02T00:00:00""
	formatted := fmt.Sprintf("%02d-%02d-%d",
		t.Day(), t.Month(), t.Year())
	return json.Marshal(formatted)
}

// UnmarshalJSON converts string into serialized object
func (t *Time) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	str := string(data)
	if len(strings.Trim(str, " ")) == 0 {
		return nil
	}
	str = strings.Trim(str, " \"")
	time, err := time.Parse("02-01-2006", str)
	if err != nil {
		return err
	}
	t.Time = time
	return nil
}

type Task struct {
	ID            string `json:"id"`
	Description   string `json:"description"`
	FolderID      string `json:"folder_id"`
	DueDate       *Time  `json:"due_date"`
	CompletedDate *Time  `json:"completed_date"`
}

type TaskService interface {
	Create(task *Task) (*Task, error)
	GetAll() ([]*Task, error)
	GetAllByFolderID(folderID string) ([]*Task, error)
	MarkAsComplete(taskID string) (*Task, error)
	Update(task *Task) (*Task, error)
}
