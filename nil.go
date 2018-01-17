package boundary

import "errors"

var ( // all of these are initialized to nil
	ip *int //any pointer starts out as nil

	f     func(int) int
	v     interface{}
	m     map[string]string
	slice []byte //also any slice

	ch  chan interface{}
	in  <-chan string
	out chan<- string
)

var ( //none of these are initialized to nil
	i int //also (u)int8, (u)int16, (u)int32, (u)int64, complex64, complex128

	a [3]*int //a[0], a[1], and a[2] are nil, but a is not

	b bool
	q struct{} //like an array, may contain nil elements, but is never nil itself
	s string
)

type Get func(string) ([]byte, error)
type Dir func(string) ([]string, error)
type Manager struct {
	Get
	Dir
}

var manager *Manager

func get(s string) ([]byte, error) {
	if manager == nil {
		return nil, errors.New("manager is not set")
	}
	return manager.Get(s)
}

func set(s string) ([]string, error) {
	if manager == nil {
		return nil, errors.New("manager is not set")
	}
	return manager.Dir(s)
}

func updateMap(m, other map[string]string) {
	for k, v := range other {
		m[k] = v
	}
}

func updateMapSeemsFixed(m, other map[string]string) {
	if m == nil {
		m = make(map[string]string)
	}

	for k, v := range other {
		m[k] = v
	}
}
