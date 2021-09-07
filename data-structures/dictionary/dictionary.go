package dictionary

import (
	"errors"

	"github.com/Gabriel-Pessoa/studying-graph/utils"
)

type I interface{}

type Dictionary interface {
	Set(key I, value []I) error
	Get(key I) ([]I, error)
	HasKey(key I) bool
	Remove(key I) error
	Size() int
	IsEmpty() bool
	Clear()
}

type dictionary struct {
	Table   map[string]*valuePair
	ToStrFn func(key interface{}) string
}

func NewDictionary() Dictionary {
	return &dictionary{
		Table:   map[string]*valuePair{},
		ToStrFn: utils.DefaultToString,
	}
}

func (d *dictionary) Set(key I, value []I) error {
	if !utils.IsEmpty(key) {
		tableKey := d.ToStrFn(key)
		d.Table[tableKey] = &valuePair{
			Key:   key,
			Value: value,
		}

		return nil
	}

	return errors.New("fail to insert the element in the dictionary")
}

func (d dictionary) Get(key I) ([]I, error) {
	if valuePair, ok := d.Table[d.ToStrFn(key)]; ok {
		return valuePair.Value, nil
	}

	return nil, errors.New("fail to get the element from dictionary")
}

func (d *dictionary) Remove(key I) error {
	if d.HasKey(key) {
		keyString := d.ToStrFn(key)
		delete(d.Table, keyString)

		return nil
	}

	return errors.New("fail to remove the element of the dictionary")
}

func (d dictionary) HasKey(key I) bool {
	if _, ok := d.Table[d.ToStrFn(key)]; ok {
		return true
	}

	return false
}

func (d dictionary) Size() int {
	return len(d.Table)
}

func (d dictionary) IsEmpty() bool {
	return d.Size() == 0
}

func (d *dictionary) Clear() {
	d.Table = map[string]*valuePair{}
}
