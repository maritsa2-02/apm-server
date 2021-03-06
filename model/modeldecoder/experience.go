// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package modeldecoder

import (
	"github.com/elastic/apm-server/model"
	"github.com/elastic/apm-server/model/modeldecoder/field"
)

func decodeUserExperience(input map[string]interface{}, fieldName field.MapperFunc, out *model.UserExperience) {
	if input == nil {
		return
	}

	if !decodeFloat64(input, "cls", &out.CumulativeLayoutShift) {
		out.CumulativeLayoutShift = -1
	}
	if !decodeFloat64(input, "fid", &out.FirstInputDelay) {
		out.FirstInputDelay = -1
	}
	if !decodeFloat64(input, "tbt", &out.TotalBlockingTime) {
		out.TotalBlockingTime = -1
	}

	out.Longtask.Count = -1
	if input := getObject(input, fieldName("longtask")); input != nil {
		decodeInt(input, "count", &out.Longtask.Count)
		decodeFloat64(input, "sum", &out.Longtask.Sum)
		decodeFloat64(input, "max", &out.Longtask.Max)
	}
}
