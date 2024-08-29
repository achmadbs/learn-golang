# Organizing Code With Packages

When we writing GO code we can split our code accross packages, and we should have at least one package per application. And a single package can also split into multiple files.

For Example:\
suppose we have a file/package called _main.go_

```go
package main

import "fmt"

func main() {
    fmt.Println("This in main!")
}
```
