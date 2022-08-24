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
