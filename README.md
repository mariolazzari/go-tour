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

## Basics

### Packages

Every Go program is made up of packages: programs start running in package main.
This program is using the packages with import paths *fmt* and *math/rand"*.

By convention, the package name is the same as the last element of the import path: the *math/rand* package comprises files that begin with the statement package *rand*.

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

### Imports

This code groups the imports into a parenthesized, *factored* import statement.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

### Exported names

In Go, a name is exported if it begins with a capital letter: *Pi* is an exported name; *pi* does not start with a capital letter, so it is not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.pi) -> unexported
	fmt.Println(math.Pi)
}
```

