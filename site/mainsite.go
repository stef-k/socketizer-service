package site

import (
	"github.com/jbrodriguez/mlog"
	"github.com/astaxie/beego/orm"
)

type MainsiteDomain struct {
	Id                       int
	Domain                   string
	ApiKey                   string
	DaysLeft                 int
	MaxConcurrentConnections int
	CurrentMonthApiCalls     int
}

// GetDomain returns the domain identified by it's API key
func FindDomainByApiKey(apiKey string) (*MainsiteDomain, error) {
	var domain MainsiteDomain
	o := orm.NewOrm()
	err := o.QueryTable("mainsite_domain").Filter("api_key", apiKey).One(&domain)
	if err == orm.ErrNoRows {
		mlog.Info("could not find domain with such API key")
	}
	return &domain, err
}

func FindDomainByName(name string) (*MainsiteDomain, error) {
	var domain MainsiteDomain
	o := orm.NewOrm()
	err := o.QueryTable("mainsite_domain").Filter("domain", name).One(&domain)
	if err == orm.ErrNoRows {
		mlog.Info("could not find domain with such name")
	}
	return &domain, err
}

// Check if the domain one way or another is active
func (d MainsiteDomain) IsActive() bool {
	return d.DaysLeft > 0
}

func (d MainsiteDomain) GetMaxConcurrentConnections() int {
	return d.MaxConcurrentConnections
}
