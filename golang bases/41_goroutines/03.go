package main

import (
    "fmt"
    "sync"
    "time"
)

func imprimirTexto(texto string, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println(texto)
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go imprimirTexto("Hello", &wg)
    wg.Add(1)
    go imprimirTexto("World", &wg)

    wg.Wait()
    
    go func(name string){
	fmt.Printf("adios %s \n",name)
    }("william")

    time.Sleep(time.Second * 1)
}
