// Copyright (c) 2018.  Yoichiro Shimizu @budougumi0617

package godecov

// GetOwner gets owner information from /api/gh/:owner.
func (cli *Client) GetOwner(owner string) (*Owner, error) {
	var o Owner
	if err := cli.get(owner, nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}
