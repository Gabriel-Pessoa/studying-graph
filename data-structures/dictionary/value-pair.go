package dictionary

import "fmt"

type I interface{}

type ValuePair interface {
	toString() string
}

type valuePair struct {
	Key   I
	Value []I
}

func (v valuePair) toString() string {
	return fmt.Sprintf("[#%v: %v]", v.Key, v.Value)
}
