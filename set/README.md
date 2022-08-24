# Set

Go의 map 을 이용한 Set 구현

### Create Set
```go
set := NewSet[int]()
```

### Add
```go
set.Add(1)
set.Add(2)
set.Add(2) // not added
```

### Delete
```go
set.Delete(2)
```

### Contain
```go
set.Contain(1) // true
set.Contain(2) // false
```

### Length
```go
set.Length() // 1
```

### Items
```go
set.Items() // []int{1}
```
