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

package v1

import (
	"strings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LocatorSpec defines the desired state of Locator
type LocatorSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	
	// Specifies the number of Locators desired in the cluster, validation to constrain to 1-3
	// +kubebuilder:validation:Enum=1;3
	Replicas	int	`json:"replicas"`

	
	Image		string	`json:"image,omitempty"`
	
	// On by default, but optionally can be turned off
	// +optional
	Metrics		*bool	`json:"metrics,omitempty"`

	// On by default but optionally can be turned off
	// +optional
	Pulse		*bool  	`json:"pulse,omitempty"`

	// +optional
	NamePrefix	string 	`json:"name-prefix,omitempty"`

}

// LocatorStatus defines the observed state of Locator
type LocatorStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
	LocatorStatus string `json:"locatorStatus,omitempty"`
}

// +kubebuilder:object:root=true

// We want a status subresource so that we behave like built-in kubernetes type
// +kubebuilder:subresource:status

// Locator is the Schema for the locators API
type Locator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LocatorSpec   `json:"spec,omitempty"`
	Status LocatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocatorList contains a list of Locator
type LocatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Locator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Locator{}, &LocatorList{})
}
