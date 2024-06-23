package cluster

import (
	"key-value-store/store"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/memberlist"
)

type Node struct {
	Name string
	Addr string
}

type Cluster struct {
	*memberlist.Memberlist
	LocalNode *Node
	Store     *store.KeyValueStore
}

func NewCluster(node *Node, store *store.KeyValueStore, clusterAddress string) (*Cluster, error) {
	config := memberlist.DefaultLocalConfig()
	config.Name = node.Name
	config.BindAddr = strings.Split(node.Addr, ":")[0]
	bindPort := strings.Split(node.Addr, ":")[1]
	config.BindPort, _ = strconv.Atoi(bindPort)

	memberList, err := memberlist.Create(config)
	if err != nil {
		return nil, err
	}

	cluster := &Cluster{
		Memberlist: memberList,
		LocalNode:  node,
		Store:      store,
	}

	if clusterAddress != "" {
		addresses := strings.Split(clusterAddress, ",")
		log.Println(" CLuster addresses : ", addresses)
		err := cluster.Join(addresses)
		if err != nil {
			return nil, err
		}
		log.Println("Cluster joined:", addresses)
	}
	log.Println("Cluster created with members:", cluster.Members())

	return cluster, nil
}

func (c *Cluster) Join(seeds []string) error {
	log.Println("seeds :", seeds)
	_, err := c.Memberlist.Join(seeds)
	if err != nil {
		log.Printf("Failed to join cluster: %v", err)
		return err
	}
	log.Printf("Node joined the cluster: %s", c.LocalNode.Name)
	return nil
}

// func (c *Cluster) NotifyMsg(msg []byte) {

// }
