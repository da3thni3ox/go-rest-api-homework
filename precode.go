package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Task ...
type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

// Ниже напишите обработчики для каждого эндпоинта
// ...

// GET Получение списка задач
func getTasks(w http.ResponseWriter, r *http.Request) {

	resp, err := json.Marshal(tasks)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

// GET Получение задачи по ID
func getTaskById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	task, ok := tasks[id]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(task)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// POST добавление новой задачи
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&newTask)

	fmt.Println(newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//TODO проверить наличие заголовка content-type
	tasks[newTask.ID] = newTask

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// DELETE удаление задачи
func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, ok := tasks[id]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	delete(tasks, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := chi.NewRouter()

	// здесь регистрируйте ваши обработчики
	// ...

	r.Get("/tasks", getTasks)
	r.Get("/tasks/{id}", getTaskById)
	r.Post("/tasks", createTask)
	r.Delete("/tasks/{id}", deleteTask)
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
