package godecov

import "time"

// PullResponse is struct for response from /api/gh/:owner/:repo/pull/:number
type PullResponse struct {
	Pull struct {
		Head   string `json:"head"`
		Title  string `json:"title"`
		Author struct {
			Username  string `json:"username"`
			ServiceID string `json:"service_id"`
			Name      string `json:"name"`
			Service   string `json:"service"`
			Email     string `json:"email"`
		} `json:"author"`
		ComparedTo string      `json:"compared_to"`
		Pullid     int         `json:"pullid"`
		State      string      `json:"state"`
		Base       string      `json:"base"`
		Commentid  interface{} `json:"commentid"`
		Diff       interface{} `json:"diff"`
	} `json:"pull"`
	Head struct {
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
		Totals  Totals
			C    int           `json:"C"`
			B    int           `json:"b"`
			D    int           `json:"d"`
			F    int           `json:"f"`
			H    int           `json:"h"`
			M    int           `json:"M"`
			C    string        `json:"c"`
			N    int           `json:"N"`
			P    int           `json:"p"`
			M    int           `json:"m"`
			Diff []interface{} `json:"diff"`
			S    int           `json:"s"`
			N    int           `json:"n"`
		} `json:"totals"`
		Pullid      string      `json:"pullid"`
		Notified    bool        `json:"notified"`
		State       string      `json:"state"`
		Updatestamp interface{} `json:"updatestamp"`
		Branch      string      `json:"branch"`
		Report      string      `json:"report"`
		Message     string      `json:"message"`
		Merged      bool        `json:"merged"`
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
	Base string `json:"base"`
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Owner   Owner         `json:"owner"`
	Diff    interface{}   `json:"diff"`
	Changes []interface{} `json:"changes"`
}
