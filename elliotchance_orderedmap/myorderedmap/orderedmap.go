package myorderedmap

import "container/list"

type orderedMapElement struct {
	key, value interface{}
}

type Element struct {
	element *list.Element
}

func (e *Element) Next() *Element {
	if e.element.Next() == nil {
		return nil
	}
	return &Element{e.element.Next()}
}

func (e *Element) Prev() *Element {
	if e.element.Prev() == nil {
		return nil
	}
	return &Element{e.element.Prev()}
}

func (e *Element) Value() interface{} {
	return e.element.Value.(*orderedMapElement).value
}

func (e *Element) Key() interface{} {
	return e.element.Value.(*orderedMapElement).key
}

type OrderedMap struct {
	ll *list.List
	mm map[interface{}]*list.Element
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		ll: list.New(),
		mm: make(map[interface{}]*list.Element),
	}
}

// Set update key value
func (m *OrderedMap) Set(key, value interface{}) {
	if e, ok := m.mm[key]; ok {
		m.ll.Remove(e)
	}
	v := &orderedMapElement{key: key, value: value}
	e := m.ll.PushBack(v)
	m.mm[key] = e
}

func (m *OrderedMap) Get(key interface{}) (interface{}, bool) {
	e, ok := m.mm[key]
	if !ok {
		return nil, false
	}
	return e.Value.(*orderedMapElement).value, true
}

func (m *OrderedMap) Keys() []interface{} {
	keys := make([]interface{}, len(m.mm))
	i := 0
	for ee := m.ll.Front(); ee != nil; ee = ee.Next() {
		keys[i] = ee.Value.(*orderedMapElement).key
		i++
	}
	return keys
}

func (m *OrderedMap) Delete(key interface{}) {
	e, ok := m.mm[key]
	if !ok {
		return
	}
	m.ll.Remove(e)
	delete(m.mm, e.Value.(*orderedMapElement).key)
}

func (m *OrderedMap) Front() *Element {
	return &Element{m.ll.Front()}
}

func (m *OrderedMap) Back() *Element {
	return &Element{m.ll.Back()}
}
