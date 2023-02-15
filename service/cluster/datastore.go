package cluster

type Cluster struct {
	ID     string
	Name   string
	Status string
}

var clusters = make(map[string]*Cluster)

func NewCluster() *Cluster {
	return &Cluster{
		ID:     "cls-1892893",
		Name:   "My Cluster",
		Status: "pending",
	}
}

func (c *Cluster) Save() {
	clusters[c.ID] = c
}
