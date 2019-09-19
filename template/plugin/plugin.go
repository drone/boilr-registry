// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/registry"
)

// New returns a new registry plugin.
func New(param1, param2 string) registry.Plugin {
	return &plugin{
		// TODO replace or remove these configuration
		// parameters. They are for demo purposes only.
		param1: param1,
		param2: param2,
	}
}

type plugin struct {
	// TODO replace or remove these configuration
	// parameters. They are for demo purposes only.
	param1 string
	param2 string
}

func (p *plugin) List(ctx context.Context, req *registry.Request) ([]*drone.Registry, error) {
	// TODO replace or remove
	// we could only expose registry credentials to specific
	// repositories or organizations.
	if req.Repo.Namespace != p.param1 {
		return nil, nil
	}

	// TODO replace or remove
	// we can return a list of credentials for
	// multiple registries.
	credentials := []*drone.Registry{
		{
			Address:  "index.docker.io",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
	}

	return credentials, nil
}
