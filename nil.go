package boundary

import (
	"errors"
)

var ( //none of these are initialized to nil
	i int //also (u)int8, (u)int16, (u)int32, (u)int64, complex64, complex128

	a [3]*int //a[0], a[1], and a[2] are nil, but a is not

	b bool
	q struct{} //like an array, structs may contain nil elements, but is never nil itself
	s string
)

// this example taken from real code in github.com/eyecuelab/kit/assets

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
	return manager.Get(s) // this can cause a panic, because manager.Get could be a nil function pointer
}

func set(s string) ([]string, error) {
	if manager == nil {
		return nil, errors.New("manager is not set")
	}
	return manager.Dir(s) //ditto
}

func updateMap(m, other map[string]string) {
	for k, v := range other {
		m[k] = v //can cause panic if m is nil map.
	}
}

func updateMapSeemsFixed(m, other map[string]string) {
	if m == nil {
		m = make(map[string]string) //this creates a new map which drops out of scope when the function returns;
		//it's actually even worse than the first one, which will at least panic to tell you something is going wrong!
	}

	for k, v := range other {
		m[k] = v
	}
}
