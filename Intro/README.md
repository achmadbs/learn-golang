# Organizing Code With Packages

When we writing GO code we can split our code accross packages, and we should have at least one package per application. And a single package can also split into multiple files.

For example:\
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
You may be wondering why we should use _main_ as our package name, why not _app_ or _index_?\
technically we can naming our entry point (main in this case) or index in JavaScript any name we want.

For example as i mentioned above we have a file called _app.go_, when we run it\
```bash
go run app.go
This is first app!
```
it will run without any problem, but the thing is we dont always run our Go code just like that, it only convenient during development process.

For deployment process we should build our Go code to make it available for other people, because those people might not have Go installed in their machine, so they definetely cannot run that command we tried above.

To make our Go available to be executed by other although they dont have Go installed, we can use build command
```bash
go build
```
