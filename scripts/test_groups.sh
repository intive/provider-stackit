#!/bin/bash
EX="examples/cluster"

EXAMPLES_KMS="$EX/kms/keyring.yaml,$EX/kms/key.yaml,$EX/kms/wrappingkey.yaml"
EXAMPLES_SVCACC="$EX/serviceaccount/serviceaccount.yaml,$EX/serviceaccount/key.yaml"
EXAMPLES_OBJECTSTORAGE="$EX/objectstorage/bucket.yaml,$EX/objectstorage/credentials-group.yaml,$EX/objectstorage/credential.yaml"
EXAMPLES_DNS="$EX/dns/zone.yaml,$EX/dns/recordset.yaml"
EXAMPLES_SECRETSMANAGER="$EX/secretsmanager/instance.yaml,$EX/secretsmanager/user.yaml"
EXAMPLES_LOGS="$EX/logs/instance.yaml,$EX/logs/accesstoken.yaml"
EXAMPLES_OBSERVABILITY="$EX/observability/instance.yaml,$EX/observability/credential.yaml,$EX/observability/alertgroup.yaml,$EX/observability/logalertgroup.yaml,$EX/observability/scrapeconfig.yaml"
EXAMPLES_MONGODBFLEX="$EX/mongodbflex/instance.yaml,$EX/mongodbflex/user.yaml"
EXAMPLES_LOGME="$EX/logme/instance.yaml,$EX/logme/credential.yaml"
EXAMPLES_MARIADB="$EX/mariadb/instance.yaml,$EX/mariadb/credential.yaml"
EXAMPLES_OPENSEARCH="$EX/opensearch/instance.yaml,$EX/opensearch/credential.yaml"
EXAMPLES_REDIS="$EX/redis/instance.yaml,$EX/redis/credential.yaml"
EXAMPLES_RABBITMQ="$EX/rabbitmq/instance.yaml,$EX/rabbitmq/credential.yaml"
EXAMPLES_SQLSERVERFLEX="$EX/sqlserverflex/instance.yaml,$EX/sqlserverflex/user.yaml"
EXAMPLES_POSTGRESFLEX="$EX/postgresflex/instance.yaml,$EX/postgresflex/user.yaml,$EX/postgresflex/database.yaml"
EXAMPLES_NETWORK_AND_COMPUTE="$EX/network/network.yaml,$EX/network/networkinterface.yaml,$EX/compute/keypair.yaml,$EX/compute/affinitygroup.yaml,$EX/compute/volume.yaml,$EX/compute/server.yaml,$EX/compute/volumeattach.yaml,$EX/compute/networkinterfaceattach.yaml,$EX/loadbalancer/loadbalancer.yaml"


case "$1" in
  kms)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_KMS make uptest
    ;;
  serviceaccount)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_SVCACC make uptest
    ;;
  objectstorage)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_OBJECTSTORAGE make uptest
    ;;
  dns)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_DNS make uptest
    ;;
  secretsmanager)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_SECRETSMANAGER make uptest
    ;;
  logs)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_LOGS make uptest
    ;;
  observability)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_OBSERVABILITY make uptest
    ;;
  mongodbflex)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_MONGODBFLEX make uptest
    ;;
  logme)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_LOGME make uptest
    ;;
  mariadb)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_MARIADB make uptest
    ;;
  opensearch)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_OPENSEARCH make uptest
    ;;
  redis)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_REDIS make uptest
    ;;
  rabbitmq)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_RABBITMQ make uptest
    ;;
  sqlserverflex)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_SQLSERVERFLEX make uptest
    ;;
  postgresflex)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_POSTGRESFLEX make uptest
    ;;
  network-and-compute)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_NETWORK_AND_COMPUTE make uptest
    ;;
  all)
    UPTEST_EXAMPLE_LIST=$EXAMPLES_KMS,$EXAMPLES_SVCACC,$EXAMPLES_OBJECTSTORAGE,$EXAMPLES_DNS,$EXAMPLES_SECRETSMANAGER,$EXAMPLES_LOGS,$EXAMPLES_OBSERVABILITY,$EXAMPLES_MONGODBFLEX,$EXAMPLES_LOGME,$EXAMPLES_MARIADB,$EXAMPLES_OPENSEARCH,$EXAMPLES_REDIS,$EXAMPLES_RABBITMQ,$EXAMPLES_SQLSERVERFLEX,$EXAMPLES_POSTGRESFLEX,$EXAMPLES_NETWORK_AND_COMPUTE make uptest
    ;;
  *)
    echo "Usage: $0 <group>"
    echo "Groups: all, kms, serviceaccount, objectstorage, dns, secretsmanager, logs, observability, mongodbflex, logme, mariadb, opensearch, redis, rabbitmq, sqlserverflex, postgresflex, network-and-compute"
    exit 1
    ;;
esac
