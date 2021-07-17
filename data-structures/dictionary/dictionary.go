package dictionary

import (
	"errors"

	"github.com/Gabriel-Pessoa/studying-graph/utils"
)

type Dictionary interface {
	Set(key, value interface{}) error
	Get(key interface{}) (interface{}, error)
}

type dictionary struct {
	table   map[string]*valuePair
	toStrFn func(key interface{}) string
}

func NewDictionary() Dictionary {
	return &dictionary{
		table:   map[string]*valuePair{},
		toStrFn: utils.DefaultToString,
	}
}

func (d *dictionary) Set(key, value interface{}) error {
	if !utils.IsEmpty(key) && !utils.IsEmpty(value) {
		tableKey := d.toStrFn(key)
		d.table[tableKey] = &valuePair{
			Key:   key,
			Value: value,
		}
		return nil
	}
	return errors.New("fail to insert the element")
}

func (d dictionary) Get(key interface{}) (interface{}, error) {
	if valuePair, ok := d.table[d.toStrFn(key)]; ok {
		return valuePair.Value, nil
	}
	return "", errors.New("fail to get the element")
}
