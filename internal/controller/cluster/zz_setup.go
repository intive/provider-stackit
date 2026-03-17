// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	credential "github.com/intive/provider-stackit/internal/controller/cluster/mariadb/credential"
	instance "github.com/intive/provider-stackit/internal/controller/cluster/mariadb/instance"
	instancemongodbflex "github.com/intive/provider-stackit/internal/controller/cluster/mongodbflex/instance"
	user "github.com/intive/provider-stackit/internal/controller/cluster/mongodbflex/user"
	bucket "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/bucket"
	credentialobjectstorage "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/credential"
	credentialsgroup "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/credentialsgroup"
	alertgroup "github.com/intive/provider-stackit/internal/controller/cluster/observability/alertgroup"
	instanceobservability "github.com/intive/provider-stackit/internal/controller/cluster/observability/instance"
	logalertgroup "github.com/intive/provider-stackit/internal/controller/cluster/observability/logalertgroup"
	scrapeconfig "github.com/intive/provider-stackit/internal/controller/cluster/observability/scrapeconfig"
	credentialopensearch "github.com/intive/provider-stackit/internal/controller/cluster/opensearch/credential"
	instanceopensearch "github.com/intive/provider-stackit/internal/controller/cluster/opensearch/instance"
	database "github.com/intive/provider-stackit/internal/controller/cluster/postgresflex/database"
	instancepostgresflex "github.com/intive/provider-stackit/internal/controller/cluster/postgresflex/instance"
	userpostgresflex "github.com/intive/provider-stackit/internal/controller/cluster/postgresflex/user"
	providerconfig "github.com/intive/provider-stackit/internal/controller/cluster/providerconfig"
	credentialrabbitmq "github.com/intive/provider-stackit/internal/controller/cluster/rabbitmq/credential"
	instancerabbitmq "github.com/intive/provider-stackit/internal/controller/cluster/rabbitmq/instance"
	credentialredis "github.com/intive/provider-stackit/internal/controller/cluster/redis/credential"
	instanceredis "github.com/intive/provider-stackit/internal/controller/cluster/redis/instance"
	instancesecretsmanager "github.com/intive/provider-stackit/internal/controller/cluster/secretsmanager/instance"
	usersecretsmanager "github.com/intive/provider-stackit/internal/controller/cluster/secretsmanager/user"
	instancesqlserverflex "github.com/intive/provider-stackit/internal/controller/cluster/sqlserverflex/instance"
	usersqlserverflex "github.com/intive/provider-stackit/internal/controller/cluster/sqlserverflex/user"
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
		instancesqlserverflex.Setup,
		usersqlserverflex.Setup,
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
		instancesqlserverflex.SetupGated,
		usersqlserverflex.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
