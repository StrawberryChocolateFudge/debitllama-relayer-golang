package workers

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
)

func StartWorkers(
	devenv bool,
	xrelayer string,
	authkey string,
	password string,
) {

	scheduler := gocron.NewScheduler(time.UTC)

	_, err := scheduler.Every(30).Minutes().Do(func() {
		url := GetUrl(devenv)
		created, _, err := GetAssignmentCount(url, xrelayer, authkey)
		if err != nil {
			log.Fatalln(err)
		}

		spawnWorkers(created, xrelayer, authkey, devenv)
	})

	if err != nil {
		log.Fatalln(err)
	}

	scheduler.StartAsync()

	// Graceful shutdown on sigterm
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit
	fmt.Printf("Received signal: %s\n", sig)
	fmt.Println("Server gracefully shutting down. Waiting for running jobs to finish...")
	scheduler.Stop()

}

func spawnWorkers(count int, xrelayer string, authkey string, devenv bool) {
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		fmt.Println("Spawning worker ", i)
		go work(&wg, xrelayer, authkey, devenv)
		// Sleep before starting another go routine
		time.Sleep(50 * time.Millisecond)
	}

	wg.Wait()
}

func work(wg *sync.WaitGroup, xrelayer string, authkey string, devenv bool) {
	err := SolveAnAssignment(xrelayer, authkey, devenv)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}
