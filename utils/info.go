package utils

import "projects.iccode.net/stef-k/socketizer-service/models"

// Return total Domains and total Clients
func TotalCons() (int, int) {
	clientSum := 0
	_, d := models.ListDomains()
	for _, domain := range models.DomainPool {
		clientSum += len(domain.ClientPool)
	}
	return d, clientSum
}
