# Basics

For my reference, i'd like to write about basic Go programming in this section.\
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

> _An interpreted language is a type of programming language in which the source code is executed line by line by an interpreter during runtime. In an interpreted language, the interpreter reads the source code, interprets each line, and executes the corresponding instructions._

var declares 1 or more variables.

```go
var initialString string = "initial string"
```

You can declare multiple variables at once.

```go
var a,b int = 1, 2
```

Go will infer the type of initialized variables. In the example below Go will automatically mark **_d_** as a Boolean type.

```go
var d = true
```

Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.

```go
var zeroInt int
```

The := syntax is shorthand for declaring and initializing a variable. (_This syntax is only available inside functions._)

```go
shortHandVariable := "This is shorthand"
```

## Constant

Go supports _constants_ of character, string, boolean, and numeric values.

_to declares constant variable we can use const keyword_

```go
const initConstant = 3e20
```

Constant are like variables in Go, except they can't be modified once they have been declared.

```go
const initConstant = 3e20
initConstant = 300;
```

> If we try to modify a constant after it was declared like the example above. We will get a compile-time error like so "cannot assign to initConstant"

Constant can be **untyped**. This can be useful when we working with numbers with arbitrary precision such as _float_, _int_. If the constant is **untyped**, it is explicity converted, where **typed** constant are not.
Lets see how we can use constant with an arbitrary number precision.

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    const n = 500000000
    const d = 3e20 / n

    fmt.Println(d)
    fmt.Println(int64(d))
    fmt.Println(math.Sin(n))
}
```

> In the example above math.Sin is expects a float64 type, since we didn't explicity give it a type it will automatically converted to a float64 type.

```bash
$ go run constant.go
6e+11
600000000000
-0.28470407323754404
```

## String Interpolation

In my opinion, string formatting/interpolation in Go is currently _less_ elegant than JavaScript.

**Examples**: \
<br />
%v - Default representation

The `%v` variants prints the Go syntax to represent a value. Usually we can use this if we are unsure what else to use. With that said it is better to use type-spesific variants.

```go
fmt.Println("I am a %v programmer", "Golang") // I am a Golang programmer
fmt.Println("I have %v years of experience", 4) // I have 4 years of experience
```

%s - String representation

```go
fmt.Println("This is a %v type interpolation", "string") // This is a string type representation
```

%d - Integer in decimal form representation

```go
fmt.Println("I am a software engineer with %d years of experience", 4) // I am a software engineer with 4 years of experience
```

%f - Decimal representation

```go
fmt.Println("I am %f years old", 28.5) // I am 28.5 years old
```

## Loop

In go there are 5 basic for loop patterns

**Three component loop**

```go
sum := 0
for i := 0; i < 5; i++ {
    i++
}
fmt.Println(sum) // 10
```

**While loop**

```go
i := 0
for i < 5; i++ {
    i += 2
}
fmt.Println(i) // 6
```

**Infinite Loop**

```go
i := 0
for {
    i++
}
fmt.Println(i) // will cause infinite loop and this line will never reached
```

**For-each range Loop**

```go
arrayOfString := [2]string{"hello", "world"}
for i, str := range arrayOfString {
    fmt.Println(i, str)
}
```

```bash
0 hello
1 world
```

**Exit a loop**

```go
i := 0
for i < 5; i++ {
    if i%2 === 0 // skip even number
        continue
    i++
}
fmt.Println(i, str) // 3
```

In Go, there are two methods to skip or terminate an iteration. In the given example we utilize **_continue_** statement to bypass a condition and proceed with the rest of the loop. An alternative approach is to use **_break_** statement which allows for immediate termination of the iteration when a spesific condition is met.

## Functions

Functions in Go can take zero or more arguments.\
Similliar with TypeScript, the variable type comes after the variable name.
<br />
<br />
For example, the following function:

```go
func counter(x int): int {
    return x++
}
```

Accepts `x` parameter that has integer type and returns another integer value.

**Multiple Parameter**\
When we have a function with multiple arguments are the same type, the type only needs to be declared after the last one
<br />
<br />
For example, the following function:

```go
import (
	"fmt"
	"strconv"
)

func add(x, y int, result string) string {
	resAdd := x + y
	return "result " + strconv.Itoa(resAdd)
}

func main() {
	var addValue string = add(5, 6, "The result is")
	fmt.Println(addValue)
}
```
```bash
The result is 11
```

since `x` and `y` are the same type, we only need to declare type at y which the last argument that has **_int_** type (_assuming they are in order._)\
If they are not in order they need to be defined separately.


