package main

import (
	"testing"

	"github.com/concourse/concourse/atc"
)

func TestGitHubUriConstruction(t *testing.T) {
	cases := []struct {
		ResourceConfig atc.ResourceConfig
		Result         string
	}{
		{atc.ResourceConfig{Source: atc.Source{"uri": "https://github.com/some/repo"}}, "https://github.com/some/repo"}, // Test for concourse/git-resource
		{atc.ResourceConfig{Source: atc.Source{"repository": "some/repo"}}, "https://github.com/some/repo"},             // Test for telia-oss/github-pr-resource
		{atc.ResourceConfig{Source: atc.Source{"owner": "some", "repository": "repo"}}, "https://github.com/some/repo"}, // Test for concourse/github-release-resource
	}

	for nr, c := range cases {
		uri, ok := ConstructGitHubUriFromConfig(c.ResourceConfig)
		if !ok || uri != c.Result {
			t.Errorf("Test case %d failed.", nr+1)
		}
	}
}
