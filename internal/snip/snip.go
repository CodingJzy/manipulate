// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snip

import (
	"fmt"
	"strings"
)

// Snip snips the given token from the given error.
func Snip(err error, token string) error {

	if len(token) == 0 || err == nil {
		return err
	}

	return fmt.Errorf("%s",
		strings.Replace(
			err.Error(),
			token,
			"[snip]",
			-1),
	)
}
