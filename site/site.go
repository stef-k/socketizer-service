package site

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

func InitDB() *gorm.DB {
	//db, err := gorm.Open("postgres", "socketizer=h5epb4N1shOz5i0AzqQN9zyxzBDMkdavJsTvuUBIui4WjFAIBt@tcp(127.0.0.1:5333)/socketizer")
	db, err := gorm.Open("postgres", "postgresql://socketizer:h5epb4N1shOz5i0AzqQN9zyxzBDMkdavJsTvuUBIui4WjFAIBt@127.0.0.1:5433")

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	db.SingularTable(true)
	return db
}
