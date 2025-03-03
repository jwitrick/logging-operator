// Copyright © 2023 Kube logging authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package output

// +name:"Openobserve"
// +weight:"200"
type _hugoOpenobserve interface{} //nolint:deadcode,unused

// +docName:"Sending messages over Openobserve"
/*
## Example

{{< highlight yaml >}}
apiVersion: logging.banzaicloud.io/v1beta1
kind: SyslogNGOutput
metadata:
  name: openobserve
spec:
  openobserve:
    url: "https://some-openobserve-endpoint"
    port: 5040
    organization: "default"
    stream: "default"
    user: "username"
    password:
      valueFrom:
        secretKeyRef:
          name: openobserve
          key: password
{{</ highlight >}}
More information at https://axoflow.com/docs/axosyslog-core/chapter-destinations/openobserve/
*/
type _docOpenobserve interface{} //nolint:deadcode,unused

// +name:"Openobserve"
// +url:"https://axoflow.com/docs/axosyslog-core/chapter-destinations/openobserve/"
// +description:"Sending messages over Openobserve"
// +status:"Testing"
type _metaOpenobserve interface{} //nolint:deadcode,unused

// +kubebuilder:object:generate=true
type OpenobserveOutput struct {
	HTTPOutput `json:",inline"`
	// Name of the organization in Openobserve.
	Organization string `json:"organization,omitempty"`
	// Name of the stream in Openobserve.
	Stream string `json:"stream,omitempty"`
	// Arguments to the `$format-json()` template function.
	// Default: --scope rfc5424 --exclude DATE --key ISODATE @timestamp=${ISODATE}"
	Record string `json:"record,omitempty"`
}

func (o *OpenobserveOutput) BeforeRender() {
	if o.Organization == "" {
		o.Organization = "default"
	}

	if o.Stream == "" {
		o.Stream = "default"
	}

}
