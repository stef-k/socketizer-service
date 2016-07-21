package site

import (
	"github.com/jbrodriguez/mlog"
	"github.com/astaxie/beego/orm"
)

type MainsiteSettings struct {
	Id                       int
	ServiceKey               string
	FreeKeys                 bool
	InBeta                   bool
	MaxConcurrentConnections int
}

func GetSettings() (*MainsiteSettings, error) {

	var settings MainsiteSettings
	o := orm.NewOrm()
	err := o.QueryTable("mainsite_settings").One(&settings)
	if err == orm.ErrNoRows {
		mlog.Info("could not read settings from DB")
	}
	return &settings, err
}
