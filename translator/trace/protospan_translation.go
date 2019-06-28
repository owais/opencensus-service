// Copyright 2018, OpenTelemetry Authors
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

package tracetranslator

// Some of the keys used to represent OC proto constructs as tags or annotations in other formats.
const (
	AnnotationDescriptionKey = "description"

	MessageEventIDKey               = "message.id"
	MessageEventTypeKey             = "message.type"
	MessageEventCompressedSizeKey   = "message.compressed_size"
	MessageEventUncompressedSizeKey = "message.uncompressed_size"

	TagStatusCode       = "status.code"
	TagStatusMsg        = "status.message"
	TagHTTPStatusCode   = "http.status_code"
	TagHTTPStatusMsg    = "http.status_message"
	TagZipkinCensusCode = "census.status_code"
	TagZipkinCensusMsg  = "census.status_description"
)
