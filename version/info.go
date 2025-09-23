// Copyright 2016 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package version

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"text/template"
	"time"
)

// Build information. Populated at build-time.
var (
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
	GoVersion = runtime.Version()
	GoOS      = runtime.GOOS
	GoArch    = runtime.GOARCH

	computedRevision string
	computedTags     string
)

// versionInfoTmpl contains the template used by Info.
var versionInfoTmpl = `
{{.program}}, version {{.version}} (branch: {{.branch}}, revision: {{.revision}})
  build user:       {{.buildUser}}
  build date:       {{.buildDate}}
  go version:       {{.goVersion}}
  platform:         {{.platform}}
  tags:             {{.tags}}
`

// Print returns version information.
func Print(program string) string {
	m := map[string]string{
		"program":   program,
		"version":   Version,
		"revision":  GetRevision(),
		"branch":    Branch,
		"buildUser": BuildUser,
		"buildDate": BuildDate,
		"goVersion": GoVersion,
		"platform":  GoOS + "/" + GoArch,
		"tags":      GetTags(),
	}
	t := template.Must(template.New("version").Parse(versionInfoTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", m); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}

// Info returns version, branch and revision information.
func Info() string {
	return fmt.Sprintf("(version=%s, branch=%s, revision=%s)", Version, Branch, GetRevision())
}

// BuildContext returns goVersion, platform, buildUser and buildDate information.
func BuildContext() string {
	return fmt.Sprintf("(go=%s, platform=%s, user=%s, date=%s, tags=%s)", GoVersion, GoOS+"/"+GoArch, BuildUser, BuildDate, GetTags())
}

func GetRevision() string {
	if Revision != "" {
		return Revision
	}
	return computedRevision
}

func GetTags() string {
	return computedTags
}

func PrometheusUserAgent() string {
	return ComponentUserAgent("Prometheus")
}

func ComponentUserAgent(component string) string {
	return component + "/" + Version
}

func init() {
	computedRevision, computedTags = computeRevision()
	initVersionInfo()
}

func computeRevision() (string, string) {
	var (
		rev      = "unknown"
		tags     = "unknown"
		modified bool
	)

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return rev, tags
	}
	for _, v := range buildInfo.Settings {
		if v.Key == "vcs.revision" {
			rev = v.Value
		}
		if v.Key == "vcs.modified" {
			if v.Value == "true" {
				modified = true
			}
		}
		if v.Key == "-tags" {
			tags = v.Value
		}
	}
	if modified {
		return rev + "-modified", tags
	}
	return rev, tags
}

func initVersionInfo() {
	Version = readVersionFile()
	BuildDate = time.Now().UTC().Format("20060102-15:04:05")
	BuildUser = getBuildUser()
	Branch = getBranch()
}

func readVersionFile() string {
	version := readFileFromRoot("VERSION")
	if version == "" {
		return "unknown"
	}

	versionExtra := readFileFromRoot("VERSION_EXTRA")
	if versionExtra != "" {
		return version + "-" + versionExtra
	}
	return version
}

func readFileFromRoot(filename string) string {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}

	var mainModule string
	if buildInfo.Main.Path != "" {
		mainModule = buildInfo.Main.Path
	}

	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		filePath := filepath.Join(wd, filename)
		if data, err := os.ReadFile(filePath); err == nil {
			return strings.TrimSpace(string(data))
		}

		parent := filepath.Dir(wd)
		if parent == wd {
			break
		}
		wd = parent
	}

	return ""
}

func getBuildUser() string {
	if user := os.Getenv("USER"); user != "" {
		return user
	}
	if user := os.Getenv("USERNAME"); user != "" {
		return user
	}
	return "unknown"
}

func getBranch() string {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	for _, setting := range buildInfo.Settings {
		if setting.Key == "vcs.branch" {
			return setting.Value
		}
	}
	return "unknown"
}
