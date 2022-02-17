package main

import (
	"gek_downloader"
	"gek_exec"
	"os"
	"os/exec"
)

var (
	// 支持的系统
	supportedOS = []string{"linux", "freebsd"}
)

// 功能实现函数
// 下载应用到临时目录
func download(downloadLink, location string) (err error) {
	// 建立下载文件夹
	// 已经存在就删除重建
	_, err = os.Stat(location)
	if os.IsExist(err) {
		err = os.RemoveAll(location)
		if err != nil {
			return err
		}
	}
	err = os.MkdirAll(location, 755)
	if err != nil {
		return err
	}

	// 下载文件到文件夹
	err = gek_downloader.Downloader(downloadLink, location, "")
	if err != nil {
		return err
	}
	return nil
}

// 从压缩文件中按照给定的的文件列表解压需要的文件到输出路径
func extract(archiveFile string, fileList []string, location string) (err error) {
	// 如果输出路径不存在则创建
	_, err = os.Stat(location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(location, 755)
		if err != nil {
			return err
		}
	}
	// 循环解压
	for _, file := range fileList {
		err = gek_exec.Run(exec.Command("unzip", "-o", archiveFile, file, "-d", location))
		if err != nil {
			return err
		}
	}
	return nil
}

// 给文件列表中的文件赋权
func chmod(fileList []string, mode os.FileMode) (err error) {
	for _, f := range fileList {
		err = os.Chmod(f, mode)
		if err != nil {
			return err
		}
	}
	return nil
}
