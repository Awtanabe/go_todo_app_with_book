package entity

import "time"

// int64だと分かりずらいから TaskID?
type TaskID int64
type TaskStatus string

// 定数定義  TaskStatusはstringか
const (
	TaskStatusTodo   TaskStatus = "todo"
	TaskStatusDosing TaskStatus = "doing"
	TaskTstatusDone  TaskStatus = "done"
)

// 単体
type Task struct {
	ID      TaskID     `json:"id"`
	Title   string     `json:"title"`
	Status  TaskStatus `json:"status"`
	Created time.Time  `json:"created"`
}

// 配列
// time がバイト数大きい見たい だからTaskはポインタなんだろうな
type Tasks []Task

func (o Task) TableName() string {
	return "todo.task"
}
