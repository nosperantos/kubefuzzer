package kubefuzzer

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// getClientset initializes and returns a Kubernetes clientset
func getClientset() (*kubernetes.Clientset, error) {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		return nil, fmt.Errorf("could not find kubeconfig")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

// readEndpoints reads endpoints from a file and returns them as a slice of strings
func readEndpoints(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var endpoints []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		endpoints = append(endpoints, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return endpoints, nil
}

// splitEndpoints splits the endpoints into 20 chunks
func splitEndpoints(endpoints []string) [][]string {
	chunkSize := (len(endpoints) + 19) / 20
	var chunks [][]string
	for i := 0; i < len(endpoints); i += chunkSize {
		end := i + chunkSize
		if end > len(endpoints) {
			end = len(endpoints)
		}
		chunks = append(chunks, endpoints[i:end])
	}
	return chunks
}

// callEndpoint simulates calling an endpoint
func callEndpoint(endpoint string) {
	// Simulate the call with a print statement
	fmt.Println("Calling endpoint:", endpoint)
}

// processChunk processes a chunk of endpoints
func processChunk(chunk []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, endpoint := range chunk {
		callEndpoint(endpoint)
	}
}

// main function to read endpoints, split them into chunks, and process them with threads
func main() {
	// Define a command line flag for the dictionary filename
	filePath := flag.String("file", "", "Path to the endpoints file")
	flag.Parse()

	// Check if the file path is provided
	if *filePath == "" {
		fmt.Println("Usage: kubefuzzer -file <path_to_endpoints_file>")
		os.Exit(1)
	}

	endpoints, err := readEndpoints(*filePath)
	if err != nil {
		log.Fatalf("Failed to read endpoints: %v", err)
	}

	chunks := splitEndpoints(endpoints)
	var wg sync.WaitGroup

	for _, chunk := range chunks {
		wg.Add(1)
		go processChunk(chunk, &wg)
	}

	wg.Wait()
}
