//go: build windows
//go:build windows
// +build windows

/*
 * Copyright (c) 2022, MegaEase
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package daemon

import (
	"golang.org/x/sys/windows"
)

const (
	processQueryLimitedInformation = 0x1000

	stillActive = 259
)

func processExists(pid int) bool {
	h, err := windows.OpenProcess(processQueryLimitedInformation, false, uint32(pid))
	if err != nil {
		return false
	}
	var c uint32
	err = windows.GetExitCodeProcess(h, &c)
	windows.Close(h)
	if err != nil {
		return c == stillActive
	}
	return true
}
