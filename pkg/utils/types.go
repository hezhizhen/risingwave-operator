// Copyright 2022 Singularity Data
//
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

package utils

// WebhookType defines an immutable type for a webhook. Use NewWebhookType to instantiate.
type WebhookType string

// Valid webhook types.
const (
	ValidatingWebhookType WebhookType = "validating"
	MutatingWebhookType   WebhookType = "mutating"
)

// IsValidating is true if wt is a validating webhook, else false.
func (wt WebhookType) IsValidating() bool {
	return wt == ValidatingWebhookType
}

// IsMutating is true if wt is a mutating webhook, else false.
func (wt WebhookType) IsMutating() bool {
	return !wt.IsValidating()
}

func (wt WebhookType) String() string {
	return string(wt)
}
