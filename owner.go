package godecov

// GetOwner gets owner information from /api/gh/:owner.
func (cli *Client) GetOwner(owner string) (*Owner, error) {
	var o Owner
	if err := cli.get(owner, nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}
