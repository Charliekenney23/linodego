package main

import (
	"log"

	golinode "github.com/chiefy/go-linode"
	"github.com/dnaeon/go-vcr/recorder"
)

func main() {
	// Start our recorder
	r, err := recorder.New("test/fixtures")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it

	c, err := golinode.NewClient(nil, r)
	if err != nil {
		log.Fatalf("Failed to create linode client: %s", err)
	}
	c.SetDebug(false)

	_, err = c.ListRegions()
	if err != nil {
		log.Fatalf("Failed to get linode regions: %s", err)
	}
	log.Println("Succesfully got linode regions")

	_, err = c.ListInstances()
	if err != nil {
		log.Fatalf("Failed to get linode instances: %s", err)
	}
	log.Println("Succesfully got linode instances")

	_, err = c.ListDistributions()
	if err != nil {
		log.Fatalf("Failed to get linode distributions: %s", err)
	}
	log.Println("Succesfully got linode distributions")

	_, err = c.GetInstance(1234)
	if err != nil {
		log.Fatalf("Failed to get linode instance ID 1234: %s", err)
	}
	log.Println("Succesfully got linode instance ID 1234")

	_, err = c.GetInstance(4090913)
	if err != nil {
		log.Fatalf("Failed to get linode instance ID 4090913: %s", err)
	}
	log.Println("Succesfully got linode instance ID 4090913")

	log.Printf("Successfully retrieved linode requests!")
}