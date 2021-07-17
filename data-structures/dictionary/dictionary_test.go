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
		wantResult   interface{}
	}{
		{
			name: "insert invalid with empty key",
			setValues: &valuePair{
				Key:   "",
				Value: "tyrion@email.com",
			},
			wantErrorSet: errors.New("fail to insert the element"),
			wantErrorGet: errors.New("fail to get the element"),
			wantResult:   "",
		},
		{
			name: "insert invalid with empty value",
			setValues: &valuePair{
				Key:   "John",
				Value: "",
			},
			wantErrorSet: errors.New("fail to insert the element"),
			wantErrorGet: errors.New("fail to get the element"),
			wantResult:   "",
		},
		{
			name: "insert invalid with empty key and value",
			setValues: &valuePair{
				Key:   "",
				Value: "",
			},
			wantErrorSet: errors.New("fail to insert the element"),
			wantErrorGet: errors.New("fail to get the element"),
			wantResult:   "",
		},
		{
			name: "insert invalid: key and value with white spaces",
			setValues: &valuePair{
				Key:   "          ",
				Value: "          ",
			},
			wantErrorSet: errors.New("fail to insert the element"),
			wantErrorGet: errors.New("fail to get the element"),
			wantResult:   "",
		},
		{
			name: "insert invalid: key as white space and value as number",
			setValues: &valuePair{
				Key:   "   ",
				Value: 1939,
			},
			wantErrorSet: errors.New("fail to insert the element"),
			wantErrorGet: errors.New("fail to get the element"),
			wantResult:   "",
		},
		{
			name: "insert invalid: key as number and value as white space",
			setValues: &valuePair{
				Key:   1939,
				Value: "   ",
			},
			wantErrorSet: errors.New("fail to insert the element"),
			wantErrorGet: errors.New("fail to get the element"),
			wantResult:   "",
		},
		{
			name: "insert and get with success: key and value as string",
			setValues: &valuePair{
				Key:   "Gandalf",
				Value: "gandalf@email.com",
			},
			wantResult: "gandalf@email.com",
		},
		{
			name: "insert and get with success: key as number and value as string",
			setValues: &valuePair{
				Key:   1914,
				Value: "World War I",
			},
			wantResult: "World War I",
		},
		{
			name: "insert with success: key as string and value as number",
			setValues: &valuePair{
				Key:   "World War II",
				Value: 1939,
			},
			wantResult: 1939,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dicTest := NewDictionary()
			errSet := dicTest.Set(tt.setValues.Key, tt.setValues.Value)
			value, errGet := dicTest.Get(tt.setValues.Key)

			assert.Equal(t, tt.wantErrorSet, errSet)
			assert.Equal(t, tt.wantErrorGet, errGet)
			assert.Equal(t, tt.wantResult, value)
		})
	}
}
