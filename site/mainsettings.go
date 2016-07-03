package site

type MainsiteSettings struct {
	ServiceKey               string
	FreeKeys                 bool
	InBeta                   bool
	MaxConcurrentConnections int
}

func GetSettings() *MainsiteSettings {

	db := InitDB()
	defer db.Close()
	var settings MainsiteSettings

	db.First(&settings)

	return &settings
}
