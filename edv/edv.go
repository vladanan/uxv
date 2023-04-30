package main

import "fmt"
import "os"

func main() {
    fmt.Println("Hello from Solaris and ed!!")
    fmt.Println("This program is written in ed. It makes file 'tekst' with some content.")
    os.WriteFile("tekst", []byte("go pise drugi put fajl\n"), 0777)
}
