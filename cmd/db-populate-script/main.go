package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

func main() {
	const numRecords = 1000000
	const fileName = "tarefas.csv"
	const diaAtividade = "2024-01-01"
	const numWorkers = 2 // Optimize for t3.micro

	startTime := time.Now()

	// Open the file for writing
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Channel for sending tasks
	recordsChan := make(chan []string, 100) // Buffered channel

	// Wait group to ensure all workers finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for i := workerID; i < numRecords; i += numWorkers {
				uuid := uuid.New().String()
				titulo := fmt.Sprintf("Tarefa %d", i)
				importante := rand.Intn(2) == 0 // Random boolean
				now := time.Now().Format(time.RFC3339)
				record := []string{
					uuid,
					titulo,
					diaAtividade,
					fmt.Sprintf("%t", importante),
					now,
					now,
				}
				recordsChan <- record
			}
		}(w)
	}

	// Goroutine to close the channel once all workers are done
	go func() {
		wg.Wait()
		close(recordsChan)
	}()

	// Write records from the channel to the CSV
	for record := range recordsChan {
		if err := writer.Write(record); err != nil {
			fmt.Printf("Error writing record: %v\n", err)
			return
		}
	}

	fmt.Printf("CSV file '%s' generated successfully with %d records in %s.\n", fileName, numRecords, time.Since(startTime))
}
