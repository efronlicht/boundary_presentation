package boundary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getFixed(t *testing.T) {
	defer resetManager()
	_, err := getFixed("")
	assert.Error(t, err)

	manager = &Manager{Get: mockGetOK}
	_, err = getFixed("")
	assert.NoError(t, err)

	manager = &Manager{Get: mockGetErr}
	_, err = getFixed("")
	assert.Error(t, err)

}

func Test_updateFixed(t *testing.T) {
	foo := map[string]string{"foo": "foo"}
	bar := map[string]string{"bar": "bar"}
	foobar := map[string]string{"foo": "foo", "bar": "bar"}

	updateFixed(foo, bar)
	assert.Equal(t, foobar, foo)

	var m map[string]string
	assert.False(t, updateFixed(m, bar))

}
