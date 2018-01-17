package boundary

type Error string

func (err Error) Error() string { return string(err) }

const (
	ErrManagerNotSet   Error = "manager is not set"
	ErrManagerHasNoGet Error = "non-nil manager has no get function"
	ErrManagerHasNoDir Error = "non-nil manager has no dir function"
)

func getFixed(s string) ([]byte, error) {
	if manager == nil {
		return nil, ErrManagerNotSet
	} else if manager.Get == nil {
		return nil, ErrManagerHasNoGet
	}
	return manager.Get(s)
}

func dirFixed(s string) ([]string, error) {
	if manager == nil {
		return nil, ErrManagerNotSet
	} else if manager.Dir == nil {
		return nil, ErrManagerHasNoDir
	}
	return manager.Dir(s)
}

func updateFixed(m, other map[string]string) bool {
	if m == nil {
		return false
	}
	for k, v := range other {
		m[k] = v
	}
	return true
}
