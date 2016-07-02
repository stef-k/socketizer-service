package site

type MainsiteDomain struct {
	Domain                   string
	ApiKey                   string
	FreeKey                  bool
	DaysLeft                 int
	MaxConcurrentConnections int
	CurrentMonthApiCalls     int
}

// GetDomain returns the domain identified by it's API key
func FindDomainByApiKey(apiKey string) *MainsiteDomain {
	db := InitDB()
	defer db.Close()

	var domain MainsiteDomain

	db.First(&domain, "api_key = ?", apiKey)
	return &domain
}

func FindDomainByName(name string) *MainsiteDomain {
	db := InitDB()
	defer db.Close()

	var domain MainsiteDomain
	db.Find(&domain, MainsiteDomain{Domain: name})
	return &domain
}

// Check if the domain one way or another is active
func (d MainsiteDomain) IsActive() bool {
	return d.FreeKey || d.DaysLeft > 0
}

func (d MainsiteDomain) GetMaxConcurrentConnections() int {
	return d.MaxConcurrentConnections
}
