package main

import (
	"elliotchance_orderedmap/myorderedmap"
	"fmt"
)

func main() {
	//m := orderedmap.NewOrderedMap()
	//
	//m.Set(34, "bar")
	//m.Set(1002, 1.23)
	//m.Set(123, true)
	//
	//for _, key := range m.Keys() {
	//    value, _:= m.Get(key)
	//    fmt.Println(key, value)
	//}
	//
	//// Iterate through all elements from oldest to newest:
	//for el := m.Front(); el != nil; el = el.Next() {
	//    fmt.Println(el.Key, el.Value)
	//}
	//
	//// You can also use Back and Prev to iterate in reverse:
	//for el := m.Back(); el != nil; el = el.Prev() {
	//    fmt.Println(el.Key, el.Value)
	//}

	// 自己实现的 orderedMap 为按照最后一次更新key的时间排序
	m := myorderedmap.NewOrderedMap()

	m.Set(34, "bar")
	m.Set(1002, 1.23)
	m.Set(123, true)

	for _, key := range m.Keys() {
		value, _ := m.Get(key)
		fmt.Println(key, value)
	}

	m.Set(34, "dsdsds")
	for _, key := range m.Keys() {
		value, _ := m.Get(key)
		fmt.Println(key, value)
	}

	// Iterate through all elements from oldest to newest:
	for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key(), el.Value())
	}

	// You can also use Back and Prev to iterate in reverse:
	for el := m.Back(); el != nil; el = el.Prev() {
		fmt.Println(el.Key(), el.Value())
	}

	fmt.Println("delete")
	m.Delete(34)
	for _, key := range m.Keys() {
		value, _ := m.Get(key)
		fmt.Println(key, value)
	}
}
