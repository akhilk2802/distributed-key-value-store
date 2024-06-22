package cluster

import (
	"fmt"
	"key-value-store/store"
	"log"
	"net/http"
	"net/url"
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

// func NewCluster(LocalNode *Node, store *store.KeyValueStore) (*Cluster, error) {
// 	config := memberlist.DefaultLocalConfig()
// 	log.Println("config : ", config)
// 	config.Name = LocalNode.Name
// 	config.BindAddr = LocalNode.Addr

// 	list, err := memberlist.Create(config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Cluster{
// 		Memberlist: list,
// 		LocalNode:  LocalNode,
// 		Store:      store,
// 	}, nil
// }

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
		log.Println(addresses)
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

func (c *Cluster) BroadcastSet(key, value string) {
	for _, member := range c.Members() {
		if member.Name != c.LocalNode.Name {
			form := url.Values{}
			form.Set("key", key)
			form.Set("value", value)
			log.Println("address : ", member.Port)
			url := fmt.Sprintf("http://%s:%s/set", member.Addr, strconv.FormatUint(uint64(member.Port), 10))
			resp, err := http.PostForm(url, form)
			if err != nil {
				log.Printf("Failed to broadcast to member %s: %v", member.Name, err)
			} else {
				resp.Body.Close()
			}
		}
	}
}
