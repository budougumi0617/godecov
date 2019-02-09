// Copyright (c) 2019. Yoichiro Shimizu @budougumi0617

package godecov

import (
	"path"

	"github.com/pkg/errors"
)

// GetPulls gets pull requests information from /api/gh/:owner/:repo/pulls.
// See also https://docs.codecov.io/reference#section-list-pull-requests
func (cli *Client) GetPulls(owner, repo string) (*PullsResponse, error) {
	// TODO Add parameters
	if len(owner) == 0 {
		return nil, errors.New("owner name was empty")
	}
	if len(repo) == 0 {
		return nil, errors.New("repository name was empty")
	}

	var pr PullsResponse
	p := path.Join(owner, repo, "pulls")
	if err := cli.get(p, nil, &pr); err != nil {
		return nil, err
	}
	return &pr, nil
}
