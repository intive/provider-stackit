// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bucket "github.com/intive/provider-stackit/internal/controller/namespaced/objectstorage/bucket"
	database "github.com/intive/provider-stackit/internal/controller/namespaced/postgresflex/database"
	instance "github.com/intive/provider-stackit/internal/controller/namespaced/postgresflex/instance"
	user "github.com/intive/provider-stackit/internal/controller/namespaced/postgresflex/user"
	providerconfig "github.com/intive/provider-stackit/internal/controller/namespaced/providerconfig"
	credential "github.com/intive/provider-stackit/internal/controller/namespaced/redis/credential"
	instanceredis "github.com/intive/provider-stackit/internal/controller/namespaced/redis/instance"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		database.Setup,
		instance.Setup,
		user.Setup,
		providerconfig.Setup,
		credential.Setup,
		instanceredis.Setup,
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
		database.SetupGated,
		instance.SetupGated,
		user.SetupGated,
		providerconfig.SetupGated,
		credential.SetupGated,
		instanceredis.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
