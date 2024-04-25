// Copyright The TBox Authors. All rights reserved.
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

package journal

import (
	"github.com/sunsingerus/tbox/pkg/api/common"
)

const (
	// EndpointUnknown specifies initial unknown endpoint
	EndpointUnknown int32 = 0

	// NoPlane section
	NoPlane       int32 = 0
	NoPlaneCustom       = NoPlane + 2000

	// ControlPlane section
	ControlPlane       int32 = 0
	EndpointTasks            = ControlPlane + 100
	ControlPlaneCustom       = ControlPlane + 2000

	// DataPlane section
	DataPlane               int32 = 3000
	EndpointDataChunks            = DataPlane + 100
	EndpointUploadObject          = DataPlane + 200
	EndpointUploadObjects         = DataPlane + 300
	EndpointDownloadObject        = DataPlane + 400
	EndpointDownloadObjects       = DataPlane + 500
	DataPlaneCustom               = DataPlane + 2000

	// HealthPlane section
	HealthPlane           int32 = 6000
	EndpointMetrics             = HealthPlane + 100
	EndpointStatusObject        = HealthPlane + 200
	EndpointStatusObjects       = HealthPlane + 300
	HealthPlaneCustom           = HealthPlane + 2000

	// ReportsPlane section
	ReportsPlane       int32 = 9000
	EndpointReport           = ReportsPlane + 100
	EndpointReports          = ReportsPlane + 200
	ReportsPlaneCustom       = ReportsPlane + 2000
)

var (
	EndpointTypeEnum = common.NewEnum()
)

func init() {
	EndpointTypeEnum.MustRegister("EndpointUnknown", EndpointUnknown)
	// Control Plane
	EndpointTypeEnum.MustRegister("EndpointTasks", EndpointTasks)
	// Data Plane
	EndpointTypeEnum.MustRegister("EndpointDataChunks", EndpointDataChunks)
	EndpointTypeEnum.MustRegister("EndpointUploadObject", EndpointUploadObject)
	EndpointTypeEnum.MustRegister("EndpointUploadObjects", EndpointUploadObjects)
	EndpointTypeEnum.MustRegister("EndpointDownloadObject", EndpointDownloadObject)
	EndpointTypeEnum.MustRegister("EndpointDownloadObjects", EndpointDownloadObjects)
	// Health Plane
	EndpointTypeEnum.MustRegister("EndpointMetrics", EndpointMetrics)
	EndpointTypeEnum.MustRegister("EndpointStatusObject", EndpointStatusObject)
	EndpointTypeEnum.MustRegister("EndpointStatusObjects", EndpointStatusObjects)
	// Reports Plane
	EndpointTypeEnum.MustRegister("EndpointReport", EndpointReport)
	EndpointTypeEnum.MustRegister("EndpointReports", EndpointReports)
}
