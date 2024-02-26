package main

import "fmt"
import "os"

func main() {
    fmt.Println("Input one character:")
    var char rune
    _, err := fmt.Scanf("%c", &char)
    if err != nil {fmt.Printf("Erorr reading character!")}
    var str = string(char)
    var bajtovi = []byte(str)
    fmt.Println("This program is written in ed and vim. It wrote/updated file 'text' with: ", str)
    //os.WriteFile("text", []byte(&char), 0777)
    os.WriteFile("text", bajtovi, 0777)
}
