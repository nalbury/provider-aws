// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	acl "github.com/upbound/provider-aws/internal/controller/memorydb/acl"
	cluster "github.com/upbound/provider-aws/internal/controller/memorydb/cluster"
	parametergroup "github.com/upbound/provider-aws/internal/controller/memorydb/parametergroup"
	snapshot "github.com/upbound/provider-aws/internal/controller/memorydb/snapshot"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/memorydb/subnetgroup"
)

// Setup_memorydb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_memorydb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.Setup,
		cluster.Setup,
		parametergroup.Setup,
		snapshot.Setup,
		subnetgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
