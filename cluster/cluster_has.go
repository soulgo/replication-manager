// replication-manager - Replication Manager Monitoring and CLI for MariaDB and MySQL
// Copyright 2017 Signal 18 SARL
// Authors: Guillaume Lefranc <guillaume@signal18.io>
//          Stephane Varoqui  <svaroqui@gmail.com>
// This source code is licensed under the GNU General Public License, version 3.

package cluster

import "strings"

func (cluster *Cluster) IsInIgnoredHosts(host string) bool {
	ihosts := strings.Split(cluster.conf.IgnoreSrv, ",")
	for _, ihost := range ihosts {
		if host == ihost {
			return true
		}
	}
	return false
}

func (cluster *Cluster) IsInPreferedHosts(host string) bool {
	ihosts := strings.Split(cluster.conf.PrefMaster, ",")
	for _, ihost := range ihosts {
		if host == ihost {
			return true
		}
	}
	return false
}

func (cluster *Cluster) IsProvision() bool {
	for _, s := range cluster.servers {
		if s.State == stateFailed || s.State == stateSuspect /*&& misc.Contains(cluster.ignoreList, s.URL) == false*/ {
			return false
		}
	}
	return true
}

func (cluster *Cluster) IsInHostList(host string) bool {
	for _, v := range cluster.hostList {
		if v == host {
			return true
		}
	}
	return false
}

func (cluster *Cluster) IsMasterFailed() bool {
	// get real master or the virtual master
	mymaster := cluster.GetMaster()
	if mymaster == nil {
		return true
	}
	if mymaster.State == stateFailed {
		return true
	} else {
		return false
	}
}

func (cluster *Cluster) IsActive() bool {
	if cluster.runStatus == "A" {
		return true
	} else {
		return false
	}
}

func (cluster *Cluster) IsVerbose() bool {
	if cluster.conf.Verbose {
		return true
	} else {
		return false
	}
}

func (cluster *Cluster) IsInFailover() bool {
	return cluster.sme.IsInFailover()
}

func (cluster *Cluster) IsDiscovered() bool {
	return cluster.sme.IsDiscovered()
}
