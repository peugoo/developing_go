package main

import (
	"fmt"
	"time"
)

func task2(ch chan string, str string){
    for i := 0; i < 5; i++ {
        ch <- str + " "
    }
}


func main(){
    channel := make(chan string, 10)
    N := "asd"
    go task2(channel, N)
    time.Sleep(time.Second)
    close(channel)
    for i := range channel{
        fmt.Print(i)
    }
}
