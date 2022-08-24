# Stack

Generic 을 포함한 Stack 구현

### Create Stack
```go
stack := Stack[int]{}
```

### Push
```go
stack.Push(10)
stack.Push(100)
```

### Pop
```go
stack.Pop() // 100
stack.Pop() // 10
stack.Pop() // 0 (zero value)
```

### Peek
```go
stack.Peek() // 0 (zero value)
stack.Push(100)
stack.Peek() // 100
```

### Length
```go
stack.Length() // 1
```

### IsEmpty
```go
stack.IsEmpty() // false
stack.Pop()
stack.IsEmpty() // true
```