package myasync

import "reflect"

type Tasks []interface{}

func Waterfall(tasks Tasks, args ...interface{}) (returns []interface{}, err error) {
	var funcArgs []reflect.Value
	var funcReturns []reflect.Value
	if len(tasks) == 0 {
		return []interface{}{}, nil
	}

	funcArgs = interfaces2values(args)

	for _, task := range tasks {
		funcReturns = callFunc(task, funcArgs...)
		if funcReturns, err = extractError(funcReturns); err != nil {
			break
		}
		funcArgs = funcReturns
	}

	returns = values2interfaces(funcReturns)

	return
}
