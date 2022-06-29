// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package cmd

import "time"

const (
	BaseDirName = ".avalanche-cli"
	subnetEvm   = "SubnetEVM"
	customVM    = "Custom"

	forceFlag = "force"

	maxLogFileSize   = 4
	maxNumOfLogFiles = 5
	retainOldFiles   = 0 // retain all old log files

	healthCheckInterval = 100 * time.Millisecond
)
