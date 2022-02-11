package main

var (
	// 工具链
	linuxToolbox = []string{"unzip", "service", "systemd"}

	// 代理文件
	linuxRepo     = "v2fly/v2ray-core"
	linuxRepoList = map[string]string{
		"linux_386":      "v2ray-linux-32.zip",
		"linux_amd64":    "v2ray-linux-64.zip",
		"linux_arm":      "v2ray-linux-arm32-v7a.zip ",
		"linux_arm64":    "v2ray-linux-arm64-v8a.zip",
		"linux_mips":     "v2ray-linux-mips32.zip",
		"linux_mips64":   "v2ray-linux-mips64.zip",
		"linux_mipsle":   "v2ray-linux-mips32le.zip",
		"linux_mips64le": "v2ray-linux-mips64le.zip",
		"linux_riscv64":  "v2ray-linux-riscv64.zip",
	}
	// 资源文件
	linuxResources = []string{"https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geoip.dat", "https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geosite.dat"}
	// 服务名称
	linuxServiceName = "v2ray.service"
	// 服务内容
	linuxServiceContent = `[Unit]
Description=V2ray Service
After=network.target nss-lookup.target

[Service]
User=nobody
CapabilityBoundingSet=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
NoNewPrivileges=true
ExecStart=/usr/local/bin/proxy/v2ray -confdir /usr/local/bin/proxy/
Restart=on-failure
RestartPreventExitStatus=23

[Install]
WantedBy=multi-user.target`
)
