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
	defaultConfig          = "config.json"
	defaultInstallLocation = "/usr/local/bin/proxy/"
)

// 功能实现函数
// 下载应用到临时目录
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

// 从压缩文件中按照给定的的文件列表解压需要的文件到输出路径
func extract(archiveFile string, fileList []string, outputLocation string) (err error) {
	// 如果输出路径不存在则创建
	_, _, err = gek_file.Exist(outputLocation)
	if err != nil {
		err = gek_file.CreateDir(outputLocation)
		if err != nil {
			return err
		}
	}

	// 循环解压
	for _, file := range fileList {
		err = gek_exec.Run(exec.Command("unzip", "-o", archiveFile, file, "-d", outputLocation))
		if err != nil {
			return err
		}
	}

	return nil
}

// 给文件列表中的文件赋权
func chmod(fileList []string, mode os.FileMode) (err error) {
	for _, f := range fileList {
		err = os.Chmod(defaultInstallLocation+f, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

// 清理应用文件
func resetApp(appLocation string) (err error) {
	// 删除应用文件
	// 同时删除安装文件夹,特别注意如果为 /usr/local/bin/ 这种公共文件夹,仅删除文件
	err = os.RemoveAll(appLocation)
	if err != nil {
		return err
	}
	return nil
}

// 安装配置文件
func installConfig(installLocation string, config string) (err error) {
	// 如果不存在安装文件夹,返回错误
	_, _, err = gek_file.Exist(installLocation)
	if err != nil {
		return err
	}

	// config不为空则移动config文件到默认配置文件
	if config != "" {
		err = gek_exec.Run(exec.Command("cp", "-f", config, installLocation+defaultConfig))
		if err != nil {
			return err
		}
	} else {
		_, _, err = gek_file.Exist(installLocation + defaultConfig)
		if err != nil {
			_, err := gek_file.CreateFile(installLocation+defaultConfig, "{}")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// 测试配置文件
func testConfig(appLocation string, config string) (err error) {
	// 查看app是否存在
	exist, appLocation, _ := gek_exec.Exist(appLocation)
	if !exist {
		return fmt.Errorf("can not find appLocation")
	}

	if config == "" {
		config = defaultInstallLocation + defaultConfig
	}
	err = gek_exec.Run(exec.Command(appLocation, "-test", "-c", config))
	if err != nil {
		return err
	}

	return nil
}

// 下载资源
func downloadResource(resourcesList []string, downloadLocation string) (err error) {
	// 查找安装路径
	exist, _, _ := gek_file.Exist(downloadLocation)
	if !exist {
		return fmt.Errorf("can not find install location %s", downloadLocation)
	}

	// 下载文件到安装路径
	for _, resource := range resourcesList {
		err = gek_downloader.Downloader(resource, downloadLocation, "")
		if err != nil {
			return err
		}
	}

	return nil
}

// 安装服务
func installService(serviceName string, serviceContent string) (err error) {
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
func uninstallService(serviceName string, serviceContent string) (err error) {
	service := gek_service.NewService(serviceName, serviceContent)
	// 卸载服务
	err = service.Uninstall()
	if err != nil {
		return err
	}
	return nil
}

// 重载服务
func reloadService(serviceName string, serviceContent string) (err error) {
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
