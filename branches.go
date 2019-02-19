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
	Commit struct {
		Author    Author `json:"author"`
		Timestamp string `json:"timestamp"`
		Totals    struct {
			C    int         `json:"C"`
			B    int         `json:"b"`
			D    int         `json:"d"`
			F    int         `json:"f"`
			H    int         `json:"h"`
			M    int         `json:"M"`
			Cs   string      `json:"c"`
			N    int         `json:"N"`
			P    int         `json:"p"`
			M2   int         `json:"m"`
			Diff TotalsArray `json:"diff"`
			S    int         `json:"s"`
			N3   int         `json:"n"`
		} `json:"totals"`
		Commitid string `json:"commitid"`
		CiPassed bool   `json:"ci_passed"`
		Message  string `json:"message"`
	} `json:"commit"`
	Updatestamp string `json:"updatestamp"`
	Branch      string `json:"branch"`
}
