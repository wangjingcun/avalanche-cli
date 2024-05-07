// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package metrics

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ava-labs/avalanche-cli/pkg/utils"

	"github.com/posthog/posthog-go"
	"github.com/spf13/cobra"
)

// telemetryToken value is set at build and install scripts using ldflags
var (
	telemetryToken    = ""
	telemetryInstance = "https://app.posthog.com"
)

func GetCLIVersion() string {
	wdPath, err := os.Getwd()
	if err != nil {
		return ""
	}
	versionPath := filepath.Join(wdPath, "VERSION")
	content, err := os.ReadFile(versionPath)
	if err != nil {
		return ""
	}
	return string(content)
}

func HandleTracking(cmd *cobra.Command, flags map[string]string) {
	if !cmd.HasSubCommands() && CheckCommandIsNotCompletion(cmd) {
		TrackMetrics(cmd, flags)
	}
}

func CheckCommandIsNotCompletion(cmd *cobra.Command) bool {
	result := strings.Fields(cmd.CommandPath())
	if len(result) >= 2 && result[1] == "completion" {
		return false
	}
	return true
}

func TrackMetrics(command *cobra.Command, flags map[string]string) {
	if telemetryToken == "" || utils.IsE2E() {
		return
	}
	client, _ := posthog.NewWithConfig(telemetryToken, posthog.Config{Endpoint: telemetryInstance})

	defer client.Close()

	usr, _ := user.Current() // use empty string if err
	hash := sha256.Sum256([]byte(fmt.Sprintf("%s%s", usr.Username, usr.Uid)))
	userID := base64.StdEncoding.EncodeToString(hash[:])
	telemetryProperties := make(map[string]interface{})
	telemetryProperties["command"] = command.CommandPath()
	telemetryProperties["version"] = GetCLIVersion()
	telemetryProperties["os"] = runtime.GOOS
	for propertyKey, propertyValue := range flags {
		telemetryProperties[propertyKey] = propertyValue
	}
	_ = client.Enqueue(posthog.Capture{
		DistinctId: userID,
		Event:      "cli-command",
		Properties: telemetryProperties,
	})
}
