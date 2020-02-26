package model

import (
	"fmt"
)

type Task struct {
	UID  int    `json:"uid"`
	ID   int    `json:"id" gorm:"primary_key"`
	Menu string `json:"menu"`
	Rep  int    `json:"rep"`
	Set  int    `json:"set"`
}

type Tasks []Task

func CreateTask(task *Task) {
	db.Create(task)
}

func FindTasks(t *Task) Tasks {
	var tasks Tasks
	db.Where(t).Find(&tasks)
	return tasks
}

func DeleteTask(t *Task) error {
	if rows := db.Where(t).Delete(&Task{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Task (%v) to delete", t)
	}
	return nil
}

func UpdateTask(t *Task) error {
	rows := db.Model(t).Update(map[string]interface{}{
		"menu": t.Menu,
		"set":  t.Set,
		"rep":  t.Rep,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Task (%v) to update", t)
	}
	return nil
}
