package site

import (
	"github.com/astaxie/beego/orm"
	"github.com/jbrodriguez/mlog"
)

// Stats model to hold some basic application - server stats
type Stats struct {
	TotalClients         int64
	MaxConcurrentClients int
}

func GetStats()  (*Stats, error){
	var stats Stats
	o := orm.NewOrm()

	err := o.QueryTable("mainsite_stats").One(&stats)
	if err == orm.ErrNoRows {
		mlog.Info("could not read stats from DB")
	}
	return &stats, err
}

//IncreaseTotalClientsBy increases the total of all time connected clients
func IncreaseTotalClientsBy(number int)  {
	stats, err := GetStats()

	if err != nil {
		mlog.Warning("Could not update TotalClients stats")
	} else {
		var client int64
		client = int64(number)
		o := orm.NewOrm()
		stats.TotalClients = client
		o.Update(&stats, "TotalClients")
	}
}

//UpdateMaxConcurrentClients updates the all time high of concurrent clients
func UpdateMaxConcurrentClients(clientNumber int)  {
	stats, err := GetStats()

	if err != nil {
		mlog.Warning("Could not update MaxConcurrentClients stats")
	} else {
		if stats.MaxConcurrentClients < clientNumber {
			o := orm.NewOrm()
			stats.MaxConcurrentClients = clientNumber
			o.Update(&stats, "MaxConcurrentClients")
		}
	}
}
