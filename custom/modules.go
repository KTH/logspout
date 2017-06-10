package main

import (
	_ "github.com/kth/logspout/adapters/syslog"
	_ "github.com/kth/logspout/transports/tcp"
	_ "github.com/kth/logspout/transports/tls"
	_ "github.com/kth/logspout/transports/udp"
	_ "github.com/looplab/logspout-logstash"
)
