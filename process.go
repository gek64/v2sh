package main

import (
	"fmt"
	"gek_file"
	"os"
)

var (
	// 工具链
	toolbox = []string{"unzip", "systemctl", "systemd"}
	// 应用安装目录
	installLocation = "/usr/local/bin/proxy/"
	// 应用下载文件存储的临时目录
	tempLocation = "/tmp/proxy_installer/"
	// 应用下载文件中需要的解压的文件
	extractFiles = []string{"v2ray", "v2ctl", "geoip.dat", "geosite.dat"}
	// 目标文件名
	targetFiles = []string{"v2ray", "v2ctl"}
	// 代理文件
	proxyRepo = "v2fly/v2ray-core"
	proxyList = map[string]string{
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
	resourcesList = []string{"https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geoip.dat", "https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geosite.dat"}
	// 服务名称
	serviceName = "proxy.service"
	// 服务内容
	serviceContent = `[Unit]
Description=Proxy Service
After=network.target nss-lookup.target

[Service]
User=nobody
CapabilityBoundingSet=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
NoNewPrivileges=true
ExecStart=/usr/local/bin/proxy/v2ray -config /usr/local/bin/proxy/config.json
Restart=on-failure
RestartPreventExitStatus=23

[Install]
WantedBy=multi-user.target`
)

// 主要功能函数
// 安装
func install() (err error) {
	// 需要应用未安装
	exist, _, _ := gek_file.Exist(installLocation + targetFiles[0])
	if exist {
		return fmt.Errorf("app is already installed")
	}

	// 如果已经指定本地文件则不进行下载
	if cliLocal == "" {
		// 下载应用
		err = downloadApp(proxyRepo, proxyList, tempLocation)
		if err != nil {
			return err
		}
	}
	// 处理应用
	err = process()
	if err != nil {
		return err
	}

	// 安装配置文件
	err = installConfig(installLocation, cliConfig)
	if err != nil {
		return err
	}
	// 安装服务
	err = installService(serviceName, serviceContent)
	if err != nil {
		return err
	}
	return nil
}

// 处理应用下载文件
func process() (err error) {
	// 如果不存在安装文件夹,则创建
	exist, _, _ := gek_file.Exist(installLocation)
	if !exist {
		err = gek_file.CreateDir(installLocation)
		if err != nil {
			return err
		}
	}

	// 解压文件到安装文件夹
	err = extract(tempLocation+"v2ray*.zip", extractFiles, installLocation)
	if err != nil {
		return err
	}

	// 可执行文件赋权755
	err = chmod(targetFiles, 755)
	if err != nil {
		return err
	}

	// 删除临时文件夹
	err = os.RemoveAll(tempLocation)
	if err != nil {
		return err
	}
	return nil
}

// 卸载
func uninstall() (err error) {
	// 重置服务
	err = resetApp(installLocation)
	if err != nil {
		return
	}
	// 卸载服务
	err = uninstallService(serviceName, serviceContent)
	if err != nil {
		return err
	}
	return nil
}

// 更新
func update() (err error) {
	// 需要应用已安装
	exist, _, _ := gek_file.Exist(installLocation + targetFiles[0])
	if !exist {
		return fmt.Errorf("app is not installed")
	}

	// 如果已经指定本地文件则不进行下载
	if cliLocal == "" {
		// 下载应用
		err = downloadApp(proxyRepo, proxyList, tempLocation)
		if err != nil {
			return err
		}
	}
	// 处理应用
	err = process()
	if err != nil {
		return err
	}

	// 下载资源文件
	err = downloadResource(resourcesList, installLocation)
	if err != nil {
		return err
	}
	// 安装配置文件
	err = installConfig(installLocation, cliConfig)
	if err != nil {
		return err
	}
	// 重载应用
	err = reload()
	if err != nil {
		return err
	}
	return nil
}

// 重载
func reload() (err error) {
	// 安装配置文件
	err = installConfig(installLocation, cliConfig)
	if err != nil {
		return err
	}
	// 重载服务
	err = reloadService(serviceName, serviceContent)
	if err != nil {
		return err
	}
	return nil
}

// 测试
func test() (err error) {
	if cliConfig != "" {
		err = testConfig(installLocation+targetFiles[0], cliConfig)
		if err != nil {
			return err
		}
	} else {
		err = testConfig(installLocation+targetFiles[0], installLocation+defaultConfig)
		if err != nil {
			return err
		}
	}
	return nil
}
