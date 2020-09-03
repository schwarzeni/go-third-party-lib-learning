package myasync

import "reflect"

type MapTasks map[string]interface{}

type MapTaskResults struct {
	res map[string][]interface{}
	err map[string]error
}

func (m *MapTaskResults) Value(key string) interface{} {
	if _, ok := m.res[key]; !ok {
		panic("no value for key: " + key)
	}
	return m.res[key]
}

type chanReturn struct {
	returns []reflect.Value
	key     string
}

func (m *MapTaskResults) Error(key string) error {
	if v, ok := m.err[key]; ok {
		return v
	}
	return nil
}

func Concurrent(mapTasks MapTasks) MapTaskResults {
	resMap := MapTaskResults{make(map[string][]interface{}), make(map[string]error)}
	resChan := make(chan chanReturn)

	for k, v := range mapTasks {
		go runConcurrentFunc(resChan, k, v)
	}

	for range mapTasks {
		cr := <-resChan
		returns, err := extractError(cr.returns)
		if err != nil {
			resMap.err[cr.key] = err
		} else {
			resMap.res[cr.key] = values2interfaces(returns)
		}
	}
	return resMap
}

func runConcurrentFunc(resChan chan chanReturn, k string, v interface{}) {
	resChan <- chanReturn{
		returns: callFunc(v),
		key:     k,
	}
}
