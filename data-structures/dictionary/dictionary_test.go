package dictionary

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAndGetDictionary(t *testing.T) {
	var tests = []struct {
		name         string
		setValues    *valuePair
		wantErrorSet error
		wantErrorGet error
		wantResult   []I
	}{
		{
			name: "insert invalid: empty key",
			setValues: &valuePair{
				Key:   "",
				Value: []I{"tyrion@email.com"},
			},
			wantErrorSet: errors.New("fail to insert the element in the dictionary"),
			wantErrorGet: errors.New("fail to get the element from dictionary"),
			wantResult:   nil,
		},
		{
			name: "insert invalid: key as white spaces",
			setValues: &valuePair{
				Key:   "          ",
				Value: []I{},
			},
			wantErrorSet: errors.New("fail to insert the element in the dictionary"),
			wantErrorGet: errors.New("fail to get the element from dictionary"),
			wantResult:   nil,
		},
		{
			name: "insert invalid: with nil",
			setValues: &valuePair{
				Key:   nil,
				Value: []I{},
			},
			wantErrorSet: errors.New("fail to insert the element in the dictionary"),
			wantErrorGet: errors.New("fail to get the element from dictionary"),
			wantResult:   nil,
		},
		{
			name: "insert and get with success: key and value as string",
			setValues: &valuePair{
				Key:   "Gandalf",
				Value: []I{"gandalf@email.com"},
			},
			wantResult: []I{"gandalf@email.com"},
		},
		{
			name: "insert and get with success: key as number and value as string",
			setValues: &valuePair{
				Key:   1914,
				Value: []I{"World War I"},
			},
			wantResult: []I{"World War I"},
		},
		{
			name: "insert and get with success: key as string and value as number",
			setValues: &valuePair{
				Key:   "World War II",
				Value: []I{1939},
			},
			wantResult: []I{1939},
		},
		{
			name: "insert and get with success: key as string and value as slice",
			setValues: &valuePair{
				Key: "books",
				Value: []I{
					[]string{"The lord of the rings", "The hobbit", "The silmarillion"},
				},
			},
			wantResult: []I{
				[]string{"The lord of the rings", "The hobbit", "The silmarillion"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dicTest := NewDictionary()

			assert.Equal(t, 0, dicTest.Size())
			assert.Equal(t, true, dicTest.IsEmpty())

			errSet := dicTest.Set(tt.setValues.Key, tt.setValues.Value)
			value, errGet := dicTest.Get(tt.setValues.Key)

			assert.Equal(t, tt.wantErrorSet, errSet)
			assert.Equal(t, tt.wantErrorGet, errGet)
			assert.Equal(t, tt.wantResult, value)
		})
	}
}

func TestRemoveDictionary(t *testing.T) {
	errRemove := errors.New("fail to remove the element of the dictionary")

	const min = 1
	const max = 5
	const size = max - min + 1

	dicTest := NewDictionary()

	for i := min; i <= max; i++ {
		err := dicTest.Set(i, []I{})
		assert.Equal(t, nil, err)
	}

	assert.Equal(t, size, dicTest.Size())

	for i := min; i <= max; i++ {
		assert.Equal(t, nil, dicTest.Remove(i))
	}

	// elements do not exist
	for i := min; i <= max; i++ {
		assert.Equal(t, errRemove, dicTest.Remove(i))
	}

	assert.Equal(t, true, dicTest.IsEmpty())
}

func TestHasKeyDictionary(t *testing.T) {
	const min = 1
	const max = 5
	const size = max - min + 1

	dicTest := NewDictionary()

	for i := min; i <= max; i++ {
		err := dicTest.Set(i, []I{})
		assert.Equal(t, nil, err)
	}

	assert.Equal(t, size, dicTest.Size())

	for i := min; i <= max; i++ {
		assert.Equal(t, true, dicTest.HasKey(i))
		assert.Equal(t, nil, dicTest.Remove(i))
		assert.Equal(t, false, dicTest.HasKey(i))
	}
}

func TestClearDictionary(t *testing.T) {
	dicTest := NewDictionary()

	dicTest.Clear()

	dicTest.Set(1, []I{})
	dicTest.Set(2, []I{})

	dicTest.Clear()
	assert.Equal(t, true, dicTest.IsEmpty())
}
