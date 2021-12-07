package main

import (
	"flag"
	"fmt"
	"gek_toolbox"
	"log"
	"os"
)

var (
	cliInstall   bool
	cliUninstall bool
	cliUpdate    bool
	cliReload    bool
	cliLocal     string
	cliConfig    string
	cliHelp      bool
	cliVersion   bool
)

func init() {
	flag.BoolVar(&cliInstall, "I", false, "install")
	flag.BoolVar(&cliUninstall, "R", false, "uninstall")
	flag.BoolVar(&cliUpdate, "U", false, "update")
	flag.BoolVar(&cliReload, "r", false, "reload")
	flag.StringVar(&cliLocal, "l", "", "use local file without download from network")
	flag.StringVar(&cliConfig, "c", "", "use local config")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
    proxyctl [Command] [Arguments]

Arguments:
	-l <bin.zip>      : use local file without download from network
	-c <config.json>  : use local config

Command:
	-I                : install
	-R                : uninstall
	-U                : update
	-r                : reload
	-t                : test config
	-h                : show help
	-v                : show version

Example:
    1) proxyctl -I -l bins.zip -c config.json   : Install proxy and service
    2) proxyctl -U                              : Update proxy and resources
    3) proxyctl -R                              : Uninstall proxy and service
	4) proxyctl -t -c config.json               : Test config
    5) proxyctl -r                              : Reload service`
		fmt.Println(helpInfo)
	}

	// 如果无 args 或者 指定 h 参数,打印用法后退出
	if len(os.Args) == 1 || cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println("v1.00")
		os.Exit(0)
	}

	// 检查运行库是否完整
	err := gek_toolbox.CheckToolbox(toolbox)
	if err != nil {
		log.Fatal(err)
	}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release`
	fmt.Println(versionInfo)
}

func main() {
	if cliInstall {
		if cliEE {
			err := install(qbteeRepo, qbteeList, true)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := install(qbtRepo, qbtList, false)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if cliUpdate {
		if cliEE {
			err := update(qbteeRepo, qbteeList, true)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := update(qbtRepo, qbtList, false)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if cliUninstall {
		err := uninstall()
		if err != nil {
			log.Println(err)
		}
	}
	if cliReload {
		err := reload()
		if err != nil {
			log.Println(err)
		}
	}
}
