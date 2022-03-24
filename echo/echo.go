package main

import "fmt"
import "os"
import "time"

func main() {
	startTime := time.Now()
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%s %d\n", os.Args[i], i)
	}
	endTime := time.Now()
	fmt.Printf("Clock time: %d\n", endTime.Sub(startTime))
}
