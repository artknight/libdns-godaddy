package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	godaddy "github.com/caoyongzheng/libdns-godaddy"
	"github.com/libdns/libdns"
)

func main() {
	token := os.Getenv("GODADDY_TOKEN")
	if token == "" {
		fmt.Printf("DNSPOD_TOKEN not set\n")
		return
	}
	zone := os.Getenv("ZONE")
	if zone == "" {
		fmt.Printf("ZONE not set\n")
		return
	}
	provider := godaddy.Provider{
		APIToken: token,
	}

	records, err := provider.GetRecords(context.TODO(), zone)
	if err != nil {
		log.Fatalln("ERROR: %s\n", err.Error())
	}

	testName := "libdns-test"
	hasTestName := false

	for _, record := range records {
		fmt.Printf("%s (.%s): %s, %s\n", record.Name, zone, record.Value, record.Type)
		if record.Name == testName {
			hasTestName = true
		}
	}

	if !hasTestName {
		appendedRecords, err := provider.AppendRecords(context.TODO(), zone, []libdns.Record{
			libdns.Record{
				Type: "TXT",
				Name: testName,
				TTL:  time.Duration(600) * time.Second,
			},
		})

		if err != nil {
			log.Fatalln("ERROR: %s\n", err.Error())
		}

		fmt.Println("appendedRecords")
		fmt.Println(appendedRecords)
	} else {
		deleteRecords, err := provider.DeleteRecords(context.TODO(), zone, []libdns.Record{
			libdns.Record{
				Type: "TXT",
				Name: testName,
			},
		})

		if err != nil {
			log.Fatalln("ERROR: %s\n", err.Error())
		}

		fmt.Println("deleteRecords")
		fmt.Println(deleteRecords)
	}
}
