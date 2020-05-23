package fields_test

import (
	"errors"
	"github.com/matthisstenius/field-parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		raw   interface{}
		error error
	}{
		{
			name: "it should succeed",
			raw:  map[string]interface{}{"key": "value"},
		},
		{
			name:  "it should fail if not map[string]interface{}",
			raw:   "fields",
			error: errors.New("raw must be map[string]interface{}"),
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			// When
			_, err := fields.Parse(td.raw)

			// Then
			assert.Equal(t, td.error, err)
		})
	}
}

func TestFields_String(t *testing.T) {
	tests := []struct {
		name  string
		raw   map[string]interface{}
		key   string
		out   string
		error error
	}{
		{
			name: "it should succeed",
			key:  "key",
			raw:  map[string]interface{}{"key": "value"},
			out:  "value",
		},
		{
			name:  "it should fail if key is missing",
			raw:   map[string]interface{}{},
			error: errors.New("could not extract as string"),
		},
		{
			name:  "it should fail if value is not string",
			raw:   map[string]interface{}{"key": 1},
			error: errors.New("could not extract as string"),
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			// When
			f, _ := fields.Parse(td.raw)

			// When
			out, err := f.String(td.key)

			// Then
			assert.Equal(t, td.error, err)
			if err != nil {
				return
			}

			assert.Equal(t, td.out, out)
		})
	}
}

func TestFields_Int64(t *testing.T) {
	tests := []struct {
		name  string
		raw   map[string]interface{}
		key   string
		out   int64
		error error
	}{
		{
			name: "it should succeed",
			key:  "key",
			raw:  map[string]interface{}{"key": int64(1)},
			out:  1,
		},
		{
			name:  "it should fail if key is missing",
			raw:   map[string]interface{}{},
			error: errors.New("could not extract as int64"),
		},
		{
			name:  "it should fail if value is not int",
			raw:   map[string]interface{}{"key": "1"},
			error: errors.New("could not extract as int64"),
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			// When
			f, _ := fields.Parse(td.raw)

			// When
			out, err := f.Int64(td.key)

			// Then
			assert.Equal(t, td.error, err)
			if err != nil {
				return
			}

			assert.Equal(t, td.out, out)
		})
	}
}

func TestFields_Bool(t *testing.T) {
	tests := []struct {
		name  string
		raw   map[string]interface{}
		key   string
		out   bool
		error error
	}{
		{
			name: "it should succeed",
			key:  "key",
			raw:  map[string]interface{}{"key": true},
			out:  true,
		},
		{
			name:  "it should fail if key is missing",
			raw:   map[string]interface{}{},
			error: errors.New("could not extract as bool"),
		},
		{
			name:  "it should fail if value is not bool",
			raw:   map[string]interface{}{"key": "1"},
			error: errors.New("could not extract as bool"),
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			// When
			f, _ := fields.Parse(td.raw)

			// When
			out, err := f.Bool(td.key)

			// Then
			assert.Equal(t, td.error, err)
			if err != nil {
				return
			}

			assert.Equal(t, td.out, out)
		})
	}
}
