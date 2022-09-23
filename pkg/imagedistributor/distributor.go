// Copyright © 2022 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package imagedistributor

import (
	"net"
)

type Interface interface {
	// Distribute each files under mounted cluster image directory to target hosts.
	Distribute(imageName string, hosts []net.IP) error
	// Restore will do some clean works via infra driver, like delete rootfs.
	Restore(targetDir string, hosts []net.IP) error
}