package dictionary

import "fmt"

type ValuePair interface {
	toString() string
}

type valuePair struct {
	Key, Value interface{}
}

func (v valuePair) toString() string {
	return fmt.Sprintf("[#%v: %v]", v.Key, v.Value)
}
