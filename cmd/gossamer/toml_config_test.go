// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"path/filepath"
	"testing"

	"github.com/ChainSafe/gossamer/dot"
	"github.com/ChainSafe/gossamer/lib/utils"

	"github.com/stretchr/testify/require"
)

// TestLoadConfig tests loading a toml configuration file
func TestLoadConfig(t *testing.T) {
	cfg, cfgFile := newTestConfigWithFile(t)

	genFile := dot.NewTestGenesisRawFile(t, cfg)

	cfg.Init.Genesis = genFile

	err := dot.InitNode(cfg)
	require.NoError(t, err)

	err = loadConfig(dotConfigToToml(cfg), cfgFile)
	require.NoError(t, err)
}

// TestLoadConfigGssmr tests loading the toml configuration file for gssmr
func TestLoadConfigGssmr(t *testing.T) {
	cfg := dot.GssmrConfig()
	require.NotNil(t, cfg)

	cfg.Global.BasePath = t.TempDir()
	cfg.Init.Genesis = utils.GetGssmrGenesisPathTest(t)

	err := dot.InitNode(cfg)
	require.NoError(t, err)

	projectRootPath := utils.GetProjectRootPathTest(t)
	gssmrConfigPath := filepath.Join(projectRootPath, "./chain/gssmr/config.toml")

	err = loadConfig(dotConfigToToml(cfg), gssmrConfigPath)
	require.NoError(t, err)
}

func TestLoadConfigKusama(t *testing.T) {
	cfg := dot.KusamaConfig()
	require.NotNil(t, cfg)

	cfg.Global.BasePath = t.TempDir()
	cfg.Init.Genesis = utils.GetKusamaGenesisPath(t)

	err := dot.InitNode(cfg)
	require.NoError(t, err)

	projectRootPath := utils.GetProjectRootPathTest(t)
	kusamaConfigPath := filepath.Join(projectRootPath, "./chain/kusama/config.toml")

	err = loadConfig(dotConfigToToml(cfg), kusamaConfigPath)
	require.NoError(t, err)
}
