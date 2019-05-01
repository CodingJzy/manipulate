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

package manipvortex

import (
	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

// Transaction is the event that captures the transaction for later processing. It is
// also the structure stored in the transaction logs.
type Transaction struct {
	Date     time.Time
	mctx     manipulate.Context
	Object   elemental.Identifiable
	Method   elemental.Operation
	Deadline time.Time
}
