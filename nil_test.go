package boundary

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func resetManager()                       { manager = nil }
func mockGetOK(s string) ([]byte, error)  { return []byte("foo"), nil }
func mockGetErr(s string) ([]byte, error) { return nil, errors.New("bad get") }

func recoverAndFail(t *testing.T) {
	if r := recover(); r != nil {
		t.Errorf("panic: %v", r)
	}
}

func Test_get(t *testing.T) {
	defer resetManager()
	_, err := get("")
	assert.Error(t, err)

	manager = &Manager{Get: mockGetOK}
	_, err = get("")
	assert.NoError(t, err)

	manager = &Manager{Get: mockGetErr}
	_, err = get("")
	assert.Error(t, err)

	t.Run("panic", func(t *testing.T) {
		defer recoverAndFail(t)
		manager = &Manager{}
		_, err = get("") // this will panic!
		assert.Error(t, err)
	})
}

func Test_update(t *testing.T) {
	foo := map[string]string{"foo": "foo"}
	bar := map[string]string{"bar": "bar"}
	foobar := map[string]string{"foo": "foo", "bar": "bar"}

	updateMap(foo, bar)
	assert.Equal(t, foobar, foo)

	var m map[string]string
	t.Run("panic", func(t *testing.T) {
		defer recoverAndFail(t)
		updateMap(m, bar) // this will panic!
		assert.Equal(t, bar, m)
	})
}

func Test_updateSeemsFixed(t *testing.T) {
	foo := map[string]string{"foo": "foo"}
	bar := map[string]string{"bar": "bar"}
	foobar := map[string]string{"foo": "foo", "bar": "bar"}

	updateMapSeemsFixed(foo, bar)
	assert.Equal(t, foobar, foo)

	var m map[string]string
	updateMapSeemsFixed(m, bar)
	assert.Equal(t, bar, m, "call to make(map[string]string) reassigned in the inner scope")

}
