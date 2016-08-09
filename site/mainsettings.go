package site

import (
	"github.com/jbrodriguez/mlog"
	"github.com/astaxie/beego/orm"
	"projects.iccode.net/stef-k/socketizer-service/models"
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

// GetAllClients returns the number of all connected clients
func GetAllClients() int {

	clientSum := 0
	for _, domain := range models.DomainPool {
		clientSum += len(domain.ClientPool)
	}
	return clientSum
}
