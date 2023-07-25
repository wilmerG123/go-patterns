package singleton

import "sync"

var (
	emp  *empleado
	once sync.Once
)

func GetInstance() *empleado {
	once.Do(func() {
		emp = &empleado{}
	})
	return emp
}

type empleado struct {
	id  int
	age string
	mux sync.Mutex
}

func (e *empleado) SetId(n int) {
	e.mux.Lock()
	e.id = n
	e.mux.Unlock()
}

func (e *empleado) GetId() int {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.id
}

func (e *empleado) SetAge(n string) {
	e.mux.Lock()
	e.age = n
	e.mux.Unlock()
}

func (e *empleado) GetAge() string {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.age
}

func (e *empleado) IncrementId() {
	e.mux.Lock()
	e.id++
	e.mux.Unlock()
}
