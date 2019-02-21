package godecov

import (
	"path"

	"github.com/pkg/errors"
)

// GetBranches gets branches information from /api/gh/:owner/:repo/branches
// See also https://docs.codecov.io/reference#branches
func (cli *Client) GetBranches(owner, repo string) (*BranchesResponse, error) {
	if len(owner) == 0 {
		return nil, errors.New("owner name was empty")
	}
	if len(repo) == 0 {
		return nil, errors.New("repository name was empty")
	}

	var br BranchesResponse
	p := path.Join(owner, repo, "branches")
	if err := cli.get(p, nil, &br); err != nil {
		return nil, err
	}
	return &br, nil
}

// BranchesResponse is struct for response from /api/gh/:owner/:repo/branches
type BranchesResponse struct {
	Owner    Owner    `json:"owner"`
	Repo     Repo     `json:"repo"`
	Meta     Meta     `json:"meta"`
	Branches []Branch `json:"branches"`
}

// Branch :
type Branch struct {
	Commit      Commit2 `json:"commit"`
	Updatestamp string  `json:"updatestamp"`
	Branch      string  `json:"branch"`
}
