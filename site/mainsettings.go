package site

type MainSettings struct {
	FreeKeys                 bool
	InBeta                   bool
	MaxConcurrentConnections int
}

func GetSettings() *MainSettings {

	db := InitDB()
	defer db.Close()
	var settings MainSettings

	db.First(&settings)

	return &settings
}
