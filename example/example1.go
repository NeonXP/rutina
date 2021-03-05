// +build ignore

package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/neonxp/rutina/v3"
)

func main() {
	r := rutina.Mew(context.Background())
	r.Go(jobWithoutErrors)
	r.Go(jobWithErrorAfter10sec)
	if err := r.Wait(); err != nil {
		log.Println("Exited with errro:", err)
	}
	log.Panicln("Completed")
}

func jobWithoutErrors(ctx context.Context) error {
	<-ctx.Done()
	return nil
}

func jobWithErrorAfter10sec(ctx context.Context) error {
	<-time.After(10 * time.Second)
	return errors.New("error!")
}
