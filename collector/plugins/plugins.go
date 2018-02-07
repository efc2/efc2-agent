package plugins

import (
	// registry all plugins
	_ "github.com/efc2/efc2-agent/collector/plugins/apache"
	_ "github.com/efc2/efc2-agent/collector/plugins/docker"
	_ "github.com/efc2/efc2-agent/collector/plugins/haproxy"
	_ "github.com/efc2/efc2-agent/collector/plugins/memcached"
	_ "github.com/efc2/efc2-agent/collector/plugins/mongodb"
	_ "github.com/efc2/efc2-agent/collector/plugins/mysql"
	_ "github.com/efc2/efc2-agent/collector/plugins/nginx"
	_ "github.com/efc2/efc2-agent/collector/plugins/phpfpm"
	_ "github.com/efc2/efc2-agent/collector/plugins/postgres"
	_ "github.com/efc2/efc2-agent/collector/plugins/redis"
	_ "github.com/efc2/efc2-agent/collector/plugins/system"
)
