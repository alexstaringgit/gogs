// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gogits/gogs/models"
	"github.com/gogits/gogs/modules/middleware"
)

func Dashboard(ctx *middleware.Context) {
	ctx.Data["Title"] = "Admin Dashboard"
	ctx.Data["Stats"] = models.GetStatistic()
	ctx.HTML(200, "admin/dashboard")
}

func Users(ctx *middleware.Context) {
	ctx.Data["Title"] = "User Management"

	var err error
	ctx.Data["Users"], err = models.GetUsers(100, 0)
	if err != nil {
		ctx.Handle(200, "admin.Users", err)
		return
	}
	ctx.HTML(200, "admin/users")
}

func Repositories(ctx *middleware.Context) {
	ctx.Data["Title"] = "Repository Management"
	var err error
	ctx.Data["Repos"], err = models.GetRepos(100, 0)
	if err != nil {
		ctx.Handle(200, "admin.Repositories", err)
		return
	}
	ctx.HTML(200, "admin/repos")
}
