// Copyright (c) 2018.  Yoichiro Shimizu @budougumi0617

package godecov

import (
	"encoding/json"
	"errors"
	"time"
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
	// C    int           `json:"C"`
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

// TotalsArray is included many Codcov responses.
type TotalsArray struct {
	// C    int           `json:"C"`
	Files         int    `json:"f"` // [0] files count
	Lines         int    `json:"n"` // [1] lines count
	Hits          int    `json:"h"` // [2] hits count
	Partials      int    `json:"p"` // [3] partials count
	Missed        int    `json:"m"` // [4] missed count
	CoverageRatio string `json:"c"` // [5] coverage ratio
	Sessions      int    `json:"s"` // [9] sessions count
	Messages      int    `json:"M"` // messages count
	N             int    `json:"N"` // Unknown count
	Branches      int    `json:"b"` // branches count
	Methods       int    `json:"d"` // methods count
	// Diff []interface{} `json:"diff"`
}

func (ta *TotalsArray) UnmarshalJSON(data []byte) error {
	var row interface{}
	err := json.Unmarshal(data, &row)
	if err != nil {
		return err
	}
	arr, ok := row.([]interface{})
	if !ok {
		return errors.New("failed type cast")
	}
	// TODO Use NullInt
	f, ok := arr[0].(float64)
	if !ok {
		return errors.New("failed type cast f")
	}
	ta.Files = int(f)
	n, ok := arr[1].(float64)
	if !ok {
		return errors.New("failed type cast n")
	}
	ta.Lines = int(n)
	h, ok := arr[2].(float64)
	if !ok {
		return errors.New("failed type cast h")
	}
	ta.Hits = int(h)
	p, ok := arr[3].(float64)
	if !ok {
		return errors.New("failed type cast p")
	}
	ta.Partials = int(p)
	m, ok := arr[4].(float64)
	if !ok {
		return errors.New("failed type cast m")
	}
	ta.Missed = int(m)
	c, ok := arr[5].(string)
	if !ok {
		return errors.New("failed type cast c")
	}
	ta.CoverageRatio = c

	// 6-9, Unknown values

	s, ok := arr[9].(float64)
	if !ok {
		return errors.New("failed type cast s")
	}
	ta.Sessions = int(s)
	return nil
}

// TODO Define MarshalJSON for TotalsArray

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
	Author    Author `json:"author"`
	Timestamp string `json:"timestamp"`
	// Totals    Totals `json:"totals"`
	Commitid string `json:"commitid"`
	CiPassed bool   `json:"ci_passed"`
	Message  string `json:"message"`
}

// Owner is struct for response from /api/gh/:owner.
type Owner struct {
	Repos []struct {
		Fork     interface{} `json:"fork"`
		Name     string      `json:"name"`
		Language string      `json:"language"`
		Cache    struct {
			Commit struct {
				Author struct {
					Username  string      `json:"username"`
					ServiceID string      `json:"service_id"`
					Name      interface{} `json:"name"`
					Service   string      `json:"service"`
					Email     interface{} `json:"email"`
				} `json:"author"`
				Timestamp string      `json:"timestamp"`
				Totals    TotalsArray `json:"totals"`
				Commitid  string      `json:"commitid"`
				CiPassed  interface{} `json:"ci_passed"`
				Message   string      `json:"message"`
			} `json:"commit"`
		} `json:"cache"`
		Activated    bool    `json:"activated"`
		Private      bool    `json:"private"`
		Updatestamp  string  `json:"updatestamp"`
		LatestCommit string  `json:"latest_commit"`
		Branch       string  `json:"branch"`
		Coverage     float64 `json:"coverage"`
		Repoid       string  `json:"repoid"`
	} `json:"repos"`
	Meta struct {
		Status int `json:"status"`
		Limit  int `json:"limit"`
		Page   int `json:"page"`
	} `json:"meta"`
	Owner struct {
		Username    string      `json:"username"`
		Name        interface{} `json:"name"`
		Service     string      `json:"service"`
		Updatestamp string      `json:"updatestamp"`
		ServiceID   string      `json:"service_id"`
	} `json:"owner"`
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
	Meta   Meta   `json:"meta"`
	Owner  Author `json:"owner"`
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
