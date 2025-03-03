## Data Types

### Intergers:
- int: Default integer type (size depends on platform, 32bit or 64bit)
- int8, int16, int32
- uint:
- unit8, uint16,

### Floating-Point Numbers:
- Float32: Single-precision floating-point.
- Float64: Double precesion floating-point

### Boolean
- bool: Can have only two values, true or false

### Strings
- string: Asequence of characters. Strings in Go are immutable (cannot be changed affer creation)

### Array
- Fixed-size collections of elements of the size type.

### Slices
- Dynamic arras. Unlike arrays. slices can grow and shrink in size.

### Structs
- A collection of fields (variables) grouped under a single name, which can be of different types.

### Maps
- A collection of key-value pairs, where keys must be unique.

### Pointers
- Variables that store the memory address of another variable. They are often used to pass variables by reference instead of by value.


`Examples`
```
var keyword
```

### Build and Run Go
```
go run main.go
```

```
go run --work main.go
```

```
go build main.go
./main
```

### Show go ENV
```
go env
```

```
GOOS='windows' go build main.go
```