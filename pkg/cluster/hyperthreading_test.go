package cluster

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Azure/ARO-RP/pkg/api"
	testdatabase "github.com/Azure/ARO-RP/test/database"
)

func TestPopulateHyperthreadingMode(t *testing.T) {
	ctx := context.Background()

	// Define the DB instance we will use to run the PatchWithLease function
	key := "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup/providers/Microsoft.RedHatOpenShift/openShiftClusters/resourceName"

	// Run tests
	for _, tt := range []struct {
		name                       string
		m                          manager
		expectedHyperthreadingMode api.HyperthreadingMode
		expectedErr                error
	}{
		{
			name: "hyperthreading disabled when registered for feature in subscription",
			m: manager{
				doc: &api.OpenShiftClusterDocument{
					Key: strings.ToLower(key),
					OpenShiftCluster: &api.OpenShiftCluster{
						ID: key,
						Properties: api.OpenShiftClusterProperties{
							ProvisioningState: api.ProvisioningStateSucceeded,
							ClusterProfile:    api.ClusterProfile{},
						},
					},
				},
				subscriptionDoc: &api.SubscriptionDocument{
					Subscription: &api.Subscription{
						Properties: &api.SubscriptionProperties{
							RegisteredFeatures: []api.RegisteredFeatureProfile{
								{
									Name:  api.FeatureFlagHyperthreadingModeDisabled,
									State: "Registered",
								},
							},
						},
					},
				},
			},
			expectedHyperthreadingMode: api.HyperthreadingDisabled,
			expectedErr:                nil,
		},
		{
			name: "hyperthreading enabled when not registered for feature in subscription",
			m: manager{
				doc: &api.OpenShiftClusterDocument{
					Key: strings.ToLower(key),
					OpenShiftCluster: &api.OpenShiftCluster{
						ID: key,
						Properties: api.OpenShiftClusterProperties{
							ProvisioningState: api.ProvisioningStateSucceeded,
							ClusterProfile:    api.ClusterProfile{},
						},
					},
				},
				subscriptionDoc: &api.SubscriptionDocument{
					Subscription: &api.Subscription{
						Properties: &api.SubscriptionProperties{
							RegisteredFeatures: []api.RegisteredFeatureProfile{
								{},
							},
						},
					},
				},
			},
			expectedHyperthreadingMode: api.HyperthreadingEnabled,
			expectedErr:                nil,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// Create the DB to test the cluster
			openShiftClustersDatabase, _ := testdatabase.NewFakeOpenShiftClusters()
			fixture := testdatabase.NewFixture().WithOpenShiftClusters(openShiftClustersDatabase)
			fixture.AddOpenShiftClusterDocuments(tt.m.doc)
			err := fixture.Create()
			if err != nil {
				t.Fatal(err)
			}
			tt.m.db = openShiftClustersDatabase

			// Run populateHyperthreadingMode and assert the correct results
			err = tt.m.populateHyperthreadingMode(ctx)
			assert.Equal(t, tt.expectedErr, err, "Unexpected error exception")
			assert.Equal(t, tt.expectedHyperthreadingMode, tt.m.doc.OpenShiftCluster.Properties.ClusterProfile.HyperthreadingMode, "Hyperthreading was not populated as expected exception")
		})
	}
}
