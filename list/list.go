package list

import (
	"encoding/json"
	"errors"
	"time"

	log "github.com/limetext/log4go"
	"github.com/zoli/emru/list/task"
)

type (
	List struct {
		tasks     []*task.Task
		CreatedAt time.Time
	}

	Lists map[string]*List
)

var TaskNotFound = errors.New("Task not found")

func New() *List {
	return &List{CreatedAt: time.Now()}
}

func (l *List) add(t *task.Task) {
	for l.Exists(t.ID) {
		t.ID++
	}
	log.Finest("Adding task %v", *t)
	l.tasks = append(l.tasks, t)
}

func (l *List) Add(t *task.Task) task.Task {
	l.add(t)
	return *t
}

func (l *List) remove(i int) {
	log.Debug("Removing task %d: %v", i, *l.tasks[i])
	if e := len(l.tasks) - 1; i != e {
		copy(l.tasks[i:], l.tasks[i+1:])
	} else {
		l.tasks = l.tasks[:e]
	}
}

func (l *List) Remove(id int) error {
	i := l.Index(id)
	if i == -1 {
		return TaskNotFound
	}
	l.remove(i)
	return nil
}

func (l *List) update(i int, t task.Task) {
	log.Debug("Updating task %v to %v", *l.tasks[i], t)
	nt := l.tasks[i]
	nt.Title = t.Title
	nt.Body = t.Body
	nt.Done = t.Done
}

func (l *List) Update(id int, t task.Task) error {
	i := l.Index(id)
	if i == -1 {
		return TaskNotFound
	}
	l.update(i, t)
	return nil
}

func (l *List) Get(id int) (task.Task, error) {
	i := l.Index(id)
	if i == -1 {
		return task.Task{}, TaskNotFound
	}
	return *l.tasks[i], nil
}

func (l *List) Tasks() []task.Task {
	r := make([]task.Task, 0, len(l.tasks))
	for _, t := range l.tasks {
		r = append(r, *t)
	}
	return r
}

func (l *List) clear() {
	log.Finest("Clearing all tasks")
	l.tasks = nil
}

func (l *List) Clear() {
	l.clear()
}

func (l *List) Exists(id int) bool {
	for _, t := range l.tasks {
		if t.ID == id {
			return true
		}
	}
	return false
}

func (l *List) Index(id int) int {
	for i, t := range l.tasks {
		if t.ID == id {
			return i
		}
	}
	return -1
}

func (l *List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Tasks     []task.Task `json:"tasks"`
		CreatedAt time.Time   `json:"created_at"`
	}{
		l.Tasks(),
		l.CreatedAt,
	})
}
