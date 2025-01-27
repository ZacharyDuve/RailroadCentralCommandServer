package main

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...

func main() {
	log.Println("Starting Railroad-Central-Command-Server, have fun training")
	rootCTX := context.Background()

	ctx, _ := context.WithTimeout(rootCTX, time.Second*5)

	totalSleepSec := 0
	for {
		select {
		case <-ctx.Done():
			{
				log.Println("Ctx canceled")
				break
			}
		default:
			{
				log.Println("Sleeping for a second", totalSleepSec)
				totalSleepSec++
				time.Sleep(time.Second)
			}
		}
	}

	log.Println("All done")
	select {
	case <-rootCTX.Done():
		{
			log.Println("Was done at the root as well")
		}
	default:
		{
			log.Println("only child was canceled")
		}
	}
}
