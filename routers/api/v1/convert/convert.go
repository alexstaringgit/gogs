// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package convert

import (
	"fmt"

	"github.com/Unknwon/com"

	api "github.com/gogits/go-gogs-client"
	"github.com/gogits/git-module"

	"github.com/gogits/gogs/models"
	"github.com/gogits/gogs/modules/setting"
)

// ToApiBranch converts user to its API format.
func ToApiBranch(b *models.Branch,c *git.Commit) *api.Branch {
	return &api.Branch{
			Name: b.Name,
			Commit: ToApiCommit(c),
		}
}
// ToApiCommit converts user to its API format.
func ToApiCommit(c *git.Commit) *api.PayloadCommit {
	return &api.PayloadCommit{
		ID: c.ID.String(),
		Message: c.Message(),
		URL: "Not implemented",
		Author: &api.PayloadAuthor{
			Name: c.Committer.Name,
			Email: c.Committer.Email,
			/* UserName: c.Committer.UserName, */
		},
	}
}
// ToApiUser converts user to its API format.
func ToApiUser(u *models.User) *api.User {
	return &api.User{
		ID:        u.Id,
		UserName:  u.Name,
		FullName:  u.FullName,
		Email:     u.Email,
		AvatarUrl: u.AvatarLink(),
	}
}

func ToApiEmail(email *models.EmailAddress) *api.Email {
	return &api.Email{
		Email:    email.Email,
		Verified: email.IsActivated,
		Primary:  email.IsPrimary,
	}
}

// ToApiRepository converts repository to API format.
func ToApiRepository(owner *models.User, repo *models.Repository, permission api.Permission) *api.Repository {
	cl := repo.CloneLink()
	return &api.Repository{
		Id:          repo.ID,
		Owner:       *ToApiUser(owner),
		FullName:    owner.Name + "/" + repo.Name,
		Private:     repo.IsPrivate,
		Fork:        repo.IsFork,
		HtmlUrl:     setting.AppUrl + owner.Name + "/" + repo.Name,
		CloneUrl:    cl.HTTPS,
		SshUrl:      cl.SSH,
		Permissions: permission,
	}
}

// ToApiPublicKey converts public key to its API format.
func ToApiPublicKey(apiLink string, key *models.PublicKey) *api.PublicKey {
	return &api.PublicKey{
		ID:      key.ID,
		Key:     key.Content,
		URL:     apiLink + com.ToStr(key.ID),
		Title:   key.Name,
		Created: key.Created,
	}
}

// ToApiHook converts webhook to its API format.
func ToApiHook(repoLink string, w *models.Webhook) *api.Hook {
	config := map[string]string{
		"url":          w.URL,
		"content_type": w.ContentType.Name(),
	}
	if w.HookTaskType == models.SLACK {
		s := w.GetSlackHook()
		config["channel"] = s.Channel
		config["username"] = s.Username
		config["icon_url"] = s.IconURL
		config["color"] = s.Color
	}

	return &api.Hook{
		ID:      w.ID,
		Type:    w.HookTaskType.Name(),
		URL:     fmt.Sprintf("%s/settings/hooks/%d", repoLink, w.ID),
		Active:  w.IsActive,
		Config:  config,
		Events:  w.EventsArray(),
		Updated: w.Updated,
		Created: w.Created,
	}
}

// ToApiDeployKey converts deploy key to its API format.
func ToApiDeployKey(apiLink string, key *models.DeployKey) *api.DeployKey {
	return &api.DeployKey{
		ID:       key.ID,
		Key:      key.Content,
		URL:      apiLink + com.ToStr(key.ID),
		Title:    key.Name,
		Created:  key.Created,
		ReadOnly: true, // All deploy keys are read-only.
	}
}

func ToApiOrganization(org *models.User) *api.Organization {
	return &api.Organization{
		ID:          org.Id,
		AvatarUrl:   org.AvatarLink(),
		UserName:    org.Name,
		FullName:    org.FullName,
		Description: org.Description,
		Website:     org.Website,
		Location:    org.Location,
	}
}
