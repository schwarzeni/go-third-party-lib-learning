package myasync

import "reflect"

func callFunc(task interface{}, args ...reflect.Value) (returns []reflect.Value) {
	fn := reflect.ValueOf(task)
	if fn.Kind() != reflect.Func {
		panic("variable " + fn.String() + " should be a function")
	}
	returns = fn.Call(args)
	return
}

func extractError(fullReturns []reflect.Value) (returns []reflect.Value, err error) {
	var ok bool
	if len(fullReturns) == 0 {
		return fullReturns, nil
	}
	err, ok = fullReturns[len(fullReturns)-1].Interface().(error)
	if !ok {
		return fullReturns, nil
	}
	return fullReturns[:len(fullReturns)-1], err
}

func values2interfaces(values []reflect.Value) (interfaces []interface{}) {
	interfaces = make([]interface{}, len(values))
	for i := range interfaces {
		interfaces[i] = values[i].Interface()
	}
	return interfaces
}

func interfaces2values(interfaces []interface{}) (values []reflect.Value) {
	values = make([]reflect.Value, len(interfaces))
	for i := range values {
		values[i] = reflect.ValueOf(interfaces[i])
	}
	return values
}
