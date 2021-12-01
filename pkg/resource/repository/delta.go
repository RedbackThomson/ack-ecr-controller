// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package repository

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.EncryptionConfiguration, b.ko.Spec.EncryptionConfiguration) {
		delta.Add("Spec.EncryptionConfiguration", a.ko.Spec.EncryptionConfiguration, b.ko.Spec.EncryptionConfiguration)
	} else if a.ko.Spec.EncryptionConfiguration != nil && b.ko.Spec.EncryptionConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.EncryptionConfiguration.EncryptionType, b.ko.Spec.EncryptionConfiguration.EncryptionType) {
			delta.Add("Spec.EncryptionConfiguration.EncryptionType", a.ko.Spec.EncryptionConfiguration.EncryptionType, b.ko.Spec.EncryptionConfiguration.EncryptionType)
		} else if a.ko.Spec.EncryptionConfiguration.EncryptionType != nil && b.ko.Spec.EncryptionConfiguration.EncryptionType != nil {
			if *a.ko.Spec.EncryptionConfiguration.EncryptionType != *b.ko.Spec.EncryptionConfiguration.EncryptionType {
				delta.Add("Spec.EncryptionConfiguration.EncryptionType", a.ko.Spec.EncryptionConfiguration.EncryptionType, b.ko.Spec.EncryptionConfiguration.EncryptionType)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.EncryptionConfiguration.KMSKey, b.ko.Spec.EncryptionConfiguration.KMSKey) {
			delta.Add("Spec.EncryptionConfiguration.KMSKey", a.ko.Spec.EncryptionConfiguration.KMSKey, b.ko.Spec.EncryptionConfiguration.KMSKey)
		} else if a.ko.Spec.EncryptionConfiguration.KMSKey != nil && b.ko.Spec.EncryptionConfiguration.KMSKey != nil {
			if *a.ko.Spec.EncryptionConfiguration.KMSKey != *b.ko.Spec.EncryptionConfiguration.KMSKey {
				delta.Add("Spec.EncryptionConfiguration.KMSKey", a.ko.Spec.EncryptionConfiguration.KMSKey, b.ko.Spec.EncryptionConfiguration.KMSKey)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ImageScanningConfiguration, b.ko.Spec.ImageScanningConfiguration) {
		delta.Add("Spec.ImageScanningConfiguration", a.ko.Spec.ImageScanningConfiguration, b.ko.Spec.ImageScanningConfiguration)
	} else if a.ko.Spec.ImageScanningConfiguration != nil && b.ko.Spec.ImageScanningConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ImageScanningConfiguration.ScanOnPush, b.ko.Spec.ImageScanningConfiguration.ScanOnPush) {
			delta.Add("Spec.ImageScanningConfiguration.ScanOnPush", a.ko.Spec.ImageScanningConfiguration.ScanOnPush, b.ko.Spec.ImageScanningConfiguration.ScanOnPush)
		} else if a.ko.Spec.ImageScanningConfiguration.ScanOnPush != nil && b.ko.Spec.ImageScanningConfiguration.ScanOnPush != nil {
			if *a.ko.Spec.ImageScanningConfiguration.ScanOnPush != *b.ko.Spec.ImageScanningConfiguration.ScanOnPush {
				delta.Add("Spec.ImageScanningConfiguration.ScanOnPush", a.ko.Spec.ImageScanningConfiguration.ScanOnPush, b.ko.Spec.ImageScanningConfiguration.ScanOnPush)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ImageTagMutability, b.ko.Spec.ImageTagMutability) {
		delta.Add("Spec.ImageTagMutability", a.ko.Spec.ImageTagMutability, b.ko.Spec.ImageTagMutability)
	} else if a.ko.Spec.ImageTagMutability != nil && b.ko.Spec.ImageTagMutability != nil {
		if *a.ko.Spec.ImageTagMutability != *b.ko.Spec.ImageTagMutability {
			delta.Add("Spec.ImageTagMutability", a.ko.Spec.ImageTagMutability, b.ko.Spec.ImageTagMutability)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.LifecyclePolicy, b.ko.Spec.LifecyclePolicy) {
		delta.Add("Spec.LifecyclePolicy", a.ko.Spec.LifecyclePolicy, b.ko.Spec.LifecyclePolicy)
	} else if a.ko.Spec.LifecyclePolicy != nil && b.ko.Spec.LifecyclePolicy != nil {
		if *a.ko.Spec.LifecyclePolicy != *b.ko.Spec.LifecyclePolicy {
			delta.Add("Spec.LifecyclePolicy", a.ko.Spec.LifecyclePolicy, b.ko.Spec.LifecyclePolicy)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RegistryID, b.ko.Spec.RegistryID) {
		delta.Add("Spec.RegistryID", a.ko.Spec.RegistryID, b.ko.Spec.RegistryID)
	} else if a.ko.Spec.RegistryID != nil && b.ko.Spec.RegistryID != nil {
		if *a.ko.Spec.RegistryID != *b.ko.Spec.RegistryID {
			delta.Add("Spec.RegistryID", a.ko.Spec.RegistryID, b.ko.Spec.RegistryID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.Tags, b.ko.Spec.Tags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}
