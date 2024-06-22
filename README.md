# Distributed Key-Value Store

This project is a distributed key-value store implemented in Go, leveraging the HashiCorp Memberlist library for cluster membership and data replication. The goal is to provide a fault-tolerant, highly available, and consistent key-value store.

## Features

- Distributed architecture with node clustering.
- Data replication across multiple nodes.
- Fault tolerance and high availability.
- RESTful API for interacting with the key-value store.

## Architecture

The system is composed of the following components:

- **Cluster**: Manages node membership and communication between nodes.
- **Key-Value Store**: Stores data with thread-safe operations and integrates with the cluster for data replication.
- **API**: Provides a RESTful interface for clients to interact with the key-value store.


## Usage

Starting the Server
To start a node in the distributed key-value store:

- go run main.go


You can start multiple nodes on different machines or on the same machine using different ports by modifying the Addr field in main.go.

## API Endpoints
The following API endpoints are available to interact with the key-value store:

### PUT /put

Description: Store a key-value pair.
Request Body:
{
  "key": "your-key",
  "value": "your-value"
}

### GET /get/key=your-key

Description: Retrieve the value for a given key.
Query Parameters:
key: The key to retrieve.


### DELETE /delete?key=your-key

Description: Delete a key-value pair.
Query Parameters:
key: The key to delete.

### POST /set/{key}/{value}

Description: Store a key-value pair via URL parameters.
URL Parameters:
key: The key to store.
value: The value to store.

## Data Replication
Data replication is achieved using the HashiCorp Memberlist library. When a key-value pair is stored, updated, or deleted, the change is broadcast to all nodes in the cluster, ensuring consistency across the distributed store.

#### Example Implementation
Here's a high-level overview of how data replication is implemented:

- Cluster Initialization: Each node initializes a Cluster instance with a unique name and address.
- Joining the Cluster: Nodes can join an existing cluster by specifying the addresses of existing nodes.
- Broadcasting Updates: When a key-value pair is modified, the change is serialized and broadcast to all other nodes in the cluster.
- Handling Incoming Messages: Each node listens for incoming messages and updates its local key-value store accordingly.
