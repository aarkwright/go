package main

import (
	"fmt"
	"runtime"	// use multiple processors
	"os"		// write to log
	"time"		// for timestamps
)

func main() {
	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./goroutine_mutex_lock.log")
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <- logCh
			if ok {
				f, _ := os.OpenFile("./goroutine_mutex_lock.log", os.O_APPEND, os.ModeAppend)

				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()

	mutex := make(chan bool, 1)

	for i:=1;i<=10;i++ {
		for j:=1;j<=100;j++{
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
				<- mutex
			}()
		}
	}

	fmt.Scanln()
}