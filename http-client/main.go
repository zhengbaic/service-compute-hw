package main

import (
	"fmt"
	"time"

	"./Agency"
)

func main() {
	t1 := time.Now()
	Agency.SyncRun()
	fmt.Println("Time of SyncRun:", time.Since(t1))

	t2 := time.Now()
	Agency.AsyncRun()
	fmt.Println("Time of AsyncRun:", time.Since(t2))
}
