// Copyright (c) 2018.  Yoichiro Shimizu @budougumi0617

package godecov

import (
	"github.com/pkg/errors"
)

// GetOwner gets owner information from /api/gh/:owner.
func (cli *Client) GetOwner(owner string) (*OwnerResponse, error) {
	if len(owner) == 0 {
		return nil, errors.New("owner name was empty")
	}

	var o OwnerResponse
	if err := cli.get(owner, nil, &o); err != nil {
		return nil, err
	}
	return &o, nil
}
