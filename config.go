// Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"github.com/opensourceways/robot-framework-lib/config"
)

type configuration struct {
	ConfigItems []repoConfig `json:"config_items,omitempty"`
	// Sig information url.
	SigInfoURL string `json:"sig_info_url" required:"true"`
	// Community name used as a request parameter to getRepoConfig sig information.
	CommunityName  string   `json:"community_name" required:"true"`
	ExcludeUser    []string `json:"exclude_user,omitempty"`
	UserMarkFormat string   `json:"user_mark_format" required:"true"`
}

// Validate to check the configmap data's validation, returns an error if invalid
func (c *configuration) Validate() error {
	err := config.ValidateRequiredConfig(*c)
	if err != nil {
		return err
	}
	err = config.ValidateConfigItems(c.ConfigItems)
	if err != nil {
		return err
	}
	return nil
}

type repoConfig struct {
	// Repos are either in the form of org/repos or just org.
	Repos []string `json:"repos" required:"true"`
	// ExcludedRepos are in the form of org/repo.
	ExcludedRepos  []string `json:"excluded_repos,omitempty"`
	CommunityName  string   `json:"community_name" required:"true"`
	CommandLink    string   `json:"command_link" required:"true"`
	SigInfoLink    string   `json:"sig_info_link" required:"true"`
	WelcomeMessage []string `json:"welcome_message" required:"true"`
}
