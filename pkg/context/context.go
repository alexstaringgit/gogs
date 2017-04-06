// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package context

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/i18n"
	"github.com/go-macaron/session"
	log "gopkg.in/clog.v1"
	"gopkg.in/macaron.v1"

	"github.com/gogits/gogs/models"
	"github.com/gogits/gogs/pkg/auth"
	"github.com/gogits/gogs/pkg/form"
	"github.com/gogits/gogs/pkg/setting"
)

// Context represents context of a request.
type Context struct {
	*macaron.Context
	Cache   cache.Cache
	csrf    csrf.CSRF
	Flash   *session.Flash
	Session session.Store

	User        *models.User
	IsSigned    bool
	IsBasicAuth bool

	Repo *Repository
	Org  *Organization
}

func (ctx *Context) UserID() int64 {
	if !ctx.IsSigned {
		return 0
	}
	return ctx.User.ID
}

// HasError returns true if error occurs in form validation.
func (ctx *Context) HasApiError() bool {
	hasErr, ok := ctx.Data["HasError"]
	if !ok {
		return false
	}
	return hasErr.(bool)
}

func (ctx *Context) GetErrMsg() string {
	return ctx.Data["ErrorMsg"].(string)
}

// HasError returns true if error occurs in form validation.
func (ctx *Context) HasError() bool {
	hasErr, ok := ctx.Data["HasError"]
	if !ok {
		return false
	}
	ctx.Flash.ErrorMsg = ctx.Data["ErrorMsg"].(string)
	ctx.Data["Flash"] = ctx.Flash
	return hasErr.(bool)
}

// HasValue returns true if value of given name exists.
func (ctx *Context) HasValue(name string) bool {
	_, ok := ctx.Data[name]
	return ok
}

// HTML responses template with given status.
func (ctx *Context) HTML(status int, name string) {
	log.Trace("Template: %s", name)
	ctx.Context.HTML(status, name)
}

// Success responses template with status http.StatusOK.
func (c *Context) Success(name string) {
	c.HTML(http.StatusOK, name)
}

// JSONSuccess responses JSON with status http.StatusOK.
func (c *Context) JSONSuccess(data interface{}) {
	c.JSON(http.StatusOK, data)
}

// RenderWithErr used for page has form validation but need to prompt error to users.
func (ctx *Context) RenderWithErr(msg, tpl string, f interface{}) {
	if f != nil {
		form.Assign(f, ctx.Data)
	}
	ctx.Flash.ErrorMsg = msg
	ctx.Data["Flash"] = ctx.Flash
	ctx.HTML(http.StatusOK, tpl)
}

// Handle handles and logs error by given status.
func (ctx *Context) Handle(status int, title string, err error) {
	switch status {
	case http.StatusNotFound:
		ctx.Data["Title"] = "Page Not Found"
	case http.StatusInternalServerError:
		ctx.Data["Title"] = "Internal Server Error"
		log.Error(2, "%s: %v", title, err)
		if !setting.ProdMode || (ctx.IsSigned && ctx.User.IsAdmin) {
			ctx.Data["ErrorMsg"] = err
		}
	}
	ctx.HTML(status, fmt.Sprintf("status/%d", status))
}

// NotFound renders the 404 page.
func (ctx *Context) NotFound() {
	ctx.Handle(http.StatusNotFound, "", nil)
}

// ServerError renders the 500 page.
func (c *Context) ServerError(title string, err error) {
	c.Handle(http.StatusInternalServerError, title, err)
}

// NotFoundOrServerError use error check function to determine if the error
// is about not found. It responses with 404 status code for not found error,
// or error context description for logging purpose of 500 server error.
func (c *Context) NotFoundOrServerError(title string, errck func(error) bool, err error) {
	if errck(err) {
		c.NotFound()
		return
	}
	c.ServerError(title, err)
}

func (ctx *Context) HandleText(status int, title string) {
	ctx.PlainText(status, []byte(title))
}

func (ctx *Context) ServeContent(name string, r io.ReadSeeker, params ...interface{}) {
	modtime := time.Now()
	for _, p := range params {
		switch v := p.(type) {
		case time.Time:
			modtime = v
		}
	}
	ctx.Resp.Header().Set("Content-Description", "File Transfer")
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+name)
	ctx.Resp.Header().Set("Content-Transfer-Encoding", "binary")
	ctx.Resp.Header().Set("Expires", "0")
	ctx.Resp.Header().Set("Cache-Control", "must-revalidate")
	ctx.Resp.Header().Set("Pragma", "public")
	http.ServeContent(ctx.Resp, ctx.Req.Request, name, modtime, r)
}

// Contexter initializes a classic context for a request.
func Contexter() macaron.Handler {
	return func(c *macaron.Context, l i18n.Locale, cache cache.Cache, sess session.Store, f *session.Flash, x csrf.CSRF) {
		ctx := &Context{
			Context: c,
			Cache:   cache,
			csrf:    x,
			Flash:   f,
			Session: sess,
			Repo: &Repository{
				PullRequest: &PullRequest{},
			},
			Org: &Organization{},
		}

		if len(setting.HTTP.AccessControlAllowOrigin) > 0 {
			ctx.Header().Set("Access-Control-Allow-Origin", setting.HTTP.AccessControlAllowOrigin)
			ctx.Header().Set("'Access-Control-Allow-Credentials' ", "true")
			ctx.Header().Set("Access-Control-Max-Age", "3600")
			ctx.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		}

		// Compute current URL for real-time change language.
		ctx.Data["Link"] = setting.AppSubURL + strings.TrimSuffix(ctx.Req.URL.Path, "/")

		ctx.Data["PageStartTime"] = time.Now()

		// Get user from session if logined.
		ctx.User, ctx.IsBasicAuth = auth.SignedInUser(ctx.Context, ctx.Session)

		if ctx.User != nil {
			ctx.IsSigned = true
			ctx.Data["IsSigned"] = ctx.IsSigned
			ctx.Data["SignedUser"] = ctx.User
			ctx.Data["SignedUserID"] = ctx.User.ID
			ctx.Data["SignedUserName"] = ctx.User.Name
			ctx.Data["IsAdmin"] = ctx.User.IsAdmin
		} else {
			ctx.Data["SignedUserID"] = 0
			ctx.Data["SignedUserName"] = ""
		}

		// If request sends files, parse them here otherwise the Query() can't be parsed and the CsrfToken will be invalid.
		if ctx.Req.Method == "POST" && strings.Contains(ctx.Req.Header.Get("Content-Type"), "multipart/form-data") {
			if err := ctx.Req.ParseMultipartForm(setting.AttachmentMaxSize << 20); err != nil && !strings.Contains(err.Error(), "EOF") { // 32MB max size
				ctx.Handle(500, "ParseMultipartForm", err)
				return
			}
		}

		ctx.Data["CSRFToken"] = x.GetToken()
		ctx.Data["CSRFTokenHTML"] = template.HTML(`<input type="hidden" name="_csrf" value="` + x.GetToken() + `">`)
		log.Trace("Session ID: %s", sess.ID())
		log.Trace("CSRF Token: %v", ctx.Data["CSRFToken"])

		ctx.Data["ShowRegistrationButton"] = setting.Service.ShowRegistrationButton
		ctx.Data["ShowFooterBranding"] = setting.ShowFooterBranding
		ctx.Data["ShowFooterVersion"] = setting.ShowFooterVersion

		c.Map(ctx)
	}
}
