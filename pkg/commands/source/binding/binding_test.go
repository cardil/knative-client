// Copyright © 2019 The Knative Authors
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

package binding

import (
	"bytes"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"
	sourcesv1 "knative.dev/eventing/pkg/apis/sources/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"knative.dev/client/pkg/commands"
	kndynamic "knative.dev/client/pkg/dynamic"
	clientv1 "knative.dev/client/pkg/sources/v1"
)

// Helper methods
var blankConfig clientcmd.ClientConfig

// Gvk used in tests
var deploymentGvk = schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "deployment"}

// TODO: Remove that blankConfig hack for tests in favor of overwriting GetConfig()
// Remove also in service_test.go
func init() {
	var err error
	blankConfig, err = clientcmd.NewClientConfigFromBytes([]byte(`kind: Config
version: v1
users:
- name: u
clusters:
- name: c
  cluster:
    server: example.com
contexts:
- name: x
  context:
    user: u
    cluster: c
current-context: x
`))
	if err != nil {
		panic(err)
	}
}

func executeSinkBindingCommand(sinkBindingClient clientv1.KnSinkBindingClient, dynamicClient kndynamic.KnDynamicClient, args ...string) (string, error) {
	knParams := &commands.KnParams{}
	knParams.ClientConfig = blankConfig

	output := new(bytes.Buffer)
	knParams.Output = output
	knParams.NewDynamicClient = func(namespace string) (kndynamic.KnDynamicClient, error) {
		return dynamicClient, nil
	}

	cmd := NewBindingCommand(knParams)
	cmd.SetArgs(args)
	cmd.SetOutput(output)

	sinkBindingClientFactory = func(config clientcmd.ClientConfig, namespace string) (clientv1.KnSinkBindingClient, error) {
		return sinkBindingClient, nil
	}
	defer cleanupSinkBindingClient()

	err := cmd.Execute()

	return output.String(), err
}

func cleanupSinkBindingClient() {
	sinkBindingClientFactory = nil
}

func createSinkBinding(name, service string, subjectGvk schema.GroupVersionKind, subjectName, namespace string, ceOverrides map[string]string) *sourcesv1.SinkBinding {
	sink := createServiceSink(service, namespace)
	builder := clientv1.NewSinkBindingBuilder(name).
		Namespace("default").
		Sink(&sink).
		SubjectGVK(&subjectGvk).
		SubjectName(subjectName).
		SubjectNamespace("default").
		CloudEventOverrides(ceOverrides, []string{})

	binding, _ := builder.Build()
	return binding
}

func createServiceSink(service, namespace string) duckv1.Destination {
	return duckv1.Destination{
		Ref: &duckv1.KReference{Name: service,
			Kind:       "Service",
			APIVersion: "serving.knative.dev/v1",
			Namespace:  namespace,
		},
	}
}
