// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// gogs(Go Git Service) is a Go clone of Github.
package main

import (
	"os"
	// "os/user"
	"runtime"

	"github.com/codegangsta/cli"

	"github.com/gogits/gogs/modules/base"
)

// +build go1.1

// Test that go1.1 tag above is included in builds. main.go refers to this definition.
const go11tag = true

const APP_VER = "0.0.9.0316.1"

func init() {
	base.AppVer = APP_VER
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// func checkRunUser() bool {
// 	u, err := user.Current()
// 	if err != nil {
// 		// TODO: log
// 		return false
// 	}
// 	return u.Username == base.Cfg.MustValue("", "RUN_USER")
// }

func main() {
	/*if !checkRunUser() {
		println("The command should be run as", base.Cfg.MustValue("", "RUN_USER"))
		return
	}*/

	app := cli.NewApp()
	app.Name = "Gogs"
	app.Usage = "Go Git Service"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		CmdWeb,
		CmdServ,
		CmdUpdate,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
