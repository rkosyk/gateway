// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package v1alpha1

import corev1 "k8s.io/api/core/v1"

const (
	// DefaultEnvoyReplicas is the default number of Envoy replicas.
	DefaultEnvoyReplicas = 1
	// DefaultDeploymentCPUResourceRequests for deployment cpu resource
	DefaultDeploymentCPUResourceRequests = "100m"
	// DefaultDeploymentMemoryResourceRequests for deployment memory resource
	DefaultDeploymentMemoryResourceRequests = "512Mi"
	// DefaultEnvoyAutomountServiceAccountToken is the default value of Envoy Proxy automountServiceAccountToken settings.
	DefaultEnvoyAutomountServiceAccountToken = false
)

// GroupVersionKind unambiguously identifies a Kind.
// It can be converted to k8s.io/apimachinery/pkg/runtime/schema.GroupVersionKind
type GroupVersionKind struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}

// ProviderType defines the types of providers supported by Envoy Gateway.
//
// +kubebuilder:validation:Enum=Kubernetes
type ProviderType string

const (
	// ProviderTypeKubernetes defines the "Kubernetes" provider.
	ProviderTypeKubernetes ProviderType = "Kubernetes"

	// ProviderTypeFile defines the "File" provider. This type is not implemented
	// until https://github.com/envoyproxy/gateway/issues/1001 is fixed.
	ProviderTypeFile ProviderType = "File"
)

// KubernetesDeploymentSpec defines the desired state of the Kubernetes deployment resource.
type KubernetesDeploymentSpec struct {
	// Replicas is the number of desired pods. Defaults to 1.
	//
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// PodAnnotations are the annotations that should be appended to the pods.
	// By default, no pod annotations are appended.
	//
	// +optional
	PodAnnotations map[string]string `json:"podAnnotations,omitempty"`

	// Resources required by this container.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// AutomountServiceAccountToken is the pod settings of automountServiceAccountToken. Defaults to false
	//
	// +optional
	AutomountServiceAccountToken *bool `json:"automountServiceAccountToken,omitempty"`

	// TODO: Expose config as use cases are better understood, e.g. labels.
}

// KubernetesServiceSpec defines the desired state of the Kubernetes service resource.
type KubernetesServiceSpec struct {
	// Annotations that should be appended to the service.
	// By default, no annotations are appended.
	//
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// TODO: Expose config as use cases are better understood, e.g. labels.
}

// XDSTranslatorHook defines the types of hooks that an Envoy Gateway extension may support
// for the xds-translator
//
// +kubebuilder:validation:Enum=VirtualHost;Route;HTTPListener;Translation
type XDSTranslatorHook string

const (
	XDSVirtualHost  XDSTranslatorHook = "VirtualHost"
	XDSRoute        XDSTranslatorHook = "Route"
	XDSHTTPListener XDSTranslatorHook = "HTTPListener"
	XDSTranslation  XDSTranslatorHook = "Translation"
)
