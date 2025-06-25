/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1alpha1 // github.com/openshift-online/ocm-api-model/clientapi/arohcp/v1alpha1

// AwsBackupConfig represents the values of the 'aws_backup_config' type.
//
// Backup configuration for AWS clusters
type AwsBackupConfig struct {
	bitmap_             uint32
	s3Bucket            string
	accountId           string
	identityProviderArn string
	roleArn             string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *AwsBackupConfig) Empty() bool {
	return o == nil || o.bitmap_ == 0
}

// S3Bucket returns the value of the 'S3_bucket' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Name of the S3 bucket used to save the backup
func (o *AwsBackupConfig) S3Bucket() string {
	if o != nil && o.bitmap_&1 != 0 {
		return o.s3Bucket
	}
	return ""
}

// GetS3Bucket returns the value of the 'S3_bucket' attribute and
// a flag indicating if the attribute has a value.
//
// Name of the S3 bucket used to save the backup
func (o *AwsBackupConfig) GetS3Bucket() (value string, ok bool) {
	ok = o != nil && o.bitmap_&1 != 0
	if ok {
		value = o.s3Bucket
	}
	return
}

// AccountId returns the value of the 'account_id' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// ID of the AWS Disaster Recovery (DR) account
func (o *AwsBackupConfig) AccountId() string {
	if o != nil && o.bitmap_&2 != 0 {
		return o.accountId
	}
	return ""
}

// GetAccountId returns the value of the 'account_id' attribute and
// a flag indicating if the attribute has a value.
//
// ID of the AWS Disaster Recovery (DR) account
func (o *AwsBackupConfig) GetAccountId() (value string, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.accountId
	}
	return
}

// IdentityProviderArn returns the value of the 'identity_provider_arn' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// ARN of the identity provider created in the Disaster Recovery (DR) account for the Management Cluster
func (o *AwsBackupConfig) IdentityProviderArn() string {
	if o != nil && o.bitmap_&4 != 0 {
		return o.identityProviderArn
	}
	return ""
}

// GetIdentityProviderArn returns the value of the 'identity_provider_arn' attribute and
// a flag indicating if the attribute has a value.
//
// ARN of the identity provider created in the Disaster Recovery (DR) account for the Management Cluster
func (o *AwsBackupConfig) GetIdentityProviderArn() (value string, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.identityProviderArn
	}
	return
}

// RoleArn returns the value of the 'role_arn' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// ARN of the role used by the CS Trusted Account to gain access to the Disaster Recovery (DR) account
func (o *AwsBackupConfig) RoleArn() string {
	if o != nil && o.bitmap_&8 != 0 {
		return o.roleArn
	}
	return ""
}

// GetRoleArn returns the value of the 'role_arn' attribute and
// a flag indicating if the attribute has a value.
//
// ARN of the role used by the CS Trusted Account to gain access to the Disaster Recovery (DR) account
func (o *AwsBackupConfig) GetRoleArn() (value string, ok bool) {
	ok = o != nil && o.bitmap_&8 != 0
	if ok {
		value = o.roleArn
	}
	return
}

// AwsBackupConfigListKind is the name of the type used to represent list of objects of
// type 'aws_backup_config'.
const AwsBackupConfigListKind = "AwsBackupConfigList"

// AwsBackupConfigListLinkKind is the name of the type used to represent links to list
// of objects of type 'aws_backup_config'.
const AwsBackupConfigListLinkKind = "AwsBackupConfigListLink"

// AwsBackupConfigNilKind is the name of the type used to nil lists of objects of
// type 'aws_backup_config'.
const AwsBackupConfigListNilKind = "AwsBackupConfigListNil"

// AwsBackupConfigList is a list of values of the 'aws_backup_config' type.
type AwsBackupConfigList struct {
	href  string
	link  bool
	items []*AwsBackupConfig
}

// Len returns the length of the list.
func (l *AwsBackupConfigList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Items sets the items of the list.
func (l *AwsBackupConfigList) SetLink(link bool) {
	l.link = link
}

// Items sets the items of the list.
func (l *AwsBackupConfigList) SetHREF(href string) {
	l.href = href
}

// Items sets the items of the list.
func (l *AwsBackupConfigList) SetItems(items []*AwsBackupConfig) {
	l.items = items
}

// Items returns the items of the list.
func (l *AwsBackupConfigList) Items() []*AwsBackupConfig {
	if l == nil {
		return nil
	}
	return l.items
}

// Empty returns true if the list is empty.
func (l *AwsBackupConfigList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *AwsBackupConfigList) Get(i int) *AwsBackupConfig {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *AwsBackupConfigList) Slice() []*AwsBackupConfig {
	var slice []*AwsBackupConfig
	if l == nil {
		slice = make([]*AwsBackupConfig, 0)
	} else {
		slice = make([]*AwsBackupConfig, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *AwsBackupConfigList) Each(f func(item *AwsBackupConfig) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *AwsBackupConfigList) Range(f func(index int, item *AwsBackupConfig) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
