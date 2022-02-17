package main

import (
	"fmt"
	"gek_file"
	"runtime"
)

// 主要功能函数
// 安装
func install() (err error) {

	newApplicationFromGithub()

	// 安装配置文件
	err = installConfig(installLocation, cliConfig)
	if err != nil {
		return err
	}

	// 安装服务
	switch runtime.GOOS {
	case supportedOS[0]:
		err = installService(linuxServiceName, linuxServiceContent)
		if err != nil {
			return err
		}
	case supportedOS[1]:
		err = installService(freebsdServiceName, freebsdServiceContent)
		if err != nil {
			return err
		}
	}

	return nil
}

// 卸载
func uninstall() (err error) {
	// 重置服务
	err = deleteApp(installLocation)
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
	exist, _, _ := gek_file.Exist(targetFiles[0])
	if !exist {
		return fmt.Errorf("app is not installed")
	}

	// 如果已经指定本地文件则不进行下载
	if cliLocalFile == "" {
		// 下载应用
		err = downloadApp(proxyRepo, proxyList, TEMP)
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
		err = testConfig(targetFiles[0], cliConfig)
		if err != nil {
			return err
		}
	} else {
		err = testConfig(targetFiles[0], defaultConfig)
		if err != nil {
			return err
		}
	}
	return nil
}
