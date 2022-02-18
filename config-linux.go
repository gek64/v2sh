package main

var (
	// 工具链
	linuxToolbox = []string{"unzip", "service", "systemd"}

	// 应用相关
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
	linuxBins         = []string{"v2ray", "v2ctl"}
	linuxBinsLocation = "/usr/local/bin/"
	linuxNeedExtract  = true

	// 配置文件相关
	linuxConfigName     = "config.json"
	linuxConfigContent  = "{}"
	linuxConfigLocation = "/usr/local/etc/v2ray/"

	// 资源文件相关
	linuxResources         = []string{"geoip.dat", "geosite.dat"}
	linuxResourcesUrl      = []string{"https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geoip.dat", "https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geosite.dat"}
	linuxResourcesLocation = "/usr/local/share/v2ray/"

	// 服务相关
	linuxServiceName    = "v2ray.service"
	linuxServiceContent = `[Unit]
Description=V2ray Service
After=network.target nss-lookup.target

[Service]
User=nobody
CapabilityBoundingSet=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
NoNewPrivileges=true
ExecStart=/usr/local/bin/v2ray -confdir /usr/local/etc/v2ray/
Restart=on-failure
RestartPreventExitStatus=23

[Install]
WantedBy=multi-user.target`
)
