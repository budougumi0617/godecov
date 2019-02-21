// Copyright (c) 2018.  Yoichiro Shimizu @budougumi0617

package godecov

import (
	"time"

	null "github.com/mattn/go-nulltype"
)

// Author is authore information
type Author struct {
	Service   string `json:"service"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	ServiceID string `json:"service_id"`
}

// Totals is included many Codcov responses.
type Totals struct {
	C             int    `json:"C"`
	CoverageRatio string `json:"c"` // coverage ratio
	Files         int    `json:"f"` // files count
	Lines         int    `json:"n"` // lines count
	Hits          int    `json:"h"` // hits count
	Missed        int    `json:"m"` // missed count
	Partials      int    `json:"p"` // partials count
	Branches      int    `json:"b"` // branches count
	Methods       int    `json:"d"` // methods count
	Messages      int    `json:"M"` // messages count
	Sessions      int    `json:"s"` // sessions count
	// Diff []interface{} `json:"diff"`
}

// Meta has result statuses
type Meta struct {
	Status int `json:"status"`
	Limit  int `json:"limit"` // length of data will be paginated
	Page   int `json:"page"`  // number of entries are stated
}

type errorObject struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Commit defines a git commit.
type Commit struct {
	Author    Author      `json:"author"`
	Timestamp string      `json:"timestamp"`
	Totals    TotalsArray `json:"totals"`
	Commitid  string      `json:"commitid"`
	CiPassed  bool        `json:"ci_passed"`
	Message   string      `json:"message"`
}

// Commit2 defines a git commit.
type Commit2 struct {
	Author    Author  `json:"author"`
	Timestamp string  `json:"timestamp"`
	Totals    Totals2 `json:"totals"`
	Commitid  string  `json:"commitid"`
	CiPassed  bool    `json:"ci_passed"`
	Message   string  `json:"message"`
}

// Totals2 :
type Totals2 struct {
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
}

// OwnerResponse is struct for response from /api/gh/:owner.
type OwnerResponse struct {
	Repos []struct {
		Fork     interface{} `json:"fork"`
		Name     string      `json:"name"`
		Language string      `json:"language"`
		Cache    struct {
			Commit Commit `json:"commit"`
		} `json:"cache"`
		Activated    bool    `json:"activated"`
		Private      bool    `json:"private"`
		Updatestamp  string  `json:"updatestamp"`
		LatestCommit string  `json:"latest_commit"`
		Branch       string  `json:"branch"`
		Coverage     float64 `json:"coverage"`
		Repoid       string  `json:"repoid"`
	} `json:"repos"`
	Meta  Meta   `json:"meta"`
	Owner Author `json:"owner"`
}

// Owner is author information.
type Owner struct {
	Service   string          `json:"service"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Username  string          `json:"username"`
	AvatarURL null.NullString `json:"avatar_url"`
	ServiceID string          `json:"service_id"`
}

// Repo is repository information.
type Repo struct {
	UsingIntegration bool          `json:"using_integration"`
	Name             string        `json:"name"`
	Language         string        `json:"language"`
	Deleted          bool          `json:"deleted"`
	Activated        bool          `json:"activated"`
	Private          bool          `json:"private"`
	Updatestamp      time.Time     `json:"updatestamp"`
	Branch           string        `json:"branch"`
	Active           null.NullBool `json:"active"`
	ServiceID        string        `json:"service_id"`
	ImageToken       string        `json:"image_token"`
}

// SingleRepoResponse is generated struct.
type SingleRepoResponse struct {
	Commits []struct {
		Pullid       int    `json:"pullid"`
		Author       Author `json:"author"`
		Timestamp    string `json:"timestamp"`
		ParentTotals Totals `json:"parent_totals"`
		State        string `json:"state"`
		Totals       Totals `json:"totals"`
		Commitid     string `json:"commitid"`
		CiPassed     bool   `json:"ci_passed"`
		Branch       string `json:"branch"`
		Message      string `json:"message"`
		Merged       bool   `json:"merged"`
	} `json:"commits"`
	Repo   Repo  `json:"repo"`
	Meta   Meta  `json:"meta"`
	Owner  Owner `json:"owner"`
	Commit struct {
		Commitid     string      `json:"commitid"`
		Parent       string      `json:"parent"`
		Author       Author      `json:"author"`
		Deleted      interface{} `json:"deleted"`
		Timestamp    string      `json:"timestamp"`
		ParentTotals Totals      `json:"parent_totals"`
		CiPassed     bool        `json:"ci_passed"`
		Totals       Totals      `json:"totals"`
		Pullid       interface{} `json:"pullid"`
		Notified     bool        `json:"notified"`
		State        string      `json:"state"`
		Updatestamp  interface{} `json:"updatestamp"`
		Branch       string      `json:"branch"`
		Report       string      `json:"report"`
		Message      string      `json:"message"`
		Merged       bool        `json:"merged"`
	} `json:"commit"`
}

// Head id generated struct.
type Head struct {
	Author    Author `json:"author"`
	Timestamp string `json:"timestamp"`
	Totals    struct {
		C             int    `json:"C"` // Unkonwn field
		CoverageRatio string `json:"c"` // coverage ratio
		Files         int    `json:"f"` // files count
		Lines         int    `json:"n"` // lines count
		Hits          int    `json:"h"` // hits count
		Missed        int    `json:"m"` // missed count
		Partials      int    `json:"p"` // partials count
		Branches      int    `json:"b"` // branches count
		Methods       int    `json:"d"` // methods count
		Messages      int    `json:"M"` // messages count
		Sessions      int    `json:"s"` // sessions count
		// Diff []interface{} `json:"diff"`
	} `json:"totals"`
	Commitid string `json:"commitid"`
	CiPassed bool   `json:"ci_passed"`
	Message  string `json:"message"`
}

// PullsResponse is generated struct.
type PullsResponse struct {
	Owner Author `json:"owner"`
	Pulls []struct {
		Head        Head        `json:"head"`
		Issueid     int         `json:"issueid"`
		Title       string      `json:"title"`
		Author      Author      `json:"author"`
		Pullid      int         `json:"pullid"`
		State       string      `json:"state"`
		Updatestamp string      `json:"updatestamp"`
		Base        interface{} `json:"base"`
		Diff        interface{} `json:"diff"`
	} `json:"pulls"`
	Meta Meta `json:"meta"`
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
}

// CommitsResponse is generated struct.
type CommitsResponse struct {
	Commits []struct {
		Branch    string      `json:"branch"`
		Merged    interface{} `json:"merged"`
		Message   string      `json:"message"`
		Deleted   interface{} `json:"deleted"`
		Totals    string      `json:"totals"`
		Pullid    int         `json:"pullid"`
		CiPassed  bool        `json:"ci_passed"`
		Commitid  string      `json:"commitid"`
		Timestamp string      `json:"timestamp"`
		Author    string      `json:"author"`
	} `json:"commits"`
}
