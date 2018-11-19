package godecov

// GetOwner gets owner information from /api/gh/:owner.
func (cli *Client) GetOwner(owner string) (*Owner, error) {
	resp, err := cli.get(owner, nil)
	if err != nil {
		return nil, err
	}
	var o Owner
	return cli.decodeJSON(resp, &o)
}
