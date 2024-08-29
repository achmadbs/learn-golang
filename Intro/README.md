# Organizing Code With Packages

When we writing GO code we can split our code accross packages, and we should have at least one package per application. And a single package can also split into multiple files.

For Example:\
suppose we have a file/package called _app.go_

```go
package main

import "fmt"

func main() {
    fmt.Println("This is first app!")
}
```
then we can also create another file _second_app.go_ which also belongs to the main package.

```go
package main

import "fmt"

func main() {
    fmt.Println("This in second app!")
}
```

## Main Package
You may be wondering why i mention _main_ so often, why not _app_ or _index_?
