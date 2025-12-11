// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bucket "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/bucket"
	credential "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/credential"
	credentialsgroup "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/credentialsgroup"
	database "github.com/intive/provider-stackit/internal/controller/cluster/postgresflex/database"
	instance "github.com/intive/provider-stackit/internal/controller/cluster/postgresflex/instance"
	user "github.com/intive/provider-stackit/internal/controller/cluster/postgresflex/user"
	providerconfig "github.com/intive/provider-stackit/internal/controller/cluster/providerconfig"
	credentialredis "github.com/intive/provider-stackit/internal/controller/cluster/redis/credential"
	instanceredis "github.com/intive/provider-stackit/internal/controller/cluster/redis/instance"
	instancesecretsmanager "github.com/intive/provider-stackit/internal/controller/cluster/secretsmanager/instance"
	usersecretsmanager "github.com/intive/provider-stackit/internal/controller/cluster/secretsmanager/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		credential.Setup,
		credentialsgroup.Setup,
		database.Setup,
		instance.Setup,
		user.Setup,
		providerconfig.Setup,
		credentialredis.Setup,
		instanceredis.Setup,
		instancesecretsmanager.Setup,
		usersecretsmanager.Setup,
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
		bucket.SetupGated,
		credential.SetupGated,
		credentialsgroup.SetupGated,
		database.SetupGated,
		instance.SetupGated,
		user.SetupGated,
		providerconfig.SetupGated,
		credentialredis.SetupGated,
		instanceredis.SetupGated,
		instancesecretsmanager.SetupGated,
		usersecretsmanager.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
