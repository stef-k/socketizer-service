package site

import (
	"github.com/astaxie/beego/orm"
	"github.com/jbrodriguez/mlog"
)

// Stats model to hold some basic application - server stats
type MainsiteStatistics struct {
	Id                       int
	MaxConcurrentClients     int
	MaxConcurrentDomains     int
	TotalMessagesBroadcasted int64
	TotalClients             int64
}

//IncreaseTotalClientsBy increases the total of all time connected clients
func IncreaseTotalClientsBy(number int) {
	var stats MainsiteStatistics
	o := orm.NewOrm()

	err := o.QueryTable("mainsite_statistics").One(&stats)
	if err != nil {
		if err == orm.ErrNoRows {
			mlog.Info("no Stats record found in DB, please create one record")
		} else {
			mlog.Info("could not access stats table, ", err.Error())
		}
	}

	if err != nil {
		mlog.Warning("could not update TotalClients stats")
	} else {
		var client int64
		client = int64(number)
		o := orm.NewOrm()
		stats.TotalClients = stats.TotalClients + client
		_, e := o.Update(&stats, "TotalClients")
		if e != nil {
			mlog.Warning("could not update TotalClients")
		}
	}
}

// UpdateTotalMessagesBroadcasted updates the total number of all time messages broadcasted
func UpdateTotalMessagesBroadcasted()  {
	var stats MainsiteStatistics
	o := orm.NewOrm()

	err := o.QueryTable("mainsite_statistics").One(&stats)

	if err != nil {
		if err == orm.ErrNoRows {
			mlog.Info("no Stats record found in DB, please create one record")
		} else {
			mlog.Info("could not access stats table, ", err.Error())
		}
	}

	if err != nil {
		mlog.Warning("Could not update TotalMessagesBroadcasted stats")
	} else {
		stats.TotalMessagesBroadcasted += 1
		_, e := o.Update(&stats, "TotalMessagesBroadcasted")
		if e != nil {
			mlog.Warning("could not update TotalMessagesBroadcasted")
		}
	}
}

//UpdateMaxConcurrentClients updates the all time high of concurrent clients
func UpdateMaxConcurrentClients(clientNumber int) {
	var stats MainsiteStatistics
	o := orm.NewOrm()

	err := o.QueryTable("mainsite_statistics").One(&stats)
	if err != nil {
		if err == orm.ErrNoRows {
			mlog.Info("no Stats record found in DB, please create one record")
		} else {
			mlog.Info("could not access stats table, ", err.Error())
		}
	}

	if err != nil {
		mlog.Warning("Could not update MaxConcurrentClients stats")
	} else {
		if stats.MaxConcurrentClients < clientNumber {
			o := orm.NewOrm()
			stats.MaxConcurrentClients = clientNumber
			_, e := o.Update(&stats, "MaxConcurrentClients")
			if e != nil {
				mlog.Warning("could not update MaxConcurrentClients")
			}
		}
	}
}

//UpdateMaxConcurrentDomains updates the all time high of concurrent clients
func UpdateMaxConcurrentDomains(domainNumber int) {
	var stats MainsiteStatistics
	o := orm.NewOrm()

	err := o.QueryTable("mainsite_statistics").One(&stats)
	if err != nil {
		if err == orm.ErrNoRows {
			mlog.Info("no Stats record found in DB, please create one record")
		} else {
			mlog.Info("could not access stats table, ", err.Error())
		}
	}

	if err != nil {
		mlog.Warning("Could not update MaxConcurrentClients stats")
	} else {
		if stats.MaxConcurrentDomains < domainNumber {
			o := orm.NewOrm()
			stats.MaxConcurrentDomains = domainNumber
			_, e := o.Update(&stats, "MaxConcurrentDomains")
			if e != nil {
				mlog.Warning("could not update MaxConcurrentDomains")
			}
		}
	}
}
