package main

import (
	"github.com/gek64/gek/gApp"
	"github.com/gek64/gek/gApp/compression/unzip"
	"github.com/gek64/gek/gDownloader"
	"os"
	"path/filepath"
	"runtime"
)

// 可执行文件操作
func installBinaryFile(localArchiveFile string) (err error) {
	if localArchiveFile != "" {
		// 使用本地文件
		bytes, err := os.ReadFile(localArchiveFile)
		if err != nil {
			return err
		}
		err = os.WriteFile(filepath.Join(os.TempDir(), "v2ray.zip"), bytes, 0644)
		if err != nil {
			return err
		}
	} else {
		// 从网络下载
		downloadURL, err := GetDownloadURL()
		if err != nil {
			return err
		}
		err = gDownloader.Download(downloadURL, filepath.Join(os.TempDir(), "v2ray.zip"), "")
		if err != nil {
			return err
		}
	}

	if runtime.GOOS != "windows" {
		// 解压可执行文件
		err = unzip.Decompress(filepath.Join(os.TempDir(), "v2ray.zip"), "/usr/local/bin", "v2ray")
		if err != nil {
			return err
		}
		// 可执行文件赋权
		err = os.Chmod("/usr/local/bin/v2ray", 0755)
		if err != nil {
			return err
		}
	} else {
		// windows下解压文件名需要.exe后缀
		err = unzip.Decompress(filepath.Join(os.TempDir(), "v2ray.zip"), "/usr/local/bin", "v2ray.exe")
		if err != nil {
			return err
		}
	}
	// 解压资源文件
	return unzip.Decompress(filepath.Join(os.TempDir(), "v2ray.zip"), "/usr/local/etc/v2ray", "geoip.dat", "geosite.dat")
}
func uninstallBinaryFile() (err error) {
	err = os.RemoveAll("/usr/local/etc/v2ray")
	if err != nil {
		return err
	}
	return os.RemoveAll("/usr/local/bin/v2ray")
}
func updateBinaryFile(localArchiveFile string) (err error) {
	// 服务初始化
	s, err := initService()
	if err != nil {
		return err
	}
	// 服务停止,关闭自启
	err = s.Unload()
	if err != nil {
		return err
	}
	// 执行新的二进制文件安装
	err = installBinaryFile(localArchiveFile)
	if err != nil {
		return err
	}
	// 服务启动,开启自启
	return s.Load()
}

// 配置文件操作
func installConfig(file string) (err error) {
	if file == "" {
		return os.WriteFile("/usr/local/etc/v2ray/config.json", []byte("{}"), 0644)
	}
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return os.WriteFile("/usr/local/etc/v2ray/config.json", bytes, 0644)
}

// 服务操作
func initService() (service *gApp.Service, err error) {
	var serviceName string
	// 获取初始化系统名称,及服务内容
	initSystem, serviceContent, err := GetService()
	if err != nil {
		return nil, err
	}
	// 获取服务名称
	switch initSystem {
	case "systemd":
		serviceName = "v2ray.service"
	default:
		serviceName = "v2ray"
	}
	// 初始化服务
	return gApp.NewService(initSystem, serviceName, serviceContent)
}
func installService() (err error) {
	// 服务初始化
	s, err := initService()
	if err != nil {
		return err
	}
	// 服务安装
	err = s.Install()
	if err != nil {
		return err
	}
	// 服务启动,开启自启
	return s.Load()
}
func uninstallService() (err error) {
	// 服务初始化
	s, err := initService()
	if err != nil {
		return err
	}
	// 服务停止,关闭自启
	err = s.Unload()
	if err != nil {
		return err
	}
	// 服务卸载
	return s.Uninstall()
}
func updateService() (err error) {
	// 服务初始化
	s, err := initService()
	if err != nil {
		return err
	}
	// 服务停止,关闭自启
	err = s.Unload()
	if err != nil {
		return err
	}
	// 新的服务安装
	err = s.Install()
	if err != nil {
		return err
	}
	// 服务启动,开启自启
	return s.Load()
}
func reloadService() (err error) {
	// 服务初始化
	s, err := initService()
	if err != nil {
		return err
	}
	// 服务重载
	return s.Reload()
}
