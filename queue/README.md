# Stack

Generic 을 포함한 Queue 구현

### Create Stack
```go
queue := Queue[int]{}
```

### Add
```go
queue.Add(10)
queue.Add(100)
```

### Poll
```go
queue.Poll() // 10
queue.Poll() // 100
queue.Poll() // 0 (zero value)
```

### Peek
```go
queue.Peek() // 0 (zero value)
queue.Add(100)
queue.Peek() // 100
```

### IsEmpty
```go
queue.IsEmpty() // false
queue.Poll()
queue.IsEmpty() // true
```