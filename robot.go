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
	"fmt"

	"github.com/opensourceways/robot-framework-lib/client"
	"github.com/opensourceways/robot-framework-lib/config"
	"github.com/opensourceways/robot-framework-lib/framework"
	"github.com/opensourceways/robot-framework-lib/utils"
	"github.com/sirupsen/logrus"
)

type iClient interface {
	CreatePRComment(owner, repo, number, comment string) (success bool)
	CreateIssueComment(owner, repo, number, comment string) (success bool)
	ListRepoAllMember(org, repo string) (result []client.User, success bool)
	ListSigInfo(org, repo string) (result []client.SigInfo, success bool)
	CheckIfPRCreateEvent(evt *client.GenericEvent) (yes bool)
	CheckIfIssueCreateEvent(evt *client.GenericEvent) (yes bool)
	AddPRLabels(org, repo, number string, labels []string) (success bool)
	AddIssueLabels(org, repo, number string, labels []string) (success bool)
	GetPullRequestChanges(org, repo, number string) (result []client.CommitFile, success bool)
	AddMemberships(org, user, permission, roleId string) (success bool)
}

type robot struct {
	cli iClient
	cnf *configuration
	log *logrus.Entry
}

func newRobot(c *configuration, token []byte, logger *logrus.Entry) *robot {
	cli := client.NewClient(token, logger)
	if cli == nil {
		return nil
	}
	return &robot{cli: cli, cnf: c, log: logger}
}

func (bot *robot) GetConfigmap() config.Configmap {
	return bot.cnf
}

func (bot *robot) RegisterEventHandler(p framework.HandlerRegister) {
	p.RegisterIssueHandler(bot.handleIssueEvent)
}

func (bot *robot) GetLogger() *logrus.Entry {
	return bot.log
}

func (bot *robot) handleIssueEvent(evt *client.GenericEvent, repoCnfPtr any, logger *logrus.Entry) {
	org, repo := utils.GetString(evt.Org), utils.GetString(evt.Repo)
	IssueAuthor := utils.GetString(evt.IssueAuthor)
	Commenter := utils.GetString(evt.Commenter)
	Author := utils.GetString(evt.Author)
	action := utils.GetString(evt.Action)

	logger.Info("org:", org, "repo:", repo, "IssueAuthor:", IssueAuthor, "Commenter:", Commenter, "Author:", Author, "Action:", action)
	if org == "openeuler" && repo == "openEuler-agreements" && action == "open" {
		fmt.Printf("add  ---- %s", IssueAuthor)
		bot.cli.AddMemberships("openeuler", IssueAuthor, "customized", "a2f75336f8a0417dac1b786f509255bc")
		bot.cli.AddMemberships("src-openeuler", IssueAuthor, "customized", "a2f75336f8a0417dac1b786f509255bc")
		return
	}
}
