// Copyright (c) 2019. Yoichiro Shimizu @budougumi0617

package godecov

import (
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// GetPull gets pull requests information from /api/gh/:owner/:repo/pulls.
// See also https://docs.codecov.io/reference#section-get-a-single-pull-request
func (cli *Client) GetPull(owner, repo string, no int) (*PullResponse, error) {
	// TODO Add parameters
	if len(owner) == 0 {
		return nil, errors.New("owner name was empty")
	}
	if len(repo) == 0 {
		return nil, errors.New("repository name was empty")
	}
	if no <= 0 {
		return nil, errors.New("pull request number was invalid")
	}

	var pr PullResponse
	p := path.Join(owner, repo, "pulls", strconv.Itoa(no))
	if err := cli.get(p, nil, &pr); err != nil {
		return nil, err
	}
	return &pr, nil
}

// PullResponse is struct for response from /api/gh/:owner/:repo/pull/:number
type PullResponse struct {
	Pull struct {
		Head       string      `json:"head"`
		Title      string      `json:"title"`
		Author     Author      `json:"author"`
		ComparedTo string      `json:"compared_to"`
		Pullid     int         `json:"pullid"`
		State      string      `json:"state"`
		Base       string      `json:"base"`
		Commentid  interface{} `json:"commentid"`
		Diff       interface{} `json:"diff"`
	} `json:"pull"`
	Head struct {
		Commitid     string      `json:"commitid"`
		Parent       string      `json:"parent"`
		Author       Author      `json:"author"`
		Deleted      interface{} `json:"deleted"`
		Timestamp    string      `json:"timestamp"`
		ParentTotals interface{} `json:"parent_totals"`
		CiPassed     bool        `json:"ci_passed"`
		Totals       Totals      `json:"totals"`
		Pullid       string      `json:"pullid"`
		Notified     bool        `json:"notified"`
		State        string      `json:"state"`
		Updatestamp  interface{} `json:"updatestamp"`
		Branch       string      `json:"branch"`
		Report       string      `json:"report"`
		Message      string      `json:"message"`
		Merged       bool        `json:"merged"`
	} `json:"head"`
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
	Commits []struct {
		Commitid string `json:"commitid"`
	} `json:"commits"`
	Base    string        `json:"base"`
	Meta    Meta          `json:"meta"`
	Owner   Owner         `json:"owner"`
	Diff    interface{}   `json:"diff"`
	Changes []interface{} `json:"changes"`
}
