// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	folderroleassignment "github.com/intive/provider-stackit/internal/controller/cluster/authorization/folderroleassignment"
	organizationroleassignment "github.com/intive/provider-stackit/internal/controller/cluster/authorization/organizationroleassignment"
	projectcustomrole "github.com/intive/provider-stackit/internal/controller/cluster/authorization/projectcustomrole"
	projectroleassignment "github.com/intive/provider-stackit/internal/controller/cluster/authorization/projectroleassignment"
	customdomain "github.com/intive/provider-stackit/internal/controller/cluster/cdn/customdomain"
	distribution "github.com/intive/provider-stackit/internal/controller/cluster/cdn/distribution"
	backupschedule "github.com/intive/provider-stackit/internal/controller/cluster/compute/backupschedule"
	group "github.com/intive/provider-stackit/internal/controller/cluster/compute/group"
	image "github.com/intive/provider-stackit/internal/controller/cluster/compute/image"
	pair "github.com/intive/provider-stackit/internal/controller/cluster/compute/pair"
	server "github.com/intive/provider-stackit/internal/controller/cluster/compute/server"
	serviceaccountattach "github.com/intive/provider-stackit/internal/controller/cluster/compute/serviceaccountattach"
	updateschedule "github.com/intive/provider-stackit/internal/controller/cluster/compute/updateschedule"
	volume "github.com/intive/provider-stackit/internal/controller/cluster/compute/volume"
	volumeattach "github.com/intive/provider-stackit/internal/controller/cluster/compute/volumeattach"
	recordset "github.com/intive/provider-stackit/internal/controller/cluster/dns/recordset"
	zone "github.com/intive/provider-stackit/internal/controller/cluster/dns/zone"
	instance "github.com/intive/provider-stackit/internal/controller/cluster/edgecloud/instance"
	kubeconfig "github.com/intive/provider-stackit/internal/controller/cluster/edgecloud/kubeconfig"
	token "github.com/intive/provider-stackit/internal/controller/cluster/edgecloud/token"
	git "github.com/intive/provider-stackit/internal/controller/cluster/git/git"
	key "github.com/intive/provider-stackit/internal/controller/cluster/kms/key"
	keyring "github.com/intive/provider-stackit/internal/controller/cluster/kms/keyring"
	wrappingkey "github.com/intive/provider-stackit/internal/controller/cluster/kms/wrappingkey"
	loadbalancer "github.com/intive/provider-stackit/internal/controller/cluster/loadbalancer/loadbalancer"
	observabilitycredential "github.com/intive/provider-stackit/internal/controller/cluster/loadbalancer/observabilitycredential"
	credential "github.com/intive/provider-stackit/internal/controller/cluster/logme/credential"
	instancelogme "github.com/intive/provider-stackit/internal/controller/cluster/logme/instance"
	accesstoken "github.com/intive/provider-stackit/internal/controller/cluster/logs/accesstoken"
	instancelogs "github.com/intive/provider-stackit/internal/controller/cluster/logs/instance"
	credentialmariadb "github.com/intive/provider-stackit/internal/controller/cluster/mariadb/credential"
	instancemariadb "github.com/intive/provider-stackit/internal/controller/cluster/mariadb/instance"
	tokenmodelserving "github.com/intive/provider-stackit/internal/controller/cluster/modelserving/token"
	instancemongodbflex "github.com/intive/provider-stackit/internal/controller/cluster/mongodbflex/instance"
	user "github.com/intive/provider-stackit/internal/controller/cluster/mongodbflex/user"
	area "github.com/intive/provider-stackit/internal/controller/cluster/network/area"
	arearegion "github.com/intive/provider-stackit/internal/controller/cluster/network/arearegion"
	arearoute "github.com/intive/provider-stackit/internal/controller/cluster/network/arearoute"
	groupnetwork "github.com/intive/provider-stackit/internal/controller/cluster/network/group"
	grouprule "github.com/intive/provider-stackit/internal/controller/cluster/network/grouprule"
	ip "github.com/intive/provider-stackit/internal/controller/cluster/network/ip"
	ipassociate "github.com/intive/provider-stackit/internal/controller/cluster/network/ipassociate"
	network "github.com/intive/provider-stackit/internal/controller/cluster/network/network"
	networkinterface "github.com/intive/provider-stackit/internal/controller/cluster/network/networkinterface"
	networkinterfaceattach "github.com/intive/provider-stackit/internal/controller/cluster/network/networkinterfaceattach"
	table "github.com/intive/provider-stackit/internal/controller/cluster/network/table"
	tableroute "github.com/intive/provider-stackit/internal/controller/cluster/network/tableroute"
	bucket "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/bucket"
	credentialobjectstorage "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/credential"
	credentialsgroup "github.com/intive/provider-stackit/internal/controller/cluster/objectstorage/credentialsgroup"
	alertgroup "github.com/intive/provider-stackit/internal/controller/cluster/observability/alertgroup"
	credentialobservability "github.com/intive/provider-stackit/internal/controller/cluster/observability/credential"
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
	folder "github.com/intive/provider-stackit/internal/controller/cluster/resourcemanager/folder"
	project "github.com/intive/provider-stackit/internal/controller/cluster/resourcemanager/project"
	organization "github.com/intive/provider-stackit/internal/controller/cluster/scf/organization"
	organizationmanager "github.com/intive/provider-stackit/internal/controller/cluster/scf/organizationmanager"
	instancesecretsmanager "github.com/intive/provider-stackit/internal/controller/cluster/secretsmanager/instance"
	usersecretsmanager "github.com/intive/provider-stackit/internal/controller/cluster/secretsmanager/user"
	account "github.com/intive/provider-stackit/internal/controller/cluster/serviceaccount/account"
	accountaccesstoken "github.com/intive/provider-stackit/internal/controller/cluster/serviceaccount/accountaccesstoken"
	accountkey "github.com/intive/provider-stackit/internal/controller/cluster/serviceaccount/accountkey"
	exportpolicy "github.com/intive/provider-stackit/internal/controller/cluster/sfs/exportpolicy"
	resourcepool "github.com/intive/provider-stackit/internal/controller/cluster/sfs/resourcepool"
	share "github.com/intive/provider-stackit/internal/controller/cluster/sfs/share"
	cluster "github.com/intive/provider-stackit/internal/controller/cluster/ske/cluster"
	kubeconfigske "github.com/intive/provider-stackit/internal/controller/cluster/ske/kubeconfig"
	instancesqlserverflex "github.com/intive/provider-stackit/internal/controller/cluster/sqlserverflex/instance"
	usersqlserverflex "github.com/intive/provider-stackit/internal/controller/cluster/sqlserverflex/user"
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
