package main

import (
	"fmt"
	"gek_downloader"
	"gek_exec"
	"gek_github"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var (
	// 应用安装目录
	installLocation = "/usr/local/bin/proxy/"

	// 应用下载临时目录
	tempLocation = "/tmp/proxy_installer/"
	// 应用文件
	appFiles = []string{"v2ray", "config.json", "geoip.dat", "geosite.dat"}
	// 目标文件名
	targetFiles = []string{installLocation + appFiles[0]}
	// 支持的系统
	supportedOS = []string{"linux", "freebsd"}
)

type Application struct {
	// 应用文件
	File string
	// 应用安装文件夹
	Location string
}

// 新建应用
func newApplication(file string, location string) (a Application) {
	return Application{File: file, Location: location}
}

// 新建应用(从Github)s
func newApplicationFromGithub(repo string, appMap map[string]string, location string) (a Application, err error) {
	// 获取应用链接
	downloadLink, err := gek_github.GetDownloadLink(repo, appMap)
	if err != nil {
		return Application{}, err
	}
	return newApplication(downloadLink, location), nil
}

// 安装应用
func (a Application) install(needExtract bool) (err error) {
	// 检查安装文件夹情况
	// 不存在则新建
	_, err = os.Stat(a.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(a.Location, 755)
		if err != nil {
			return err
		}
	}

	if needExtract {
		// 检查临时文件夹情况
		_, err = os.Stat(tempLocation)
		if os.IsNotExist(err) {
			err = os.MkdirAll(tempLocation, 755)
			if err != nil {
				return err
			}
		}
		// 下载需解压文件到临时文件夹
		err = gek_downloader.Downloader(a.File, tempLocation, "")
		if err != nil {
			return err
		}
		// 解压文件到安装文件夹
		err = extract(tempLocation+"*.zip", appFiles, a.Location)
		if err != nil {
			return err
		}

	} else {
		err = gek_downloader.Downloader(a.File, "", a.Location+appFiles[0])
		if err != nil {
			return err
		}
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

// 卸载应用
func (a Application) uninstall() (err error) {
	// 检测应用安装情况
	_, err = os.Stat(filepath.Join(a.Location, appFiles[0]))
	if os.IsNotExist(err) {
		return fmt.Errorf("can't find app location %s", filepath.Join(a.Location, appFiles[0]))
	}
	// 删除应用文件
	err = os.RemoveAll(filepath.Join(a.Location, appFiles[0]))
	if err != nil {
		return err
	}
	return nil
}

// 测试应用
func (a Application) test() (err error) {
	// 查看app是否存在
	exist, app, _ := gek_exec.Exist(filepath.Join(a.Location, appFiles[0]))
	if !exist {
		return fmt.Errorf("can not find app")
	}

	// 分系统运行不同的命令
	switch runtime.GOOS {
	case supportedOS[0]:
		err = gek_exec.Run(exec.Command(app, "-test"))
	case supportedOS[1]:
		err = gek_exec.Run(exec.Command(app, "test"))
	}
	return err
}
