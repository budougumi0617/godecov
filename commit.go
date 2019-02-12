package godecov

import (
	"path"

	"github.com/pkg/errors"
)

// GetCommit gets pull requests information from /api/gh/:owner/:repo/commit/:sha.
// See also https://docs.codecov.io/reference#section-get-a-single-commit
func (cli *Client) GetCommit(owner, repo, sha string) (*CommitResponse, error) {
	// TODO Add parameters
	if len(owner) == 0 {
		return nil, errors.New("owner name was empty")
	}
	if len(repo) == 0 {
		return nil, errors.New("repository name was empty")
	}
	if len(sha) == 0 {
		return nil, errors.New("commit hash was empty")
	}

	var cr CommitResponse
	p := path.Join(owner, repo, "commit", sha)
	if err := cli.get(p, nil, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// CommitResponse is struct for response from /api/gh/:owner/:repo/commit/:sha
type CommitResponse struct {
	Owner  Owner `json:"owner"`
	Repo   Repo  `json:"repo"`
	Meta   Meta  `json:"meta"`
	Commit struct {
		Commitid string `json:"commitid"`
		Parent   string `json:"parent"`
		Author   struct {
			Username  string `json:"username"`
			ServiceID string `json:"service_id"`
			Name      string `json:"name"`
			Service   string `json:"service"`
			Email     string `json:"email"`
		} `json:"author"`
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
	} `json:"commit"`
}
