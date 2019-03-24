# Overview

## Basic Use of GO 

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

### 6. Multiple Return Values in One Function
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


