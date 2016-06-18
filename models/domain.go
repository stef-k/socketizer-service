package models

import "fmt"

type Domain struct {
	Name       string   // the name of the domain i.e http://example.com
	ClientPool []*Client // Client object with key the pointer address as string
}

// NewDomain creates a new domain
func NewDomain(name string) *Domain {
	domain := new(Domain)
	domain.Name = name

	return domain
}

// AddClient adds a new client to the pool
func (d *Domain) AddClient(client *Client) {
	d.ClientPool = append(d.ClientPool, client)
}

// DeleteClient removes a client from the pool
func (d *Domain) DeleteClient(c *Client) {
	index, _ := d.FindClient(c)
	if index != -1 {
		fmt.Println("DeleteClient: ", c.Id)
		d.ClientPool = append(d.ClientPool[:index], d.ClientPool[index + 1:]...)
	}
}

// FindDomain finds a domain in the DomainPool
// and returns its index and value.
// If the domain does not exists in pool returns -1
// If the DomainPool is empty returns -1
func (d *Domain) FindClient(c *Client) (int, *Client) {
	if len(d.ClientPool) == 0 {
		return -1, &Client{}
	} else {
		for i, v := range d.ClientPool {
			if v.Id == c.Id {
				return i, v
			}
		}
		return -1, &Client{}
	}
}

// HasClient checks if the domain has a client registered in its pool
func (d *Domain) HasClient(client *Client) bool {

	for _, obj := range d.ClientPool {
		if obj.Id == client.Id {
			return true
		}
	}
	return false
}

// ListClients returns a slice with IDs of all connected clients
func (d *Domain) ListClients() []string {
	pool := make([]string, 0)
	for _, obj := range d.ClientPool {
		pool = append(pool, fmt.Sprintf("%p", obj.Id))
	}
	return pool
}

// ClientCount returns the number of domain's connected clients
func (d *Domain) ClientCount() int {
	return len(d.ClientPool)
}

// DomainBroadcast send a Message to all domain's connected clients
func (d *Domain) DomainBroadast(msg Message) {
	for _, client := range d.ClientPool {
		client.SendMessage(msg)
	}
}
