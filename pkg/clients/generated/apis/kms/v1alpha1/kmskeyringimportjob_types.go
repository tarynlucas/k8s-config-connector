// Copyright 2020 Google LLC
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

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KMSKeyRingImportJobSpec struct {
	/* Immutable. It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}. */
	ImportJobId string `json:"importJobId"`

	/* Immutable. The wrapping method to be used for incoming key material. Possible values: ["RSA_OAEP_3072_SHA1_AES_256", "RSA_OAEP_4096_SHA1_AES_256"]. */
	ImportMethod string `json:"importMethod"`

	/* Immutable. The KeyRing that this import job belongs to.
	Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''. */
	KeyRing string `json:"keyRing"`

	/* Immutable. The protection level of the ImportJob. This must match the protectionLevel of the
	versionTemplate on the CryptoKey you attempt to import into. Possible values: ["SOFTWARE", "HSM", "EXTERNAL"]. */
	ProtectionLevel string `json:"protectionLevel"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type KeyringimportjobAttestationStatus struct {
	/* The attestation data provided by the HSM when the key operation was performed.
	A base64-encoded string. */
	// +optional
	Content *string `json:"content,omitempty"`

	/* The format of the attestation data. */
	// +optional
	Format *string `json:"format,omitempty"`
}

type KeyringimportjobPublicKeyStatus struct {
	/* The public key, encoded in PEM format. For more information, see the RFC 7468 sections
	for General Considerations and Textual Encoding of Subject Public Key Info. */
	// +optional
	Pem *string `json:"pem,omitempty"`
}

type KMSKeyRingImportJobStatus struct {
	/* Conditions represent the latest available observations of the
	   KMSKeyRingImportJob's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Statement that was generated and signed by the key creator (for example, an HSM) at key creation time.
	Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
	Only present if the chosen ImportMethod is one with a protection level of HSM. */
	// +optional
	Attestation []KeyringimportjobAttestationStatus `json:"attestation,omitempty"`

	/* The time at which this resource is scheduled for expiration and can no longer be used.
	This is in RFC3339 text format. */
	// +optional
	ExpireTime *string `json:"expireTime,omitempty"`

	/* The resource name for this ImportJob in the format projects/* /locations/* /keyRings/* /importJobs/*. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The public key with which to wrap key material prior to import. Only returned if state is 'ACTIVE'. */
	// +optional
	PublicKey []KeyringimportjobPublicKeyStatus `json:"publicKey,omitempty"`

	/* The current state of the ImportJob, indicating if it can be used. */
	// +optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KMSKeyRingImportJob is the Schema for the kms API
// +k8s:openapi-gen=true
type KMSKeyRingImportJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSKeyRingImportJobSpec   `json:"spec,omitempty"`
	Status KMSKeyRingImportJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KMSKeyRingImportJobList contains a list of KMSKeyRingImportJob
type KMSKeyRingImportJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KMSKeyRingImportJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KMSKeyRingImportJob{}, &KMSKeyRingImportJobList{})
}