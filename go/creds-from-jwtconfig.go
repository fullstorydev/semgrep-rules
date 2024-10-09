package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func main() {
	authUsingJWTConfigFromJSON()

	authUsingADC()
}

func authUsingJWTConfigFromJSON() {
	data, err := ioutil.ReadFile("path/to/service-account-key.json")
	if err != nil {
		log.Fatalf("Unable to read service account key file: %v", err)
	}

	// ruleid: creds-from-jwtconfig
	jwtConfig, err := google.JWTConfigFromJSON(data, compute.CloudPlatformScope)
	if err != nil {
		log.Fatalf("Unable to parse service account key file to JWTConfig: %v", err)
	}

	client := jwtConfig.Client(context.Background())

	computeService, err := compute.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Compute service: %v", err)
	}

	instancesList, err := computeService.Instances.List("your-project-id", "us-central1-a").Do()
	if err != nil {
		log.Fatalf("Unable to list instances: %v", err)
	}

	for _, instance := range instancesList.Items {
		fmt.Printf("Instance: %s\n", instance.Name)
	}
}

func authUsingADC() {
	computeService, err := compute.NewService(context.Background(), option.WithScopes(compute.CloudPlatformScope))
	if err != nil {
		log.Fatalf("Unable to create Compute service: %v", err)
	}

	instancesList, err := computeService.Instances.List("your-project-id", "us-central1-a").Do()
	if err != nil {
		log.Fatalf("Unable to list instances: %v", err)
	}

	for _, instance := range instancesList.Items {
		fmt.Printf("Instance: %s\n", instance.Name)
	}
}
