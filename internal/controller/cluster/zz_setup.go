// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	clusterbucket "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/clusterbucket"
	clustercredential "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/clustercredential"
	clustercredentialsgroup "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/clustercredentialsgroup"
	clusterdatabase "github.com/intive/provider-stackit/internal/controller/cluster/postgresflex/clusterdatabase"
	clusterinstance "github.com/intive/provider-stackit/internal/controller/cluster/postgresflex/clusterinstance"
	clusteruser "github.com/intive/provider-stackit/internal/controller/cluster/postgresflex/clusteruser"
	providerconfig "github.com/intive/provider-stackit/internal/controller/cluster/providerconfig"
	clustercredentialrabbitmq "github.com/intive/provider-stackit/internal/controller/cluster/rabbitmq/clustercredential"
	clusterinstancerabbitmq "github.com/intive/provider-stackit/internal/controller/cluster/rabbitmq/clusterinstance"
	clustercredentialredis "github.com/intive/provider-stackit/internal/controller/cluster/redis/clustercredential"
	clusterinstanceredis "github.com/intive/provider-stackit/internal/controller/cluster/redis/clusterinstance"
	clusterinstancesecretsmanager "github.com/intive/provider-stackit/internal/controller/cluster/secretsmanager/clusterinstance"
	clusterusersecretsmanager "github.com/intive/provider-stackit/internal/controller/cluster/secretsmanager/clusteruser"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		clusterbucket.Setup,
		clustercredential.Setup,
		clustercredentialsgroup.Setup,
		clusterdatabase.Setup,
		clusterinstance.Setup,
		clusteruser.Setup,
		providerconfig.Setup,
		clustercredentialrabbitmq.Setup,
		clusterinstancerabbitmq.Setup,
		clustercredentialredis.Setup,
		clusterinstanceredis.Setup,
		clusterinstancesecretsmanager.Setup,
		clusterusersecretsmanager.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		clusterbucket.SetupGated,
		clustercredential.SetupGated,
		clustercredentialsgroup.SetupGated,
		clusterdatabase.SetupGated,
		clusterinstance.SetupGated,
		clusteruser.SetupGated,
		providerconfig.SetupGated,
		clustercredentialrabbitmq.SetupGated,
		clusterinstancerabbitmq.SetupGated,
		clustercredentialredis.SetupGated,
		clusterinstanceredis.SetupGated,
		clusterinstancesecretsmanager.SetupGated,
		clusterusersecretsmanager.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
