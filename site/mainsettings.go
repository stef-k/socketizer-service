package site

import (
	"github.com/jbrodriguez/mlog"
	"github.com/astaxie/beego/orm"
)

type MainsiteSettings struct {
	Id                       int
	ServiceIsActive          bool
	ServiceKey               string
	FreeKeys                 bool
	InBeta                   bool
	MaxConcurrentConnections int
	MaxConnection            int
}

func GetSettings() (*MainsiteSettings, error) {

	var settings MainsiteSettings
	o := orm.NewOrm()
	err := o.QueryTable("mainsite_settings").One(&settings)
	if err == orm.ErrNoRows {
		mlog.Info("no settings found in DB, please create one record")
	}
	return &settings, err
}
