// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	folderroleassignment "github.com/intive/provider-stackit/internal/controller/namespaced/authorization/folderroleassignment"
	organizationroleassignment "github.com/intive/provider-stackit/internal/controller/namespaced/authorization/organizationroleassignment"
	projectcustomrole "github.com/intive/provider-stackit/internal/controller/namespaced/authorization/projectcustomrole"
	projectroleassignment "github.com/intive/provider-stackit/internal/controller/namespaced/authorization/projectroleassignment"
	customdomain "github.com/intive/provider-stackit/internal/controller/namespaced/cdn/customdomain"
	distribution "github.com/intive/provider-stackit/internal/controller/namespaced/cdn/distribution"
	backupschedule "github.com/intive/provider-stackit/internal/controller/namespaced/compute/backupschedule"
	group "github.com/intive/provider-stackit/internal/controller/namespaced/compute/group"
	image "github.com/intive/provider-stackit/internal/controller/namespaced/compute/image"
	pair "github.com/intive/provider-stackit/internal/controller/namespaced/compute/pair"
	server "github.com/intive/provider-stackit/internal/controller/namespaced/compute/server"
	serviceaccountattach "github.com/intive/provider-stackit/internal/controller/namespaced/compute/serviceaccountattach"
	updateschedule "github.com/intive/provider-stackit/internal/controller/namespaced/compute/updateschedule"
	volume "github.com/intive/provider-stackit/internal/controller/namespaced/compute/volume"
	volumeattach "github.com/intive/provider-stackit/internal/controller/namespaced/compute/volumeattach"
	recordset "github.com/intive/provider-stackit/internal/controller/namespaced/dns/recordset"
	zone "github.com/intive/provider-stackit/internal/controller/namespaced/dns/zone"
	instance "github.com/intive/provider-stackit/internal/controller/namespaced/edgecloud/instance"
	kubeconfig "github.com/intive/provider-stackit/internal/controller/namespaced/edgecloud/kubeconfig"
	token "github.com/intive/provider-stackit/internal/controller/namespaced/edgecloud/token"
	git "github.com/intive/provider-stackit/internal/controller/namespaced/git/git"
	key "github.com/intive/provider-stackit/internal/controller/namespaced/kms/key"
	keyring "github.com/intive/provider-stackit/internal/controller/namespaced/kms/keyring"
	wrappingkey "github.com/intive/provider-stackit/internal/controller/namespaced/kms/wrappingkey"
	loadbalancer "github.com/intive/provider-stackit/internal/controller/namespaced/loadbalancer/loadbalancer"
	observabilitycredential "github.com/intive/provider-stackit/internal/controller/namespaced/loadbalancer/observabilitycredential"
	credential "github.com/intive/provider-stackit/internal/controller/namespaced/logme/credential"
	instancelogme "github.com/intive/provider-stackit/internal/controller/namespaced/logme/instance"
	accesstoken "github.com/intive/provider-stackit/internal/controller/namespaced/logs/accesstoken"
	instancelogs "github.com/intive/provider-stackit/internal/controller/namespaced/logs/instance"
	credentialmariadb "github.com/intive/provider-stackit/internal/controller/namespaced/mariadb/credential"
	instancemariadb "github.com/intive/provider-stackit/internal/controller/namespaced/mariadb/instance"
	tokenmodelserving "github.com/intive/provider-stackit/internal/controller/namespaced/modelserving/token"
	instancemongodbflex "github.com/intive/provider-stackit/internal/controller/namespaced/mongodbflex/instance"
	user "github.com/intive/provider-stackit/internal/controller/namespaced/mongodbflex/user"
	area "github.com/intive/provider-stackit/internal/controller/namespaced/network/area"
	arearegion "github.com/intive/provider-stackit/internal/controller/namespaced/network/arearegion"
	arearoute "github.com/intive/provider-stackit/internal/controller/namespaced/network/arearoute"
	groupnetwork "github.com/intive/provider-stackit/internal/controller/namespaced/network/group"
	grouprule "github.com/intive/provider-stackit/internal/controller/namespaced/network/grouprule"
	ip "github.com/intive/provider-stackit/internal/controller/namespaced/network/ip"
	ipassociate "github.com/intive/provider-stackit/internal/controller/namespaced/network/ipassociate"
	network "github.com/intive/provider-stackit/internal/controller/namespaced/network/network"
	networkinterface "github.com/intive/provider-stackit/internal/controller/namespaced/network/networkinterface"
	networkinterfaceattach "github.com/intive/provider-stackit/internal/controller/namespaced/network/networkinterfaceattach"
	table "github.com/intive/provider-stackit/internal/controller/namespaced/network/table"
	tableroute "github.com/intive/provider-stackit/internal/controller/namespaced/network/tableroute"
	bucket "github.com/intive/provider-stackit/internal/controller/namespaced/objectstorage/bucket"
	credentialobjectstorage "github.com/intive/provider-stackit/internal/controller/namespaced/objectstorage/credential"
	credentialsgroup "github.com/intive/provider-stackit/internal/controller/namespaced/objectstorage/credentialsgroup"
	alertgroup "github.com/intive/provider-stackit/internal/controller/namespaced/observability/alertgroup"
	credentialobservability "github.com/intive/provider-stackit/internal/controller/namespaced/observability/credential"
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
	folder "github.com/intive/provider-stackit/internal/controller/namespaced/resourcemanager/folder"
	project "github.com/intive/provider-stackit/internal/controller/namespaced/resourcemanager/project"
	organization "github.com/intive/provider-stackit/internal/controller/namespaced/scf/organization"
	organizationmanager "github.com/intive/provider-stackit/internal/controller/namespaced/scf/organizationmanager"
	instancesecretsmanager "github.com/intive/provider-stackit/internal/controller/namespaced/secretsmanager/instance"
	usersecretsmanager "github.com/intive/provider-stackit/internal/controller/namespaced/secretsmanager/user"
	account "github.com/intive/provider-stackit/internal/controller/namespaced/serviceaccount/account"
	accountaccesstoken "github.com/intive/provider-stackit/internal/controller/namespaced/serviceaccount/accountaccesstoken"
	accountkey "github.com/intive/provider-stackit/internal/controller/namespaced/serviceaccount/accountkey"
	exportpolicy "github.com/intive/provider-stackit/internal/controller/namespaced/sfs/exportpolicy"
	resourcepool "github.com/intive/provider-stackit/internal/controller/namespaced/sfs/resourcepool"
	share "github.com/intive/provider-stackit/internal/controller/namespaced/sfs/share"
	cluster "github.com/intive/provider-stackit/internal/controller/namespaced/ske/cluster"
	kubeconfigske "github.com/intive/provider-stackit/internal/controller/namespaced/ske/kubeconfig"
	instancesqlserverflex "github.com/intive/provider-stackit/internal/controller/namespaced/sqlserverflex/instance"
	usersqlserverflex "github.com/intive/provider-stackit/internal/controller/namespaced/sqlserverflex/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		folderroleassignment.Setup,
		organizationroleassignment.Setup,
		projectcustomrole.Setup,
		projectroleassignment.Setup,
		customdomain.Setup,
		distribution.Setup,
		backupschedule.Setup,
		group.Setup,
		image.Setup,
		pair.Setup,
		server.Setup,
		serviceaccountattach.Setup,
		updateschedule.Setup,
		volume.Setup,
		volumeattach.Setup,
		recordset.Setup,
		zone.Setup,
		instance.Setup,
		kubeconfig.Setup,
		token.Setup,
		git.Setup,
		key.Setup,
		keyring.Setup,
		wrappingkey.Setup,
		loadbalancer.Setup,
		observabilitycredential.Setup,
		credential.Setup,
		instancelogme.Setup,
		accesstoken.Setup,
		instancelogs.Setup,
		credentialmariadb.Setup,
		instancemariadb.Setup,
		tokenmodelserving.Setup,
		instancemongodbflex.Setup,
		user.Setup,
		area.Setup,
		arearegion.Setup,
		arearoute.Setup,
		groupnetwork.Setup,
		grouprule.Setup,
		ip.Setup,
		ipassociate.Setup,
		network.Setup,
		networkinterface.Setup,
		networkinterfaceattach.Setup,
		table.Setup,
		tableroute.Setup,
		bucket.Setup,
		credentialobjectstorage.Setup,
		credentialsgroup.Setup,
		alertgroup.Setup,
		credentialobservability.Setup,
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
		folder.Setup,
		project.Setup,
		organization.Setup,
		organizationmanager.Setup,
		instancesecretsmanager.Setup,
		usersecretsmanager.Setup,
		account.Setup,
		accountaccesstoken.Setup,
		accountkey.Setup,
		exportpolicy.Setup,
		resourcepool.Setup,
		share.Setup,
		cluster.Setup,
		kubeconfigske.Setup,
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
		folderroleassignment.SetupGated,
		organizationroleassignment.SetupGated,
		projectcustomrole.SetupGated,
		projectroleassignment.SetupGated,
		customdomain.SetupGated,
		distribution.SetupGated,
		backupschedule.SetupGated,
		group.SetupGated,
		image.SetupGated,
		pair.SetupGated,
		server.SetupGated,
		serviceaccountattach.SetupGated,
		updateschedule.SetupGated,
		volume.SetupGated,
		volumeattach.SetupGated,
		recordset.SetupGated,
		zone.SetupGated,
		instance.SetupGated,
		kubeconfig.SetupGated,
		token.SetupGated,
		git.SetupGated,
		key.SetupGated,
		keyring.SetupGated,
		wrappingkey.SetupGated,
		loadbalancer.SetupGated,
		observabilitycredential.SetupGated,
		credential.SetupGated,
		instancelogme.SetupGated,
		accesstoken.SetupGated,
		instancelogs.SetupGated,
		credentialmariadb.SetupGated,
		instancemariadb.SetupGated,
		tokenmodelserving.SetupGated,
		instancemongodbflex.SetupGated,
		user.SetupGated,
		area.SetupGated,
		arearegion.SetupGated,
		arearoute.SetupGated,
		groupnetwork.SetupGated,
		grouprule.SetupGated,
		ip.SetupGated,
		ipassociate.SetupGated,
		network.SetupGated,
		networkinterface.SetupGated,
		networkinterfaceattach.SetupGated,
		table.SetupGated,
		tableroute.SetupGated,
		bucket.SetupGated,
		credentialobjectstorage.SetupGated,
		credentialsgroup.SetupGated,
		alertgroup.SetupGated,
		credentialobservability.SetupGated,
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
		folder.SetupGated,
		project.SetupGated,
		organization.SetupGated,
		organizationmanager.SetupGated,
		instancesecretsmanager.SetupGated,
		usersecretsmanager.SetupGated,
		account.SetupGated,
		accountaccesstoken.SetupGated,
		accountkey.SetupGated,
		exportpolicy.SetupGated,
		resourcepool.SetupGated,
		share.SetupGated,
		cluster.SetupGated,
		kubeconfigske.SetupGated,
		instancesqlserverflex.SetupGated,
		usersqlserverflex.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
