package redhatopenshiftapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"

	"github.com/Azure/go-autorest/autorest"

	"github.com/Azure/ARO-RP/pkg/client/services/redhatopenshift/mgmt/2023-09-04/redhatopenshift"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result redhatopenshift.OperationListPage, err error)
	ListComplete(ctx context.Context) (result redhatopenshift.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*redhatopenshift.OperationsClient)(nil)

// OpenShiftVersionsClientAPI contains the set of methods on the OpenShiftVersionsClient type.
type OpenShiftVersionsClientAPI interface {
	List(ctx context.Context, location string) (result redhatopenshift.OpenShiftVersionListPage, err error)
	ListComplete(ctx context.Context, location string) (result redhatopenshift.OpenShiftVersionListIterator, err error)
}

var _ OpenShiftVersionsClientAPI = (*redhatopenshift.OpenShiftVersionsClient)(nil)

// OpenShiftClustersClientAPI contains the set of methods on the OpenShiftClustersClient type.
type OpenShiftClustersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters redhatopenshift.OpenShiftCluster) (result redhatopenshift.OpenShiftClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.OpenShiftClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.OpenShiftCluster, err error)
	List(ctx context.Context) (result redhatopenshift.OpenShiftClusterListPage, err error)
	ListComplete(ctx context.Context) (result redhatopenshift.OpenShiftClusterListIterator, err error)
	ListAdminCredentials(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.OpenShiftClusterAdminKubeconfig, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result redhatopenshift.OpenShiftClusterListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result redhatopenshift.OpenShiftClusterListIterator, err error)
	ListCredentials(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.OpenShiftClusterCredentials, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, parameters redhatopenshift.OpenShiftClusterUpdate) (result redhatopenshift.OpenShiftClustersUpdateFuture, err error)
}

var _ OpenShiftClustersClientAPI = (*redhatopenshift.OpenShiftClustersClient)(nil)

// MachinePoolsClientAPI contains the set of methods on the MachinePoolsClient type.
type MachinePoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters redhatopenshift.MachinePool) (result redhatopenshift.MachinePool, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string) (result redhatopenshift.MachinePool, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.MachinePoolListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.MachinePoolListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters redhatopenshift.MachinePoolUpdate) (result redhatopenshift.MachinePool, err error)
}

var _ MachinePoolsClientAPI = (*redhatopenshift.MachinePoolsClient)(nil)

// SecretsClientAPI contains the set of methods on the SecretsClient type.
type SecretsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters redhatopenshift.Secret) (result redhatopenshift.Secret, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string) (result redhatopenshift.Secret, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.SecretListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.SecretListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters redhatopenshift.SecretUpdate) (result redhatopenshift.Secret, err error)
}

var _ SecretsClientAPI = (*redhatopenshift.SecretsClient)(nil)

// SyncIdentityProvidersClientAPI contains the set of methods on the SyncIdentityProvidersClient type.
type SyncIdentityProvidersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters redhatopenshift.SyncIdentityProvider) (result redhatopenshift.SyncIdentityProvider, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string) (result redhatopenshift.SyncIdentityProvider, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.SyncIdentityProviderListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.SyncIdentityProviderListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters redhatopenshift.SyncIdentityProviderUpdate) (result redhatopenshift.SyncIdentityProvider, err error)
}

var _ SyncIdentityProvidersClientAPI = (*redhatopenshift.SyncIdentityProvidersClient)(nil)

// SyncSetsClientAPI contains the set of methods on the SyncSetsClient type.
type SyncSetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters redhatopenshift.SyncSet) (result redhatopenshift.SyncSet, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string) (result redhatopenshift.SyncSet, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.SyncSetListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.SyncSetListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters redhatopenshift.SyncSetUpdate) (result redhatopenshift.SyncSet, err error)
}

var _ SyncSetsClientAPI = (*redhatopenshift.SyncSetsClient)(nil)
