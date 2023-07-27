package main

import (
    "fmt"
    "github.com/gek64/gek/gApp"
    "github.com/gek64/gek/gDownloader"
    "github.com/gek64/gek/gExec"
    "io/ioutil"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

type Res struct {
    gApp.Resources
}

// 解压安装本地资源压缩文件
func (r Res) installFromLocalArchiveFile(localFile string) (err error) {
    // 检查本地文件是否存在
    _, err = os.Stat(localFile)
    if os.IsNotExist(err) {
        return fmt.Errorf("%s is not exist", localFile)
    }

    // 检测资源安装路径是否存在
    // 不存在则创建
    _, err = os.Stat(r.Location)
    if os.IsNotExist(err) {
        err = os.MkdirAll(r.Location, 0755)
        if err != nil {
            return err
        }
    }

    // 解压资源文件到资源安装路径
    for _, file := range r.Files {
        err = gExec.Run(exec.Command("unzip", "-o", "-d", r.Location, localFile, file))
        if err != nil {
            return err
        }
    }

    // 赋权0644
    for _, file := range r.Files {
        err = os.Chmod(filepath.Join(r.Location, file), 0644)
        if err != nil {
            return err
        }
    }

    return nil
}

// 下载资源压缩文件并解压安装
func (r Res) installFromInternetArchiveFile(url string, tempLocation string) (err error) {
    // 检查临时文件夹是否存在
    _, err = os.Stat(tempLocation)
    if os.IsNotExist(err) {
        appTemp := gApp.NewTemp(tempLocation)
        err = appTemp.Create()
        if err != nil {
            return err
        }

        defer func(appTemp gApp.Temp) {
            err := appTemp.Delete()
            if err != nil {
                log.Panicln(err)
            }
        }(appTemp)
    }

    // 下载资源压缩文件
    err = gDownloader.Downloader(url, tempLocation, "")
    if err != nil {
        return err
    }

    // 解压下载的资源压缩文件,到资源安装路径
    fileInfo, err := ioutil.ReadDir(tempLocation)
    if err != nil {
        return err
    }
    for _, f := range fileInfo {
        if strings.Contains(f.Name(), ".zip") {
            return r.installFromLocalArchiveFile(filepath.Join(tempLocation, f.Name()))
        }
    }
    return fmt.Errorf("can't find archive file")
}
