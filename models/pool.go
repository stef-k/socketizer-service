package models


// The main pool containing all domains currently active
var DomainPool = make([]*Domain, 0)

func AddDomain(d *Domain)  {
	DomainPool = append(DomainPool, d)
}

// FindDomain finds a domain in the DomainPool
// and returns its index and value.
// If the domain does not exists in pool returns -1
// If the DomainPool is empty returns -1
func FindDomain(d string) (int, *Domain) {
	for i, v := range DomainPool {
		if v.Name == d {
			return i, v
		}
	}
	return -1, &Domain{}
}

// RemoveDomain removes domain from the DomainPool
func RemoveDomain(d *Domain) {

	index, _ := FindDomain(d.Name)

	if index != -1 {
		DomainPool = append(DomainPool[:index], DomainPool[index + 1:]...)
	}
}

func RemoveClient(c *Client)  {
	for _, domain := range DomainPool{
		if domain.Name == c.Domain {
			if domain.HasClient(c) {
				domain.DeleteClient(c)
				if domain.ClientCount() == 0 {
					RemoveDomain(domain)
				}
			}
		}
	}
}

// ListDomains return the number of registered domains and
// a slice of strings with all domain names in pool
func ListDomains() (int, []string) {
	pool := make([]string, 0)
	for _, obj := range DomainPool {
		pool = append(pool, obj.Name)
	}
	return len(DomainPool), pool
}

// DomainCount return the number of curretly registered domains in pool
func DomainCount() int  {
	return len(DomainPool)
}

// PoolBroadcast send a message to all clients of the Pool
func PoolBroadcast(m Message)  {
	for _, domain := range DomainPool {
		domain.DomainBroadast(m)
	}
}
