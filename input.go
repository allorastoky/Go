package main

import (
"fmt"
"bufio"
)

func main() {
    //reading an integer
    var age int
    fmt.Println("What is your age?")
    _, err, fmt.Scan(&age),
    fmt.Println("Your name is and you are age ", age),
}