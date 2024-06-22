package main

import (
	"flag"
	"key-value-store/api"
	"key-value-store/cluster"
	"key-value-store/store"
	"log"
)

func main() {

	// log.Println("starting application")
	// node := &cluster.Node{Name: "node2", Addr: "localhost:8080"}
	// log.Println("Node Created")
	// store := store.NewKeyValueStore()
	// log.Println("Store Created")
	// cluster, err := cluster.NewCluster(node, store)
	// if err != nil {
	// 	log.Fatalf("Failed to create cluster: %v", err)
	// }

	// api := api.NewAPI(cluster)
	// log.Println("API Created")
	// api.Run(":8081")
	// log.Println("Started and running")

	nodeName := flag.String("node", "node1", "Name of the node")
	nodeAddr := flag.String("addr", "127.0.0.1:7946", "Address of the node")
	apiPort := flag.String("port", ":8081", "Port for the API to listen on")
	clusterAddr := flag.String("cluster", "", "Comma-separated list of existing cluster members")

	// Parse the flags
	flag.Parse()

	// Start the application
	log.Println("Starting application")

	// Create node
	node := &cluster.Node{Name: *nodeName, Addr: *nodeAddr}
	log.Println("Node created:", node)

	// Create key-value store
	store := store.NewKeyValueStore()
	log.Println("Store created")

	// Create cluster
	cluster, err := cluster.NewCluster(node, store, *clusterAddr)
	if err != nil {
		log.Fatalf("Failed to create cluster: %v", err)
	}

	// Create API
	api := api.NewAPI(cluster)
	log.Println("API created")

	// Run the API
	api.Run(*apiPort)
	log.Println("Started and running on port", *apiPort)
}
