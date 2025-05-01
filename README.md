# A tour of Go

A [tour](https://go.dev/tour/list) of [Go](https://go.dev/) language.

## Welcome

[Welcome](https://go.dev/tour/welcome/1)

### Hello

```go
// hello.go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

```sh 
go run hello.go
```

### Playground

Go [playground](https://go.dev/play/)

The service receives a Go program, compiles, links, and runs the program inside a sandbox, then returns the output.

There are limitations to the programs that can be run in the playground:
- In the playground the time begins at 2009-11-10 23:00:00 UTC (determining the significance of this date is an exercise for the reader). This makes it easier to cache programs by giving them deterministic output.
There are also limits on execution time and on CPU and memory usage, and the program cannot access external network hosts.
- The playground uses the latest stable release of Go.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
}
```