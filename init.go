package main

import (
    "fmt"
    "github.com/gek64/gek/gApp"
)

var (
    app         App
    config      gApp.Config
    resources   Res
    service     gApp.Service
    tempFolder  = "/tmp/proxy"
    needExtract = true
    cc          CC
)

func initApp(local bool) (err error) {
    var a = &app
    var c = &config
    var r = &resources
    var s = &service

    err = initConf()
    if err != nil {
        return err
    }

    // 应用初始化
    if local {
        a.Application = gApp.NewApplication(cc.Application.File, "", cc.Application.Location, cc.Application.UninstallDeleteLocation)
    } else {
        a.Application, err = gApp.NewApplicationFromGithub(cc.Application.File, cc.Application.Repo, cc.Application.RepoList, cc.Application.Location, cc.Application.UninstallDeleteLocation, tempFolder)
        if err != nil {
            return err
        }
    }

    // 配置初始化
    *c = gApp.NewConfig(cc.Config.Name, cc.Config.Content, cc.Config.Location, cc.Config.UninstallDeleteLocation)

    // 资源初始化
    r.Resources = gApp.NewResources(cc.Resources.File, cc.Resources.URL, cc.Resources.Location, cc.Application.UninstallDeleteLocation)

    // 检查系统中的init system
    _, initSystemBin := gApp.CheckInitSystem()

    // 服务初始化
    switch initSystemBin {
    case gApp.InitSystem["systemd"]:
        bytes, err := container.ReadFile("configs/v2ray.service")
        if err != nil {
            return err
        }
        *s = gApp.NewService(cc.Service.Name, string(bytes))
    case gApp.InitSystem["openrc"]:
        return fmt.Errorf("openrc does not yet support")
    case gApp.InitSystem["rc.d"]:
        bytes, err := container.ReadFile("configs/v2ray")
        if err != nil {
            return err
        }
        *s = gApp.NewService(cc.Service.Name, string(bytes))
    default:
        var supportInitSystemListString string
        for key := range gApp.InitSystem {
            supportInitSystemListString = supportInitSystemListString + ", " + key
        }
        return fmt.Errorf("no supported init system found, currently only %s are supported", supportInitSystemListString)
    }

    return nil
}
