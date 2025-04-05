package commands

import (
	"fmt"
	"time"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time (1s, 1m, 1h)>")
	}

	time_between_reqs := cmd.Args[0]

	timeBetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("failed to parse duration: %w", err)
	}
	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		ScrapeFeeds(s)
	}

}
