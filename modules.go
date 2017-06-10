package main

import (
	_ "github.com/kth/logspout/adapters/raw"
	_ "github.com/kth/logspout/adapters/syslog"
	_ "github.com/kth/logspout/httpstream"
	_ "github.com/kth/logspout/routesapi"
	_ "github.com/kth/logspout/transports/tcp"
	_ "github.com/kth/logspout/transports/tls"
	_ "github.com/kth/logspout/transports/udp"
)
