# github.com/elliotchance/orderedmap 学习笔记

2020.09.04

[github.com/elliotchance/orderedmap](github.com/elliotchance/orderedmap)

比普通的 map 增加了一个特性，当遍历 map 的时候顺序为插入 k/v 的时序

源码非常简单，除了一个 go 中的 map 外，还有一个 `container/list` 中的 List 用于记录 key 的顺序，大致数据结构如下

```go
type OrderedMap struct {
	kv map[interface{}]*list.Element
	ll *list.List
}
```

增删改查都是 `O(1)` 的时间复杂度。由于 `list.List` 使用的是双向链表实现，所以对于 `list.Delete()` 方法，只需要将删除节点的 `Prev` 和 `Next` 置为 `nil` ，同时将其前后相邻节点的 `Prev` 和 `Next` 做一些修改即可，而当前节点存储在 map 中，可以以 O(1) 的时间复杂度获取。

我自己也尝试实现了相关的方法，在 [myorderedmap](myorderedmap) 下，只不过 key 是按照最后一次更新的时序来排序的
