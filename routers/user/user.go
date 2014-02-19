// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"fmt"
	"net/http"

	"github.com/martini-contrib/render"

	"github.com/gogits/validation"

	"github.com/gogits/gogs/models"
	"github.com/gogits/gogs/utils/log"
)

func SignIn(r render.Render) {
	r.Redirect("/user/signup", 302)
}

func SignUp(req *http.Request, r render.Render) {
	if req.Method == "GET" {
		r.HTML(200, "user/signup", map[string]interface{}{
			"Title": "Sign Up",
		})
		return
	}

	u := &models.User{
		Name:   req.FormValue("username"),
		Email:  req.FormValue("email"),
		Passwd: req.FormValue("passwd"),
	}
	valid := validation.Validation{}
	ok, err := valid.Valid(u)
	if err != nil {
		log.Error("user.SignUp -> valid user: %v", err)
		return
	}
	if !ok {
		for _, err := range valid.Errors {
			log.Warn("user.SignUp -> valid user: %v", err)
		}
		return
	}

	err = models.RegisterUser(u)
	r.HTML(403, "status/403", map[string]interface{}{
		"Title": fmt.Sprintf("%v", err),
	})
}

func Delete(r render.Render) {
	u := &models.User{}
	err := models.DeleteUser(u)
	r.HTML(403, "status/403", map[string]interface{}{
		"Title": fmt.Sprintf("%v", err),
	})
}
