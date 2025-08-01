/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package names

import (
	"knative.dev/pkg/kmeta"
)

// IngressVirtualService returns the name of the VirtualService child
// resource for given Ingress that programs traffic for Ingress
// Gateways.
func IngressVirtualService(i kmeta.Accessor) string {
	return kmeta.ChildName(i.GetName(), "-ingress")
}

// MeshVirtualService returns the name of the VirtualService child
// resource for given Ingress that programs traffic for Service
// Mesh.
func MeshVirtualService(i kmeta.Accessor) string {
	return kmeta.ChildName(i.GetName(), "-mesh")
}

// DelegateVirtualService returns the name of the VirtualService child
// resource for given Ingress that programs traffic for
// delegate the route for other VirtualServices.
func DelegateVirtualService(i kmeta.Accessor) string {
	return kmeta.ChildName(i.GetName(), "-delegate")
}
