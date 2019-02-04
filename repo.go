// Copyright (c) 2019.  Yoichiro Shimizu @budougumi0617

package godecov

import (
	"path"

	"github.com/pkg/errors"
)

// GetSingleRepository gets a single repository information from /api/gh/:owner/:repo.
// See also https://docs.codecov.io/v4.3.0/reference#section-get-a-single-repository
func (cli *Client) GetSingleRepository(owner, repo string) (*SingleRepoResponse, error) {
	if len(owner) == 0 {
		return nil, errors.New("owner name was empty")
	}
	if len(repo) == 0 {
		return nil, errors.New("repository name was empty")
	}

	var sr SingleRepoResponse
	p := path.Join(owner, repo)
	if err := cli.get(p, nil, &sr); err != nil {
		return nil, err
	}
	return &sr, nil
}
