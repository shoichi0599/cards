# Overview
Learning Go by creating a project that simulates with a deck of playing cards
following the course [Go: The complete Developer's Guide](https://www.udemy.com/go-the-complete-developers-guide/).
```
Functions in Cards:
+--------------------------------------------------------------------------+
| Cards                                                                    |
|--------------------------------------------------------------------------|
| - newDeck        : Create a list of playing cards.                       |
|                    Essentially an array of strings.                      |
| - print          : Log out the contents of a deck of cards               |
| - shuffle        : Shuffles all the cards in a deck.                     |
| - deal           : Createe a "hand" of cards.                            |
| - saveToFile     : Save a list of cards to a file on the local machine.  |
| - newDeckFromFile: Load a list of cards from the local machine.          |
+--------------------------------------------------------------------------+
```

# What I Learned
## Go Command
```
go build   : Compiles a bunch of go source code files
            Example: go build
go run     : Compiles and executes one or two files
go fmt     : Formats all the code in each file in the current directory
go install : Compiles and “installs” a packages
go get     : Downloads the raw source code of someone else’s package
go test    : Runs any tests associated with the current project
```

## Go Package
There are two types of packages, one is executable package and the other is reusable package . 
### 1. Executable package 
Defines a package that can be compiled and then executed by Go.<br>
Must have a func called 'main'<br>
Syntax: ```package main```<br>
Example:
```go
package main

func main() {
}
```
### 2. Reusable package 
Defines a package that can be used as a dependency (helper code) like "fmt"<br>
Needs to be imported.<br>
Syntax: ```package "{PACKAGE_NAME}"```<br>
Example:
```go
package main

import "fmt"

func main() {
    fmt.Println("Import fmt package!")
}
```

## Basic Use of Go
### 1. Variable Declarations
```go
package main
import "fmt"

func main() {
    // var card string = "Ace of Spades"
    // ↓ abbreviated syntax
    // ":=" => when you initialize a variable
    // "="  => when you re-asign a value to a variable
    card := "Ace of Diamonds"
    fmt.Println(card)

    card = "Six of Spades"
    fmt.Println(card)  // => Six of Spades
}
```

### 2. Functions and Return Types
```go
package main
import "fmt"

func main() {
    card := newCard()
    fmt.Println(card)
}

func newCard() string {
    return "Five of Diamonds"
}
```

### 3. Slices and For Loops
```go
package main
import "fmt"

func main() {
    // Array: Fixed length list of things
    // Slice: An array that can grow or shrink
    cards := []string{"Ace of Diamonds", "Two of Diamonds"}
    cards = append(cards, "Six of Spades")

    // i: Index of this element in the slice
    // card: Current card we're iterating over
    // range: Take the slice of 'cards' and loop over it
    for i, card := range cards {
        fmt.Println(i, card)
    }
}
```

### 4. Slice Range Syntax
```go
package main
import "fmt"

func main() {
    // Index:              0         1        2         3
    fruits := []string {"apple", "banana", "grape", "orange"}
    // fruits[startIndexIncluding : upToNotIncluding]
    fmt.Println(fruits[0:2])  // => "apple" "banana"
    fmt.Println(fruits[:2])   // => "apple" "banana"
    fmt.Println(fruits[2:])   // => "grape" "orange"
}
```

### 5. Blank Identifier '_' (underscore)
```go
package main
import "fmt"

func main() {
    // '_' means that there is a variable but we are not gonna use it.
    // The blank identifier provides a way to ignore left-hand side values in an assignment.
    fruits := []string {"apple", "banana", "grape", "orange"}
    for _, fruit := range fruits {
        fmt.Println(fruit)
    }
}
```

### 6. Custom Type Declarations
```go
package main
import "fmt"

func main() {
    cards := deck{"Ace of Diamonds", "Two of Diamonds"}
    fmt.Println(cards)
}

// Create a new type of 'deck' which is a slice of string
type deck []string
```

### 7. Receiver Functions
```go
package main
import "fmt"

func main() {
    cards := deck{"Ace of Diamonds", "Two of Diamonds"}
    fmt.Println(cards)
}

// Create a new type of 'deck' which is a slice of string
type deck []string

// (d deck) -> receiver
// d: The actual copy of the deck we're working with is available
//    in the function as variable called 'd'.
// deck: Every variable of type 'deck' can call this function on itself.
// ----------------------------------------------------------------------------
// Any variable of type "deck" now gets access to the "print" method.
// The receiver sets up methods on variables that we create.
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}
```

### 8. Multiple Return Values in One Function
```go
package main
import "fmt"

func main() {
    title, pages := getBookInfo()
    fmt.Println(title, pages)  // => "War and Peace" 1000
}

// (string, int): Multiple return values. 
//                Define them after a method name.
func getBookInfo() (string, int) {
    return "War and Peace", 1000
}
```

### 9. Byte Slice
```
"Hi there" (string)
   ↓
[72 105 32 116 104 101 114 101 33] (byte slice)

Every element inside of the slice corresponds to an ASCII character code
Example: 72 => H, 105 => i, 32 => " " (space)...etc
Ref: http://www.asciitable.com/
=> byte slice is a computer friendly representation of a string
```

### 10. Type Conversion
```go
package main
import "fmt"

func main() {
    // string -> []byte
    // []byte       : Type we want
    // ("Hi there!"): Value we have
    // []byte("Hi there!")
    
    greeting := "Hi there !"
    fmt.Println([]byte(greeting))
}
```
