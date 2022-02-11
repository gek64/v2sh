package main

var (
	// 工具链
	freebsdToolbox = []string{"unzip", "service"}

	// 代理文件
	freebsdRepo     = linuxRepo
	freebsdRepoList = map[string]string{
		"freebsd_386":   "v2ray-freebsd-32.zip",
		"freebsd_amd64": "v2ray-freebsd-64.zip",
	}
	// 资源文件
	freebsdResources = linuxResources
	// 服务名称
	freebsdServiceName = "v2ray"
	// 服务内容
	freebsdServiceContent = `#!/bin/sh
# PROVIDE: v2ray
# REQUIRE: DAEMON NETWORKING
#
# Install rc.d service use the following command on freebsd
#
# mkdir /usr/local/etc/rc.d/
# ee /usr/local/etc/rc.d/v2ray && chmod +x /usr/local/etc/rc.d/v2ray
# service v2ray enable && service v2ray start

. /etc/rc.subr

name=v2ray
rcvar=${name}_enable

command="/usr/local/bin/proxy/${name}"
pidfile="/var/run/${name}.pid"

start_cmd="${name}_start"
stop_cmd="${name}_stop"

v2ray_start() {
  echo "Starting ${name}."
  /usr/sbin/daemon -cf -p ${pidfile} ${command} run -d "/usr/local/bin/proxy/"
}

v2ray_stop() {
  echo "Stopping ${name}."
  if [ -f ${pidfile} ]; then
    pkill -F ${pidfile}
  fi
}

load_rc_config $name
run_rc_command "$1"
`
)
