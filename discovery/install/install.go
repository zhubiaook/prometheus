// Copyright 2020 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package install has the side-effect of registering all builtin
// service discovery config types.
package install

import (
	_ "github.com/zhubiaook/prometheus/discovery/aws"          // register aws
	_ "github.com/zhubiaook/prometheus/discovery/azure"        // register azure
	_ "github.com/zhubiaook/prometheus/discovery/consul"       // register consul
	_ "github.com/zhubiaook/prometheus/discovery/digitalocean" // register digitalocean
	_ "github.com/zhubiaook/prometheus/discovery/dns"          // register dns
	_ "github.com/zhubiaook/prometheus/discovery/eureka"       // register eureka
	_ "github.com/zhubiaook/prometheus/discovery/file"         // register file
	_ "github.com/zhubiaook/prometheus/discovery/gce"          // register gce
	_ "github.com/zhubiaook/prometheus/discovery/hetzner"      // register hetzner
	_ "github.com/zhubiaook/prometheus/discovery/http"         // register http
	_ "github.com/zhubiaook/prometheus/discovery/ionos"        // register ionos
	_ "github.com/zhubiaook/prometheus/discovery/kubernetes"   // register kubernetes
	_ "github.com/zhubiaook/prometheus/discovery/linode"       // register linode
	_ "github.com/zhubiaook/prometheus/discovery/marathon"     // register marathon
	_ "github.com/zhubiaook/prometheus/discovery/moby"         // register moby
	_ "github.com/zhubiaook/prometheus/discovery/nomad"        // register nomad
	_ "github.com/zhubiaook/prometheus/discovery/openstack"    // register openstack
	_ "github.com/zhubiaook/prometheus/discovery/ovhcloud"     // register ovhcloud
	_ "github.com/zhubiaook/prometheus/discovery/puppetdb"     // register puppetdb
	_ "github.com/zhubiaook/prometheus/discovery/scaleway"     // register scaleway
	_ "github.com/zhubiaook/prometheus/discovery/triton"       // register triton
	_ "github.com/zhubiaook/prometheus/discovery/uyuni"        // register uyuni
	_ "github.com/zhubiaook/prometheus/discovery/vultr"        // register vultr
	_ "github.com/zhubiaook/prometheus/discovery/xds"          // register xds
	_ "github.com/zhubiaook/prometheus/discovery/zookeeper"    // register zookeeper
)
