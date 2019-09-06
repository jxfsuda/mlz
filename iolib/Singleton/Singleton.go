package Singleton

import (
	"go/types"
	"sync"
)

type Singleton struct {

}

 var singleTonMap = make(map[types.Type]interface{},99)

var once sync.Once

func GetInstance(p types.Type) interface{}{
	val ,ok :=singleTonMap[p]
	if ok {
		return val
	}
	once.Do(func() {
		val = &p{}
		singleTonMap[p] = val
	})
	return val
}
