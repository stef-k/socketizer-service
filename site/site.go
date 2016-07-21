package site

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	orm.RegisterModel(new(MainsiteSettings))
	orm.RegisterModel(new(MainsiteDomain))

	orm.RegisterDataBase("default", "postgres", "postgresql://socketizer:h5epb4N1shOz5i0AzqQN9zyxzBDMkdavJsTvuUBIui4WjFAIBt@127.0.0.1:5432", 30)
}
