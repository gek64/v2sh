package main

import (
	"fmt"
	"gek_downloader"
	"gek_exec"
	"gek_github"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	// TEMP 应用下载临时目录
	TEMP = "/tmp/proxy_installer/"
)

type Application struct {
	// app文件
	appFiles []string
	// 应用URL
	Url string
	// 是否需要解压
	needExtract bool
	// 应用安装文件夹
	Location string
}

// 新建应用
func newApplication(appFiles []string, url string, needExtract bool, location string) (a Application) {
	return Application{appFiles: appFiles, Url: url, needExtract: needExtract, Location: location}
}

// 新建应用(从Github)
func newApplicationFromGithub(appFiles []string, repo string, appMap map[string]string, needExtract bool, location string) (a Application, err error) {
	// 获取应用链接
	downloadLink, err := gek_github.GetDownloadLink(repo, appMap)
	if err != nil {
		return Application{}, err
	}
	return newApplication(appFiles, downloadLink, needExtract, location), nil
}

// 安装应用
func (a Application) install() (err error) {
	// 检查安装文件夹情况
	// 不存在则新建
	_, err = os.Stat(a.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(a.Location, 755)
		if err != nil {
			return err
		}
	}

	if a.needExtract {
		// 检查临时文件夹情况
		_, err = os.Stat(TEMP)
		if os.IsNotExist(err) {
			err = os.MkdirAll(TEMP, 755)
			if err != nil {
				return err
			}
		}
		// 结束后删除临时文件夹
		defer func(path string) {
			err = os.RemoveAll(path)
			if err != nil {
				log.Panicln(err)
			}
		}(TEMP)

		// 下载压缩文件到临时文件夹
		err = gek_downloader.Downloader(a.Url, TEMP, "")
		if err != nil {
			return err
		}

		// 读取下载的压缩文件名
		var zipFile string
		fileInfos, err := ioutil.ReadDir(TEMP)
		if err != nil {
			fmt.Println(err)
		}
		for _, f := range fileInfos {
			if strings.Contains(f.Name(), ".zip") {
				zipFile = f.Name()
				break
			}
		}
		if zipFile == "" {
			return fmt.Errorf("can't find the download application archive file")
		}
		// 解压文件到安装文件夹
		err = extract(filepath.Join(TEMP, zipFile), a.appFiles, a.Location)
		if err != nil {
			return err
		}
	} else {
		err = gek_downloader.Downloader(a.Url, "", filepath.Join(a.Location, a.appFiles[0]))
		if err != nil {
			return err
		}
	}

	// 可执行文件赋权755
	for _, appFile := range a.appFiles {
		err = chmodRecursive(filepath.Join(a.Location, appFile), 755)
		if err != nil {
			return err
		}
	}

	return nil
}

// 卸载应用
func (a Application) uninstall() (err error) {
	// 检测应用安装情况
	_, err = os.Stat(filepath.Join(a.Location, a.appFiles[0]))
	if os.IsNotExist(err) {
		return fmt.Errorf("can't find app location %s", filepath.Join(a.Location, a.appFiles[0]))
	}

	// 删除应用文件
	for _, app := range a.appFiles {
		err = os.RemoveAll(filepath.Join(a.Location, app))
		if err != nil {
			return err
		}
	}

	return nil
}

// 测试应用
func (a Application) test(configFile string) (err error) {
	// 查看app是否存在
	exist, app, _ := gek_exec.Exist(filepath.Join(a.Location, a.appFiles[0]))
	if !exist {
		return fmt.Errorf("can not find app")
	}

	// 分系统运行不同的命令
	var c string
	switch runtime.GOOS {
	case supportedOS[0]:
		c = linuxConfigLocation
	case supportedOS[1]:
		c = freebsdConfigLocation
	}

	if configFile != "" {
		err = gek_exec.Run(exec.Command(app, "-test", "-config", configFile))
	} else {
		err = gek_exec.Run(exec.Command(app, "-test", "-confdir", c))
	}

	return err
}
