package main

import (
	"fmt"
	"gek_downloader"
	"gek_exec"
	"gek_file"
	"gek_github"
	"gek_service"
	"os"
	"os/exec"
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
	targetFiles = []string{"v2ray", "v2ctl", "geoip.dat", "geosite.dat"}
	// 配置文件
	config = "config.json"
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
	resourcesRepo = "Loyalsoldier/v2ray-rules-dat"
	resourcesList = map[string]string{
		"geoip":   "geoip.dat",
		"geosite": "geosite.dat",
	}
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
func install(repo string, repoList map[string]string, needExtract bool) (err error) {
	// 需要文件未安装
	exist, _, _ := gek_file.Exist(installLocation + targetFileName)
	if exist {
		return fmt.Errorf("app is already installed")
	}
	// 下载应用
	err = downloadApp(repo, repoList)
	if err != nil {
		return err
	}
	// 处理应用
	err = processApp(needExtract)
	if err != nil {
		return err
	}
	// 安装服务
	err = installService()
	if err != nil {
		return err
	}
	return nil
}

// 卸载
func uninstall() (err error) {
	// 重置服务
	err = resetApp()
	if err != nil {
		return
	}
	// 卸载服务
	err = uninstallService()
	if err != nil {
		return err
	}
	return nil
}

// 更新
func update(repo string, repoList map[string]string, needExtract bool) (err error) {
	// 需要应用已安装
	exist, _, _ := gek_file.Exist(installLocation + targetFileName)
	if !exist {
		return fmt.Errorf("app is not installed")
	}
	// 下载应用
	err = downloadApp(repo, repoList)
	if err != nil {
		return err
	}
	// 处理应用
	err = processApp(needExtract)
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
	// 重载服务
	err = reloadService()
	if err != nil {
		return err
	}
	return nil
}

// 功能实现函数
// 下载应用
func downloadApp(repo string, repoList map[string]string, tempLocation string) (err error) {
	// 建立临时文件夹
	// 已经存在就删除重建
	exist, _, _ := gek_file.Exist(tempLocation)
	if exist {
		err = gek_exec.Run("rm -rf " + tempLocation)
		if err != nil {
			return err
		}
	}
	err = gek_file.CreateDir(tempLocation)
	if err != nil {
		return err
	}

	// 获取下载链接
	downloadLink, err := gek_github.GetDownloadLink(repo, repoList)
	if err != nil {
		return err
	}

	// 下载文件到临时文件夹
	err = gek_downloader.Downloader(downloadLink, tempLocation, "")
	if err != nil {
		return err
	}
	return nil
}

// 处理应用下载文件
func processApp(installLocation string, tempLocation string, needExtract bool, localFile string) (err error) {
	// 如果不存在安装文件夹,则创建
	exist, _, _ := gek_file.Exist(installLocation)
	if !exist {
		err := gek_file.CreateDir(installLocation)
		if err != nil {
			return err
		}
	}

	// 指定需要处理的应用本地文件
	var appFile string = ""
	if localFile != "" {
		appFile = localFile
	} else {
		appFile = tempLocation + "*.zip"
	}

	// 需要解压执行解压,不需要解压则执行改名移动
	if needExtract {
		for i, exfile := range extractFiles {
			err = gek_exec.Run(exec.Command("unzip", "-o", appFile, exfile, "-d", installLocation))
			if err != nil {
				return err
			}
			// 如果解压出的文件名与目标文件名不相同则更改
			if installLocation+exfile != installLocation+targetFiles[i] {
				err = gek_exec.Run(exec.Command("mv", installLocation+exfile, installLocation+targetFiles[i]))
				if err != nil {
					return err
				}
			}
		}
	} else {
		for i, exfile := range extractFiles {
			err = gek_exec.Run(exec.Command("mv", tempLocation+exfile, installLocation+targetFiles[i]))
			if err != nil {
				return err
			}
		}
	}

	// 赋权755
	for _, tfile := range targetFiles {
		err = os.Chmod(installLocation+tfile, 755)
		if err != nil {
			return err
		}
	}

	// 删除临时文件夹
	err = os.RemoveAll(tempLocation)
	if err != nil {
		return err
	}
	return nil
}

// 清理应用文件
func resetApp() (err error) {
	// 删除应用文件
	// 同时删除安装文件夹,特别注意如果为 /usr/local/bin/ 这种公共文件夹,仅删除文件
	err = os.RemoveAll(installLocation)
	if err != nil {
		return err
	}
	return nil
}

// 安装配置文件
func installConfig(config string, installLocation string) (err error) {
	// 如果不存在安装文件夹,则创建
	exist, _, _ := gek_file.Exist(installLocation)
	if !exist {
		err = gek_file.CreateDir(installLocation)
		if err != nil {
			return err
		}
	}

	// 配置文件移动到安装文件夹
	err = gek_exec.Run(exec.Command("mv", config, installLocation))
	if err != nil {
		return err
	}

	return nil
}

// 测试配置文件
func testConfig(config string, app string) (err error) {
	// 查看app是否存在
	exist, app, _ := gek_exec.Exist(app)
	if !exist {
		return fmt.Errorf("can not find %s", app)
	}
	// 查看config是否存在
	exist, _, _ = gek_file.Exist(config)
	if !exist {
		return fmt.Errorf("can not find %s", config)
	}

	// 测试
	err = gek_exec.Run(exec.Command(app, "-test", "-c", config))
	if err != nil {
		return err
	}

	return nil
}

// 下载资源
func downloadResource(repo string, repoList map[string]string, installLocation string) (err error) {
	// 查找安装路径
	exist, _, _ := gek_file.Exist(installLocation)
	if !exist {
		return fmt.Errorf("can not find install location %s", installLocation)
	}

	// 获取下载链接
	downloadLink, err := gek_github.GetDownloadLink(repo, repoList)
	if err != nil {
		return err
	}

	// 下载文件到安装路径
	err = gek_downloader.Downloader(downloadLink, installLocation, "")
	if err != nil {
		return err
	}

	return nil
}

// 安装服务
func installService() (err error) {
	service := gek_service.NewService(serviceName, serviceContent)
	// 安装服务
	err = service.Install()
	if err != nil {
		return err
	}
	// 查看服务状态
	err = service.Status()
	if err != nil {
		return err
	}
	return nil
}

// 卸载服务
func uninstallService() (err error) {
	service := gek_service.NewService(serviceName, serviceContent)
	// 卸载服务
	err = service.Uninstall()
	if err != nil {
		return err
	}
	return nil
}

// 重载服务
func reloadService() (err error) {
	service := gek_service.NewService(serviceName, serviceContent)
	// 重载服务
	err = service.Reload()
	if err != nil {
		return err
	}
	// 查看服务状态
	err = service.Status()
	if err != nil {
		return err
	}
	return nil
}
