// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 The Falco Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/docker/cli/cli/config"
	"github.com/docker/docker/pkg/homedir"
	"github.com/spf13/viper"

	drivertype "github.com/falcosecurity/falcoctl/pkg/driver/type"
)

var (
	// ConfigDir configuration directory for falcoctl.
	ConfigDir string
	// FalcoctlPath path inside the configuration directory where the falcoctl stores its config files.
	FalcoctlPath string
	// IndexesFile name of the file where the indexes info is stored. It lives under FalcoctlPath.
	IndexesFile string
	// IndexesDir is where the actual indexes are stored. It is a directory that lives under FalcoctlPath.
	IndexesDir string
	// ClientCredentialsFile name of the file where oauth client credentials are stored. It lives under FalcoctlPath.
	ClientCredentialsFile string
	// DefaultIndex is the default index for the falcosecurity organization.
	DefaultIndex Index
	// DefaultRegistryCredentialConfPath is the default path for the credential store configuration file.
	DefaultRegistryCredentialConfPath = filepath.Join(config.Dir(), "config.json")
	// DefaultDriver is the default config for the falcosecurity organization.
	DefaultDriver Driver

	// Useful regexps for parsing.

	// SemicolonSeparatedRegexp is a regexp matching semi-colon separated values, without trailing separator.
	SemicolonSeparatedRegexp = regexp.MustCompile(`^([^;]+)(;[^;]+)*$`)
	// CommaSeparatedRegexp is a regexp matching comma separated values, without trailing separator.
	CommaSeparatedRegexp = regexp.MustCompile(`^([^,]+)(,[^,]+)*$`)
)

const (
	// EnvPrefix is the prefix for all the environment variables.
	EnvPrefix = "SYSDIG"
	// ConfigPath is the path to the default config.
	ConfigPath = "/etc/sysdig/sysdig.yaml"
	// PluginsDir default path where plugins are installed.
	PluginsDir = "/usr/share/sysdig/plugins"
	// AssetsDir default path where assets are installed.
	AssetsDir = "/etc/sysdig/assets"
	// RulesfilesDir default path where rulesfiles are installed.
	RulesfilesDir = "/etc/sysdig"

	// FollowResync time interval how often it checks for newer version of the artifact.
	// Default values is set every 24 hours.
	FollowResync = time.Hour * 24

	//
	// Viper configuration keys.
	//

	// DriverKey is the Viper key for driver structure.
	DriverKey = "driver"
	// DriverTypeKey is the Viper key for the driver type.
	DriverTypeKey = "driver.type"
	// DriverVersionKey is the Viper key for the driver version.
	DriverVersionKey = "driver.version"
	// DriverReposKey is the Viper key for the driver repositories.
	DriverReposKey = "driver.repos"
	// DriverNameKey is the Viper key for the driver name.
	DriverNameKey = "driver.name"
	// DriverHostRootKey is the Viper key for the driver host root.
	DriverHostRootKey   = "driver.hostRoot"
	falcoHostRootEnvKey = "HOST_ROOT"
)

// Index represents a configured index.
type Index struct {
	Name    string `mapstructure:"name"`
	URL     string `mapstructure:"url"`
	Backend string `mapstructure:"backend"`
}

// OauthAuth represents an OAuth credential.
type OauthAuth struct {
	Registry     string `mapstructure:"registry"`
	ClientSecret string `mapstructure:"clientSecret"`
	ClientID     string `mapstructure:"clientID"`
	TokenURL     string `mapstructure:"tokenURL"`
}

