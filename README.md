# gooption

Rust's Option&lt;T> in Go

## Usage

- Using Some[T] value

```go
import (
    "fmt"

    "github.com/mnaufalhilmym/gooption"
)

data := gooption.Some("Example of string data")

if data.IsSome() { // true
    fmt.Println(data.Unwrap()) // print: Example of string data
}
fmt.Println(data.IsNone()) // print: false
```

- Using None[T] value

```go
import (
    "fmt"

    "github.com/mnaufalhilmym/gooption"
)

data := gooption.None[any]()
dataInt := gooption.None[int]()

if data.IsNone() { // true
    // data.Unwrap() will panic

    fmt.Println(data.UnwrapOr("Another data")) // print: Another data

    fmt.Println(data.UnwrapOrDefault()) // print: <nil>
    fmt.Println(dataInt.UnwrapOrDefault()) // print: 0

    fmt.Println(data.UnwrapOrElse(func () any {
        return []any{"TestData1", 10.01, 11}
    })) // print: ["TestData1", 10.01, 11]
    fmt.Println(dataInt.UnwrapOrElse(func () int {
        return 10
    })) // print: 10
}

fmt.Println(data.IsSome()) // print: false
```
