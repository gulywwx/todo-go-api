package models

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Task struct {
	ID      string     `json:"id,omitempty" `
	Title   string     `json:"title,omitempty" `
	Content string     `json:"content,omitempty" `
	Status  int        `json:"status" `
	IsMajor bool       `json:"is_major" `
	Created *time.Time `json:"created,omitempty"`
	Expired *time.Time `json:"expired,omitempty"`
}

type Entry struct {
	key   string
	value *Task
}

type byID []Entry

func (d byID) Len() int {
	return len(d)
}
func (d byID) Less(i, j int) bool {
	a, _ := strconv.Atoi(d[i].value.ID)
	b, _ := strconv.Atoi(d[j].value.ID)
	return a < b
}
func (d byID) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

const MAX_NUM int = 20

var (
	tasklist     = make(map[string]*Task)
	seq      int = MAX_NUM
)

func init() {
	now := time.Now()
	nextMonth := now.AddDate(0, 1, 0)
	for i := 0; i < MAX_NUM; i++ {
		tmp := fmt.Sprintf("%d", i)
		tasklist[tmp] = &Task{ID: tmp, Title: "title" + tmp, Content: "test" + tmp, Status: rand.Intn(2), IsMajor: true, Created: &now, Expired: &nextMonth}
	}
	seq = MAX_NUM

}

func GetTaskList(status int) ([]*Task, error) {
	values := []*Task{}

	slice := make(byID, 0, len(tasklist))
	for key, value := range tasklist {
		slice = append(slice, Entry{key, value})
	}

	// Sort the slice.
	sort.Sort(slice)

	if status == -1 {
		for _, entry := range slice {
			values = append(values, entry.value)
		}

	} else {
		for _, entry := range slice {
			if entry.value.Status == status {
				values = append(values, entry.value)
			}

		}

	}

	return values, nil
}

func GetTask(id string) (*Task, error) {
	task, found := tasklist[id]

	if !found {
		return nil, errors.New("task not found")
	}

	return task, nil
}

func UpdateMark(id string, isMajor bool) (*Task, error) {
	_, err := GetTask(id)

	if err != nil {
		return nil, err
	}

	tasklist[id].IsMajor = isMajor

	return tasklist[id], nil

}

func UpdateTaskStatus(id string, status int) error {
	_, err := GetTask(id)

	if err != nil {
		return err
	}

	tasklist[id].Status = status
	return nil
}

func AddTask(title string, content string, expired time.Time) error {
	seq += 1
	task := Task{
		ID:      fmt.Sprintf("%d", seq),
		Title:   title,
		Content: content,
		Expired: &expired,
		Status:  1,
	}
	tasklist[task.ID] = &task

	return nil
}

func EditTask(id string, title string, content string, expired time.Time) error {
	_, err := GetTask(id)

	if err != nil {
		return err
	}

	tasklist[id].Title = title
	tasklist[id].Content = content
	tasklist[id].Expired = &expired

	return nil
}

func DeleteTask(ID string) error {
	delete(tasklist, ID)
	return nil
}
