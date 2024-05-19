package main

import "fmt"

func main() {

    {{if .loop}}
    for i := 0; i < 10; i++ {
    {{end}}
        fmt.Println("Hello world!")
    {{if .loop}}
    }
    {{end}}
}