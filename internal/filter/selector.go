/*
# Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
*/

package filter

import "github.com/NVIDIA/nvidia-container-toolkit/internal/discover"

//go:generate moq -stub -out selector_mock.go . Selector

// Selector defines an interface for determining whether a specfied Device is selected
// by a particular configuration.
type Selector interface {
	Selected(discover.Device) bool
}