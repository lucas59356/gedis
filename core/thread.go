package core

import (
	"errors"
	"log"
	"strings"
)

var (
	// ErrWrongType When you try per example get a int in a string field
	ErrWrongType = errors.New("This key is not of this type")
	// ErrNotFound Key not found
	ErrNotFound = errors.New("Key not found")
	// ErrInvalidType Program do not support this type of data
	ErrInvalidType = errors.New("Invalid data type: data not supported")
)

//
type SingleValue struct {
	Key   string
	Type  int8
	Value interface{}
}

// Thread Object that represents a basic key/value instance
type Thread struct {
	Values   map[string]interface{}
	Types    map[string]int8
	SetQueue chan (SingleValue)
}

// NewThread Creates a new Thread object
func NewThread() *Thread {
	t := Thread{
		Values:   map[string]interface{}{},
		Types:    map[string]int8{},
		SetQueue: make(chan (SingleValue)),
	}
	t.SetLoop()
	return &t
}

// Get Gets an value
func (t *Thread) Get(key string) (interface{}, int8, error) {
	pkey := strings.ToLower(key)
	v, ok := t.Values[pkey]
	if !ok {
		return v, TypeWhatever, ErrNotFound
	}
	tp := t.Types[pkey]
	log.Printf("GEDIS_GET %s %s = %v", pkey, Types[tp], v)
	return v, tp, nil
}

// Set Sets a key on the store
func (t *Thread) Set(key string, v interface{}) (interface{}, int8, error) {
	key = strings.ToLower(key)
	tp := GuessDataType(v)
	if tp == TypeWhatever {
		return nil, TypeWhatever, ErrInvalidType
	}
	t.Values[key] = v
	t.Types[key] = tp
	t.SetQueue <- SingleValue{
		Key:   key,
		Type:  tp,
		Value: v,
	}
	return v, tp, nil
}

// SetLoop Setup loop for adding more values
func (t *Thread) SetLoop() {
	go func() {
		for item := range t.SetQueue {
			log.Printf("GEDIS_SET %s %s = %v", item.Key, Types[item.Type], item.Value)
			t.Types[item.Key] = item.Type
			t.Values[item.Key] = item.Value
		}
	}()
}

// Del Deletes a key on the store
func (t *Thread) Del(key string) error {
	key = strings.ToLower(key)
	log.Printf("GEDIS_DEL %s", key)
	delete(t.Types, key)
	delete(t.Values, key)
	return nil
}
