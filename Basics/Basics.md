# Basics

For my reference, i'd like to write about basic Go programming in this section.
Such as:

- Brief introduction about Go.
- Variables, string & numbers.
- Loops.
- Boolean & Conditionals.
- Standart Library.
- Use the go command to run code.
- Use the go package.
- Functions call.
- Call functions of an external module.

---

To kickstart my journey, let's write the classic “Hello World!”. :star_struck:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

---

## Variables

In Go _variables_ are expicitly declared and used by the compiler to check type-correctness of function calls as Go is statically typed and compiled language

**[Compiled vs Interpreted Languages](https://www.freecodecamp.org/news/compiled-versus-interpreted-languages/)**
> _A compiled language is a programming language that is converted into machine code (the end result of this process is a binary file typically referred as extension such as .exe .apk) so that the processor can execute it. This means that the program must be compiled again when some changes are made._

var declares 1 or more variables.

```go
var initialString string = "initial string"
```

You can declare multiple variables at once.

```go
var a,b int = 1, 2
```

Go will infer the type of initialized variables. In the example below Go will automatically inter Boolean type to the **_d_** variable

```go
var d = true
```

Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.

```go
var zeroInt int
```

The := syntax is shorthand for declaring and initializing a variable

```go
shortHandVariable := "This is shorthand"
```

_This syntas is only available inside functions._
