# github.com/rafaeldias/async 学习笔记

2020.08.31 - 2020.09.02

[github.com/rafaeldias/async](https://github.com/rafaeldias/async)

作者开发的这个库给出了几个简单的接口，分别为: 

- `Waterfall` 将一组函数按序执行
- `Parallel` / `Concurrent` 并发执行一组函数

样例在该库的 README 中，部分如下：

Waterfall: 

```go
func fib(p, c int) (int, int) { return c, p + c }

func main() {
    // execution in series.
    res, e := async.Waterfall(async.Tasks{
        fib,
        fib,
        fib,
        func(p, c int) (int, error) { return c, nil },
    }, 0, 1)

    if e != nil {
      fmt.Printf("Error executing a Waterfall (%s)\n", e.Error())
    }

    fmt.Println(res[0].(int)) // Prints 3
}
```

Parallel:

```go
func main() {
    res, e := async.Parallel(async.MapTasks{
        "one": func() int { /* ... */ },
        "two": func() int { /* ... */ },
        "three": func() int { /* ... */ },
    })
    if e != nil {
        fmt.Printf("Errors [%s]\n", e.Error())
    }
    fmt.Println("Results from task 'two': %v", res.Key("two"))
}
```

看了一下作者实现的代码，思路比较清晰，使用到了 Go 语言的 `reflect` 库，使用反射，实现函数类型和 `interface` 类型之间的转换，大致思路如下

```go
fn := reflect.ValueOf(fnInterface)
returns = fn.Call(args)
```

```go
fnInterface := fn.Interface()
```

因为希望将任意类型的函数作为形参传到相应的接口中，所以接口参数的定义中和函数有关的变量都应该是 `interface` 类型的

```go
type Tasks []interface{}

type MapTasks map[string]interface{}
```

大致看完实现思路后，我突然想到，如果一个面试官让你实现这两个接口: `Waterfall` 和 `Parallel`，你会怎么做呢？ 我自己尝试实现了一下，没有用到递归，其他思路类似，在 `myasync` 文件夹里。实现起来还是比较简单的。

