// Package app - логика менеджера задач
// Содержит структуры данных и CRUD операции для работы с задачами
package app

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Приоритет задачи
type Priority int

const (
	PriorityLow    Priority = 0
	PriorityMedium Priority = 1
	PriorityHigh   Priority = 2
)

// Строковое представление приоритета
func (p Priority) String() string {
	switch p {
	case PriorityHigh:
		return "Высокий"
	case PriorityMedium:
		return "Средний"
	default:
		return "Низкий"
	}
}

// Task - структура задачи
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    Priority  `json:"priority"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}

// TaskManager - менеджер задач
type TaskManager struct {
	tasks      []Task
	nextID     int
	mu         sync.RWMutex
	dataFile   string
	onModified func() // Callback при изменении данных
}

// NewTaskManager создаёт новый менеджер задач
func NewTaskManager() *TaskManager {
	// Путь к файлу с задачами в домашней директории пользователя
	homeDir, _ := os.UserHomeDir()
	dataDir := filepath.Join(homeDir, ".taskmanager_data")
	os.MkdirAll(dataDir, 0755)

	tm := &TaskManager{
		tasks:    make([]Task, 0),
		nextID:   1,
		dataFile: filepath.Join(dataDir, "tasks.json"),
	}

	// Загружаем сохранённые задачи
	tm.loadTasks()

	return tm
}

// SetOnModified устанавливает callback при изменении данных
func (tm *TaskManager) SetOnModified(callback func()) {
	tm.onModified = callback
}

// notifyModified вызывает callback при изменении
func (tm *TaskManager) notifyModified() {
	if tm.onModified != nil {
		tm.onModified()
	}
}

// AddTask добавляет новую задачу
func (tm *TaskManager) AddTask(title, description string, priority Priority) *Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task := Task{
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		Priority:    priority,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	tm.nextID++
	tm.tasks = append(tm.tasks, task)
	tm.saveTasks()
	tm.notifyModified()

	return &task
}

// GetAllTasks возвращает все задачи
func (tm *TaskManager) GetAllTasks() []Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	// Возвращаем копию среза
	result := make([]Task, len(tm.tasks))
	copy(result, tm.tasks)
	return result
}

// GetTask возвращает задачу по ID
func (tm *TaskManager) GetTask(id int) *Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			task := tm.tasks[i]
			return &task
		}
	}
	return nil
}

// UpdateTask обновляет задачу
func (tm *TaskManager) UpdateTask(id int, title, description string, priority Priority) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks[i].Title = title
			tm.tasks[i].Description = description
			tm.tasks[i].Priority = priority
			tm.saveTasks()
			tm.notifyModified()
			return true
		}
	}
	return false
}

// ToggleComplete переключает статус выполнения задачи
func (tm *TaskManager) ToggleComplete(id int) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks[i].Completed = !tm.tasks[i].Completed
			if tm.tasks[i].Completed {
				tm.tasks[i].CompletedAt = time.Now()
			} else {
				tm.tasks[i].CompletedAt = time.Time{}
			}
			tm.saveTasks()
			tm.notifyModified()
			return true
		}
	}
	return false
}

// DeleteTask удаляет задачу
func (tm *TaskManager) DeleteTask(id int) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			tm.saveTasks()
			tm.notifyModified()
			return true
		}
	}
	return false
}

// GetTaskCount возвращает количество задач
func (tm *TaskManager) GetTaskCount() int {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return len(tm.tasks)
}

// GetCompletedCount возвращает количество выполненных задач
func (tm *TaskManager) GetCompletedCount() int {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	count := 0
	for _, task := range tm.tasks {
		if task.Completed {
			count++
		}
	}
	return count
}

// GetTasksByPriority возвращает задачи с указанным приоритетом
func (tm *TaskManager) GetTasksByPriority(priority Priority) []Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	result := make([]Task, 0)
	for _, task := range tm.tasks {
		if task.Priority == priority {
			result = append(result, task)
		}
	}
	return result
}

// saveTasks сохраняет задачи в файл
func (tm *TaskManager) saveTasks() {
	data := struct {
		Tasks  []Task `json:"tasks"`
		NextID int    `json:"next_id"`
	}{
		Tasks:  tm.tasks,
		NextID: tm.nextID,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}

	os.WriteFile(tm.dataFile, jsonData, 0644)
}

// loadTasks загружает задачи из файла
func (tm *TaskManager) loadTasks() {
	data, err := os.ReadFile(tm.dataFile)
	if err != nil {
		return
	}

	var loaded struct {
		Tasks  []Task `json:"tasks"`
		NextID int    `json:"next_id"`
	}

	if err := json.Unmarshal(data, &loaded); err != nil {
		return
	}

	tm.tasks = loaded.Tasks
	tm.nextID = loaded.NextID

	if tm.nextID == 0 {
		tm.nextID = 1
	}
}

// ClearAllTasks удаляет все задачи
func (tm *TaskManager) ClearAllTasks() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.tasks = make([]Task, 0)
	tm.saveTasks()
	tm.notifyModified()
}
