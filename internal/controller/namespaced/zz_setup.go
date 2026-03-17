// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	credential "github.com/intive/provider-stackit/internal/controller/namespaced/mariadb/credential"
	instance "github.com/intive/provider-stackit/internal/controller/namespaced/mariadb/instance"
	instancemongodbflex "github.com/intive/provider-stackit/internal/controller/namespaced/mongodbflex/instance"
	user "github.com/intive/provider-stackit/internal/controller/namespaced/mongodbflex/user"
	bucket "github.com/intive/provider-stackit/internal/controller/namespaced/objectstorage/bucket"
	credentialobjectstorage "github.com/intive/provider-stackit/internal/controller/namespaced/objectstorage/credential"
	credentialsgroup "github.com/intive/provider-stackit/internal/controller/namespaced/objectstorage/credentialsgroup"
	alertgroup "github.com/intive/provider-stackit/internal/controller/namespaced/observability/alertgroup"
	instanceobservability "github.com/intive/provider-stackit/internal/controller/namespaced/observability/instance"
	logalertgroup "github.com/intive/provider-stackit/internal/controller/namespaced/observability/logalertgroup"
	scrapeconfig "github.com/intive/provider-stackit/internal/controller/namespaced/observability/scrapeconfig"
	credentialopensearch "github.com/intive/provider-stackit/internal/controller/namespaced/opensearch/credential"
	instanceopensearch "github.com/intive/provider-stackit/internal/controller/namespaced/opensearch/instance"
	database "github.com/intive/provider-stackit/internal/controller/namespaced/postgresflex/database"
	instancepostgresflex "github.com/intive/provider-stackit/internal/controller/namespaced/postgresflex/instance"
	userpostgresflex "github.com/intive/provider-stackit/internal/controller/namespaced/postgresflex/user"
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
		credential.Setup,
		instance.Setup,
		instancemongodbflex.Setup,
		user.Setup,
		bucket.Setup,
		credentialobjectstorage.Setup,
		credentialsgroup.Setup,
		alertgroup.Setup,
		instanceobservability.Setup,
		logalertgroup.Setup,
		scrapeconfig.Setup,
		credentialopensearch.Setup,
		instanceopensearch.Setup,
		database.Setup,
		instancepostgresflex.Setup,
		userpostgresflex.Setup,
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
		credential.SetupGated,
		instance.SetupGated,
		instancemongodbflex.SetupGated,
		user.SetupGated,
		bucket.SetupGated,
		credentialobjectstorage.SetupGated,
		credentialsgroup.SetupGated,
		alertgroup.SetupGated,
		instanceobservability.SetupGated,
		logalertgroup.SetupGated,
		scrapeconfig.SetupGated,
		credentialopensearch.SetupGated,
		instanceopensearch.SetupGated,
		database.SetupGated,
		instancepostgresflex.SetupGated,
		userpostgresflex.SetupGated,
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
