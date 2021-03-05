// +build ignore

package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/neonxp/rutina/v3"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	r := rutina.Mew(ctx)
	r.Go(jobWithoutErrors)
	if err := r.Wait(); err != nil {
		log.Println("Exited with errro:", err)
	}
	log.Panicln("Completed")
}

func jobWithoutErrors(ctx context.Context) error {
	log.Println("Job started")
	<-ctx.Done()
	log.Println("Job stopped")
	return nil
}
