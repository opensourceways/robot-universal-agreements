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
	"flag"
	"github.com/opensourceways/robot-framework-lib/framework"
	"github.com/opensourceways/robot-framework-lib/testdata"
	"github.com/opensourceways/robot-framework-lib/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	commandPort             = "--port=8511"
	commandExecFile         = "****"
	commandConfigFilePrefix = "--config-file="
	commandTokenFilePrefix  = "--token-path="
	commandDelToken         = "--del-token=false"
	commandHandlePath       = "--handle-path=gitcode-hook"
	dir                     = string(os.PathSeparator) + "testdata" + string(os.PathSeparator)
	configMap               = dir + "config.yaml"
	configMap1              = dir + "config1.yaml"
	token                   = dir + "token"
	sig1                    = dir + "sig1.json"
	sig2                    = dir + "sig2.json"
	sig3                    = dir + "sig3.json"
	sig4                    = dir + "sig4.json"
	sig5                    = dir + "sig5.json"
	comment1                = dir + "comment1.md"
	comment2                = dir + "comment2.md"
	comment4                = dir + "comment4.md"
	comment5                = dir + "comment5.md"
	comment6                = dir + "comment6.md"
	comment7                = dir + "comment7.md"
)

func TestGatherOptions(t *testing.T) {
	logger := framework.NewLogger()
	args := []string{
		commandExecFile,
		commandPort,
		commandConfigFilePrefix + findTestdata(nil, configMap1),
		commandHandlePath,
		commandTokenFilePrefix + findTestdata(t, token),
		commandDelToken,
	}
	opt := new(robotOptions)
	opt.gatherOptions(flag.NewFlagSet(args[0], flag.ExitOnError), logger, args[1:]...)
	assert.Equal(t, true, opt.service.Interrupt)

	args[2] = commandConfigFilePrefix + findTestdata(t, configMap)

	opt = new(robotOptions)
	opt.gatherOptions(flag.NewFlagSet(args[0], flag.ExitOnError), logger, args[1:]...)
	assert.Equal(t, false, opt.service.Interrupt)
	assert.Equal(t, "gitcode-hook", opt.service.HandlePath)
	want := &configuration{}
	_ = utils.LoadFromYaml(findTestdata(t, configMap), want)
	assert.Equal(t, *want, *opt.service.ConfigmapAgentValue.GetConfigmap().(*configuration))
	assert.Equal(t, "gf112421415123123asdada", string(opt.service.TokenValue))
}

func findTestdata(t *testing.T, path string) string {
	workDir, _ := os.Getwd()
	return testdata.FindTestDataWithDir(t, workDir, path)
}
