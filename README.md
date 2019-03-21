# Overview

## Basic Use of GO 
1. Variable Declarations
2. Functions and Return Types
3. Slices and For Loops
4. Slice Range Syntax
5. Blank Identifier '_' (underscore)
6. Multiple Return Values in One Function
```$xslt
package main

import "fmt"

func main() {
    // 1) Variable Declarations
    var card string = "Ace of Spades"
    // ↓ abbreviated syntax
    // ":=" -> when you initialize a variable
    // "=" -> when you re-asign a value to a variable
    card := "Ace of Diamonds"
    card = "Six of Spades"

    // 2) Functions and Return Types
    card := newCard()
    fmt.Println(card)

    // 3) Slices and For Loops
    cards := []string{"Ace of Diamonds", newCard()}
    cards = append(cards, "Six of Spades")
    // i: Index of this element in the array
    // card: Current card we're iterating over
    // range: Take the slice of 'cards' and loop over it
    for i, card := range cards {
        fmt.Println(i, card)
    }
    
    // 4) Slice Range Syntax
    // Index:              0         1        2         3
    fruits := []string {"apple", "banana", "grape", "orange"}
    // fruits[startIndexIncluding : upToNotIncluding]
    fmt.Println(fruits[0:2])  // => "apple" "banana"
    fmt.Println(fruits[:2])   // => "apple" "banana"
    fmt.Println(fruits[2:])   // => "grape" "orange"
    
    // 5) Blank identifier '_'
    // '_' means that there is a variable but we are not gonna use it.
    // The blank identifier provides a way to ignore left-hand side values in an assignment.
    fruits := []string {"apple", "banana", "grape", "orange"}
    for _, fruit := range fruits {
        fmt.Println(fruit)
    }
    
    // 6) Multiple Return Values in One Function
    // (string, int): Define multiple return values after a method name.
    func getBookInfo() (string, int) {
        return "War and Peace", 1000
    }
    title, pages := getBookInfo()
    fmt.Println(title, pages)  // => "War and Peace" 1000
}

func newCard() string {
    return "Five of Diamonds"
}
```
