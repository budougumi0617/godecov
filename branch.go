// Copyright (c) 2019. Yoichiro Shimizu @budougumi0617

package godecov

import (
	"path"
	"time"

	"github.com/pkg/errors"
)

// GetBranch gets pull requests information from /api/gh/:owner/:repo/branch/:branch
// See also https://docs.codecov.io/reference#section-get-a-single-branch
func (cli *Client) GetBranch(owner, repo, branch string) (*BranchResponse, error) {
	// TODO Add parameters
	if len(owner) == 0 {
		return nil, errors.New("owner name was empty")
	}
	if len(repo) == 0 {
		return nil, errors.New("repository name was empty")
	}
	if len(branch) == 0 {
		return nil, errors.New("repository name was empty")
	}

	var br BranchResponse
	p := path.Join(owner, repo, "branch", branch)
	if err := cli.get(p, nil, &br); err != nil {
		return nil, err
	}
	return &br, nil
}

// BranchResponse is struct for response from /api/gh/:owner/:repo/branch/:branch
type BranchResponse struct {
	Commits []struct {
		Pullid       interface{} `json:"pullid"`
		Author       Author      `json:"author"`
		Timestamp    string      `json:"timestamp"`
		ParentTotals Totals2     `json:"parent_totals"`
		State        string      `json:"state"`
		Totals       Totals2     `json:"totals"`
		Commitid     string      `json:"commitid"`
		CiPassed     bool        `json:"ci_passed"`
		Branch       string      `json:"branch"`
		Message      string      `json:"message"`
		Merged       bool        `json:"merged"`
	} `json:"commits"`
	Repo struct {
		UsingIntegration bool      `json:"using_integration"`
		Name             string    `json:"name"`
		Language         string    `json:"language"`
		Deleted          bool      `json:"deleted"`
		Activated        bool      `json:"activated"`
		Private          bool      `json:"private"`
		Updatestamp      time.Time `json:"updatestamp"`
		Branch           string    `json:"branch"`
		Active           bool      `json:"active"`
		ServiceID        string    `json:"service_id"`
		ImageToken       string    `json:"image_token"`
	} `json:"repo"`
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Owner  Owner `json:"owner"`
	Commit struct {
		Commitid     string      `json:"commitid"`
		Parent       string      `json:"parent"`
		Author       Author      `json:"author"`
		Deleted      interface{} `json:"deleted"`
		Timestamp    string      `json:"timestamp"`
		ParentTotals Totals2     `json:"parent_totals"`
		CiPassed     bool        `json:"ci_passed"`
		Totals       Totals2     `json:"totals"`
		Pullid       interface{} `json:"pullid"`
		Notified     bool        `json:"notified"`
		State        string      `json:"state"`
		Updatestamp  string      `json:"updatestamp"`
		Branch       string      `json:"branch"`
		Report       string      `json:"report"`
		Message      string      `json:"message"`
		Merged       bool        `json:"merged"`
	} `json:"commit"`
}
