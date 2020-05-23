package fields

import (
	"errors"
)

type Fields struct {
	raw map[string]interface{}
}

// Parse raw data into Fields. Fails if raw data is not map[string]interface{}
func Parse(raw interface{}) (*Fields, error) {
	v, ok := raw.(map[string]interface{})
	if !ok {
		return nil, errors.New("raw must be map[string]interface{}")
	}

	return &Fields{raw: v}, nil
}

// String tries to extract given key as string
func (f *Fields) String(key string) (string, error) {
	if v, ok := f.raw[key].(string); ok {
		return v, nil
	}
	return "", errors.New("could not extract as string")
}

// Int64 tries to extract given key as int64
func (f *Fields) Int64(key string) (int64, error) {
	if v, ok := f.raw[key].(int64); ok {
		return v, nil
	}
	return 0, errors.New("could not extract as int64")
}

// Bool tries to extract given key as bool
func (f *Fields) Bool(key string) (bool, error) {
	if v, ok := f.raw[key].(bool); ok {
		return v, nil
	}
	return false, errors.New("could not extract as bool")
}
