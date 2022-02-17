package main

import (
	"fmt"
	"gek_downloader"
	"os"
)

type Resources struct {
	// 资源文件
	files []string
	// 资源安装文件夹
	Location string
}

// 新建资源
func newResources(files []string, location string) (r Resources) {
	return Resources{files: files, Location: location}
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
	for _, f := range r.files {
		err = gek_downloader.Downloader(f, r.Location, "")
		if err != nil {
			return err
		}
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
