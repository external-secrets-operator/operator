/*


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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExternalConfigMapSpec defines the desired state of ExternalConfigMap
type ExternalConfigMapSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name of the ExternalBackend resource that is used to get values.
	// +kubebuilder:validation:MinLength=1
	BackendName string `json:"backendName"`

	// Keys in the backend that hold values.
	// +kubebuilder:validation:MinItems=1
	Keys []string `json:"keys"`

	// Golang template that generates the entire key-value map
	// +optional
	Template string `json:"template"`
}

// ExternalConfigMapStatus defines the observed state of ExternalConfigMap
type ExternalConfigMapStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ExternalConfigMap is the Schema for the externalconfigmaps API
type ExternalConfigMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalConfigMapSpec   `json:"spec,omitempty"`
	Status ExternalConfigMapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalConfigMapList contains a list of ExternalConfigMap
type ExternalConfigMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalConfigMap `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExternalConfigMap{}, &ExternalConfigMapList{})
}
