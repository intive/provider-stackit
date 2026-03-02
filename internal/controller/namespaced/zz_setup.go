// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bucket "github.com/intive/provider-stackit/internal/controller/namespaced/objectstorage/bucket"
	credential "github.com/intive/provider-stackit/internal/controller/namespaced/objectstorage/credential"
	credentialsgroup "github.com/intive/provider-stackit/internal/controller/namespaced/objectstorage/credentialsgroup"
	credentialopensearch "github.com/intive/provider-stackit/internal/controller/namespaced/opensearch/credential"
	instance "github.com/intive/provider-stackit/internal/controller/namespaced/opensearch/instance"
	database "github.com/intive/provider-stackit/internal/controller/namespaced/postgresflex/database"
	instancepostgresflex "github.com/intive/provider-stackit/internal/controller/namespaced/postgresflex/instance"
	user "github.com/intive/provider-stackit/internal/controller/namespaced/postgresflex/user"
	providerconfig "github.com/intive/provider-stackit/internal/controller/namespaced/providerconfig"
	credentialrabbitmq "github.com/intive/provider-stackit/internal/controller/namespaced/rabbitmq/credential"
	instancerabbitmq "github.com/intive/provider-stackit/internal/controller/namespaced/rabbitmq/instance"
	credentialredis "github.com/intive/provider-stackit/internal/controller/namespaced/redis/credential"
	instanceredis "github.com/intive/provider-stackit/internal/controller/namespaced/redis/instance"
	instancesecretsmanager "github.com/intive/provider-stackit/internal/controller/namespaced/secretsmanager/instance"
	usersecretsmanager "github.com/intive/provider-stackit/internal/controller/namespaced/secretsmanager/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		credential.Setup,
		credentialsgroup.Setup,
		credentialopensearch.Setup,
		instance.Setup,
		database.Setup,
		instancepostgresflex.Setup,
		user.Setup,
		providerconfig.Setup,
		credentialrabbitmq.Setup,
		instancerabbitmq.Setup,
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
		credentialopensearch.SetupGated,
		instance.SetupGated,
		database.SetupGated,
		instancepostgresflex.SetupGated,
		user.SetupGated,
		providerconfig.SetupGated,
		credentialrabbitmq.SetupGated,
		instancerabbitmq.SetupGated,
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