// BasicAuth represents a Basic credential.
type BasicAuth struct {
	Registry string `mapstructure:"registry"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// GcpAuth represents a Gcp activation setting.
type GcpAuth struct {
	Registry string `mapstructure:"registry"`
}

// Follow represents the follower configuration.
type Follow struct {
	Every         time.Duration `mapstructure:"every"`
	Artifacts     []string      `mapstructure:"artifacts"`
	FalcoVersions string        `mapstructure:"falcoVersions"`
	RulesfilesDir string        `mapstructure:"rulesFilesDir"`
	PluginsDir    string        `mapstructure:"pluginsDir"`
	TmpDir        string        `mapstructure:"pluginsDir"`
	NoVerify      bool          `mapstructure:"noVerify"`
}

// Install represents the installer configuration.
type Install struct {
	Artifacts     []string `mapstructure:"artifacts"`
	RulesfilesDir string   `mapstructure:"rulesFilesDir"`
	PluginsDir    string   `mapstructure:"pluginsDir"`
	ResolveDeps   bool     `mapstructure:"resolveDeps"`
	NoVerify      bool     `mapstructure:"noVerify"`
}

// Driver represents the internal driver configuration (with Type string).
type Driver struct {
	Type     []string `mapstructure:"type"`
	Name     string   `mapstructure:"name"`
	Repos    []string `mapstructure:"repos"`
	Version  string   `mapstructure:"version"`
	HostRoot string   `mapstructure:"hostRoot"`
}

func init() {
	ConfigDir = filepath.Join(homedir.Get(), ".config")
	FalcoctlPath = filepath.Join(ConfigDir, "scap-driver-loader")
	DefaultDriver = Driver{
		Type:     []string{drivertype.TypeModernBpf, drivertype.TypeBpf, drivertype.TypeKmod},
		Name:     "scap",
		Repos:    []string{"https://download.sysdig.com/scap-drivers"},
		Version:  "",
		HostRoot: string(os.PathSeparator),
	}
}

// Load is used to load the config file.
func Load(path string) error {

	// Set default driver
	viper.SetDefault(DriverTypeKey, DefaultDriver.Type)
	viper.SetDefault(DriverHostRootKey, DefaultDriver.HostRoot)
	viper.SetDefault(DriverNameKey, DefaultDriver.Name)
	viper.SetDefault(DriverReposKey, DefaultDriver.Repos)
	viper.SetDefault(DriverVersionKey, DefaultDriver.Version)
	// Bind FALCOCTL_DRIVER_HOSTROOT key to HOST_ROOT,
	// so that we manage Falco HOST_ROOT variable too.
	_ = viper.BindEnv(DriverHostRootKey, falcoHostRootEnvKey)

	viper.SetEnvPrefix(EnvPrefix)

	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores. Also, consider nested keys by replacing dot with underscore.
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	// Bind to environment variables.
	viper.AutomaticEnv()

	return nil
}

// DriverTypes retrieves the driver types of the config file.
func DriverTypes() ([]string, error) {
	// manage driver.Type as ";" separated list.
	types := viper.GetStringSlice(DriverTypeKey)
	if len(types) == 1 { // in this case it might come from the env
		if !SemicolonSeparatedRegexp.MatchString(types[0]) {
			return types, fmt.Errorf("env variable not correctly set, should match %q, got %q", SemicolonSeparatedRegexp.String(), types[0])
		}
		types = strings.Split(types[0], ";")
	}
	return types, nil
}

// DriverRepos retrieves the driver repos of the config file.
func DriverRepos() ([]string, error) {
	// manage driver.Repos as ";" separated list.
	repos := viper.GetStringSlice(DriverReposKey)
	if len(repos) == 1 { // in this case it might come from the env
		if !SemicolonSeparatedRegexp.MatchString(repos[0]) {
			return repos, fmt.Errorf("env variable not correctly set, should match %q, got %q", SemicolonSeparatedRegexp.String(), repos[0])
		}
		repos = strings.Split(repos[0], ";")
	}
	return repos, nil
}

//// StoreDriver stores a driver conf in config file.
//func StoreDriver(driverCfg *Driver, configFile string) error {
//       if err := UpdateConfigFile(DriverKey, driverCfg, configFile); err != nil {
//               return fmt.Errorf("unable to update driver in the config file %q: %w", configFile, err)
//       }
//       return nil
//}
//
//// UpdateConfigFile is used to update a section of the config file.
//// We create a brand new viper instance for doing it so that we are sure that modifications
//// are scoped to the passed key with no side effects (e.g user forgot to unset one env variable for
//// another config setting, avoid to mistakenly update it).
//func UpdateConfigFile(key string, value interface{}, path string) error {
//	v := viper.New()
//	// we keep these for consistency, but not actually used
//	// since we explicitly set the filepath later
//	v.SetConfigName("falcoctl")
//	v.SetConfigType("yaml")
//
//	absolutePath, err := filepath.Abs(path)
//	if err != nil {
//		return err
//	}
//
//	v.SetConfigFile(absolutePath)
//
//	if err := v.ReadInConfig(); err != nil {
//		return fmt.Errorf("config: error reading config file: %w", err)
//	}
//
//	v.Set(key, value)
//
//	if err := v.WriteConfig(); err != nil {
//		return fmt.Errorf("unable to set key %q to config file: %w", key, err)
//	}
//
//	return nil
//}

