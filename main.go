package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	LoadConfig()
	setupLogger()
}

var createCmd = &cobra.Command{
	Use:   "create [events-file]",
	Short: "Create RabbitMQ queues for events",
	Long:  `Read events from a file and create cert_ and dispatch_ queues for each event`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		eventsFile := args[0]

		events, err := readEventsFile(eventsFile)
		if err != nil {
			log.Fatalf("Error reading events file: %v", err)
		}

		fmt.Printf("Found %d events\n", len(events))

		if err := createQueues(events); err != nil {
			log.Fatalf("Error creating queues: %v", err)
		}

		fmt.Println("Queue creation completed successfully")
	},
}

var processCmd = &cobra.Command{
	Use:   "process [csv-file]",
	Short: "Process CSV file and add student data to certificate queue",
	Long:  `Read student data from CSV file and publish to cert_ queue as JSON payloads`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		csvFile := args[0]

		students, err := parseCSVFile(csvFile)
		if err != nil {
			log.Fatalf("Error parsing CSV file: %v", err)
		}

		if len(students) == 0 {
			fmt.Println("No valid student records found")
			return
		}

		eventName := extractEventName(csvFile)
		queueName := fmt.Sprintf("cert_%s", eventName)

		fmt.Printf("Processing %d students for event: %s\n", len(students), eventName)
		fmt.Printf("Publishing to queue: %s\n", queueName)

		if err := publishToQueue(queueName, students); err != nil {
			log.Fatalf("Error publishing to queue: %v", err)
		}

		fmt.Println("CSV processing completed successfully")
	},
}

func main() {
	var rootCmd = &cobra.Command{Use: "sigil"}
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(processCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
