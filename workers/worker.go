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
	args CliArgs,
) {

	scheduler := gocron.NewScheduler(time.UTC)

	_, err := scheduler.Every(30).Minutes().Do(func() {
		url := GetUrl(args.Devenv)
		created, _, err := GetAssignmentCount(url, args.Xrelayer, args.Authkey)
		if err != nil {
			log.Fatalln(err)
		}

		if created == nil {
			return
		}
		jobs, chainids := SumAssignmentsCountFromFlags(args.Chains, created)

		if args.Devenv {
			spawnWorkers(1, args.Xrelayer, args.Authkey, args.Devenv, args.Password, chainids)
		} else {
			spawnWorkers(jobs, args.Xrelayer, args.Authkey, args.Devenv, args.Password, chainids)
		}

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

func spawnWorkers(count int, xrelayer string, authkey string, devenv bool, password string, chainids []ChainIds) {
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		fmt.Println("Spawning worker ", i)
		go work(&wg, xrelayer, authkey, devenv, password, chainids[i])
		// Sleep before starting another go routine
		time.Sleep(50 * time.Millisecond)
	}

	wg.Wait()
}

func work(wg *sync.WaitGroup, xrelayer string, authkey string, devenv bool, password string, chainids ChainIds) {
	err := SolveAnAssignment(xrelayer, authkey, devenv, password, chainids)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}
