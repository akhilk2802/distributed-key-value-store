# Distributed Key-Value Store

## 📌 Overview
This project is a **distributed key-value store** implemented in **Go**, utilizing the **HashiCorp Memberlist library** for **cluster membership and data replication**. Designed to be **fault-tolerant, highly available, and consistent**, this system ensures **efficient key-value storage** across multiple nodes in a distributed environment.

---

*Implemented on local machine, on multiple threads*

## 🚀 Features
- **⚡ Distributed Architecture**: Supports clustering of multiple nodes.
- **📡 Data Replication**: Ensures consistency across nodes using HashiCorp Memberlist.
- **🔄 Fault Tolerance & High Availability**: Handles node failures gracefully.
- **🌍 RESTful API**: Provides an easy-to-use interface for interaction.

---

## 🏗️ Architecture
The system consists of the following core components:

1️⃣ **Cluster**: Manages **node membership** and inter-node communication.

2️⃣ **Key-Value Store**: Provides **thread-safe operations** and integrates with the cluster for data replication.

3️⃣ **API**: Exposes **RESTful endpoints** for external clients to interact with the key-value store.

---

## Folder Structure 
```plaintext
.
├── README.md
├── api
│   └── api.go
├── cluster
│   └── cluster.go
├── go.mod
├── go.sum
├── main.go
└── store
    └── store.go

```

---

## 🚀 Usage

### **Starting the Server**
To start a node in the distributed key-value store:
```sh
go run main.go
```
You can launch multiple nodes either:
- On **different machines**.
- On the **same machine** using different ports (modify the `Addr` field in `main.go`).

---

## 🔗 API Endpoints
The following API endpoints are available to interact with the key-value store:

### **📌 Store a Key-Value Pair**
#### **PUT /put**
**Description**: Stores a key-value pair.
#### **Request Body**:
```json
{
  "key": "your-key",
  "value": "your-value"
}
```

---

### **📌 Retrieve a Value**
#### **GET /get/key=your-key**
**Description**: Retrieves the value associated with a given key.
#### **Query Parameters**:
- `key`: The key to retrieve.

---

### **📌 Delete a Key-Value Pair**
#### **DELETE /delete?key=your-key**
**Description**: Deletes the specified key-value pair.
#### **Query Parameters**:
- `key`: The key to delete.

---

### **📌 Store a Key-Value Pair via URL Parameters**
#### **POST /set/{key}/{value}**
**Description**: Stores a key-value pair using URL parameters.
#### **URL Parameters**:
- `key`: The key to store.
- `value`: The corresponding value.

---

## 🔄 Data Replication
**Replication** is achieved using the **HashiCorp Memberlist library**, ensuring consistency across all nodes in the distributed store. 

### **🔍 How Data Replication Works**
1️⃣ **Cluster Initialization**: Each node initializes a `Cluster` instance with a unique name and address.

2️⃣ **Joining the Cluster**: Nodes can join an existing cluster by specifying the addresses of existing nodes.

3️⃣ **Broadcasting Updates**: Any changes (create/update/delete) to a key-value pair are serialized and broadcast across all nodes.

4️⃣ **Handling Incoming Messages**: Each node listens for incoming updates and synchronizes its local key-value store accordingly.

---
🚀 **Distributed Key-Value Store - Scalable, Fault-Tolerant, and Highly Available!**

