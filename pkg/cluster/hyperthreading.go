package cluster

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/util/feature"
)

func (m *manager) populateHyperthreadingMode(ctx context.Context) error {
	subProperties := m.subscriptionDoc.Subscription.Properties
	var hyperthreadingMode api.HyperthreadingMode

	if feature.IsRegisteredForFeature(subProperties, api.FeatureFlagHyperthreadingModeDisabled) {
		hyperthreadingMode = api.HyperthreadingDisabled
	} else {
		hyperthreadingMode = api.HyperthreadingEnabled
	}
	return patchHyperthreadingMode(m, ctx, hyperthreadingMode)
}

func patchHyperthreadingMode(m *manager, ctx context.Context, hyperthreadingMode api.HyperthreadingMode) error {
	var err error
	m.doc, err = m.db.PatchWithLease(ctx, m.doc.Key, func(doc *api.OpenShiftClusterDocument) error {
		doc.OpenShiftCluster.Properties.ClusterProfile.HyperthreadingMode = hyperthreadingMode
		return nil
	})
	return err
}
