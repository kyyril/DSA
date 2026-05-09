package main

import (
	"fmt"
	"sort"
	"sync"
)

type Task struct {
	Name      string
	Priority  int
	DependsOn []string
}

type TaskScheduler struct {
	mu            sync.Mutex
	allTasks      map[string]*Task
	completedTasks map[string]bool
	// pendingTasks menyimpan task yang belum bisa diproses karena dependensi
	pendingTasks  []*Task
}

func NewScheduler() *TaskScheduler {
	return &TaskScheduler{
		allTasks:       make(map[string]*Task),
		completedTasks: make(map[string]bool),
	}
}

func (s *TaskScheduler) AddTask(name string, priority int, dependsOn []string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newTask := &Task{
		Name:      name,
		Priority:  priority,
		DependsOn: dependsOn,
	}
	s.allTasks[name] = newTask
	s.pendingTasks = append(s.pendingTasks, newTask)

	// Urutkan berdasarkan prioritas agar task tanpa dependensi yang prioritasnya tinggi diproses duluan
	s.sortPending()
}

func (s *TaskScheduler) sortPending() {
	sort.SliceStable(s.pendingTasks, func(i, j int) bool {
		return s.pendingTasks[i].Priority < s.pendingTasks[j].Priority
	})
}

func (s *TaskScheduler) ProcessNext() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, task := range s.pendingTasks {
		// Cek apakah semua dependensi sudah selesai
		canProcess := true
		for _, dep := range task.DependsOn {
			if !s.completedTasks[dep] {
				canProcess = false
				break
			}
		}

		if canProcess {
			// Hapus task dari pending
			s.pendingTasks = append(s.pendingTasks[:i], s.pendingTasks[i+1:]...)
			// Tandai sebagai selesai
			s.completedTasks[task.Name] = true
			return task.Name
		}
	}

	return "" // Tidak ada task yang siap diproses saat ini
}

func main() {
	s := NewScheduler()

	// Task B bergantung pada A
	// Task C bergantung pada B
	s.AddTask("C", 1, []string{"B"}) // Priority 1 (Urgent) tapi nunggu B
	s.AddTask("A", 3, []string{})    // Priority 3
	s.AddTask("B", 2, []string{"A"}) // Priority 2 tapi nunggu A

	fmt.Println("Order of Execution:")
	for {
		name := s.ProcessNext()
		if name == "" {
			break
		}
		fmt.Println("-> Completed:", name)
	}
}