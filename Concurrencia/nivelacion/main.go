package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func countErrors(files []string) int {
	var wg sync.WaitGroup
	errorChannel := make(chan int, len(files))

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			count := processFile(file)
			errorChannel <- count
		}(file)
	}

	wg.Wait()
	close(errorChannel)

	totalErrors := 0

	// Read from the channel and accumulate the errors
	for count := range errorChannel {
		totalErrors += count
	}
	return totalErrors
}

func processFile(file string) int {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", file, err)
		return 0
	}
	defer f.Close()

	errorCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			errorCount++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", file, err)
	}

	return errorCount
}

func main() {
	files := []string{"/home/jaimonetti/dev/laboraGo/Golang/Concurrencia/nivelacion/log1.txt"}
	totalErrors := countErrors(files)
	fmt.Println("Total errors:", totalErrors)
}
