package main

import (
	"fmt"
	"gek_downloader"
	"os"
)

type Resources struct {
	// 资源文件
	Files []string
	// 资源链接
	Urls []string
	// 资源安装文件夹
	Location string
}

// 新建资源
func newResources(files []string, urls []string, location string) (r Resources) {
	return Resources{Files: files, Urls: urls, Location: location}
}

// 安装资源
func (r Resources) install() (err error) {
	// 检测资源安装路径是否存在
	// 不存在则创建
	_, err = os.Stat(r.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(r.Location, 755)
		if err != nil {
			return err
		}
	}
	// 下载资源文件到资源安装路径
	for _, url := range r.Urls {
		err = gek_downloader.Downloader(url, r.Location, "")
		if err != nil {
			return err
		}
	}

	// 赋权755
	err = chmodRecursive(r.Location, 755)
	if err != nil {
		return err
	}

	return nil
}

// 安装资源(从本地文件解压,无需网络下载)
func (r Resources) installFromLocal(localFile string) (err error) {
	// 检查本地文件是否存在
	_, err = os.Stat(localFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s is not exist", localFile)
	}

	// 检测资源安装路径是否存在
	// 不存在则创建
	_, err = os.Stat(r.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(r.Location, 755)
		if err != nil {
			return err
		}
	}

	// 解压资源文件到资源安装路径
	err = extract(localFile, r.Files, r.Location)
	if err != nil {
		return err
	}

	// 赋权755
	err = chmodRecursive(r.Location, 755)
	if err != nil {
		return err
	}

	return nil
}

// 卸载资源,并删除资源安装文件夹
func (r Resources) uninstall() (err error) {
	// 检测资源安装路径是否存在
	_, err = os.Stat(r.Location)
	if os.IsNotExist(err) {
		return fmt.Errorf("can't find resources location %s", r.Location)
	}
	// 删除资源安装文件夹,及所有资源文件
	err = os.RemoveAll(r.Location)
	if err != nil {
		return err
	}
	return nil
}
