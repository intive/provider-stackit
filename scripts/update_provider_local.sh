#!/bin/bash
set -e

make generate
make build

# Required for testing local terraform provider
# Force deletion of existing package in cluster, otherwise it will
# skip updating the package in cluster when the version string did not change.
make local.xpkg.sync
REVISION=$(kubectl get providerrevisions.pkg.crossplane.io -o name 2>/dev/null | head -1)
if [ -n "$REVISION" ]; then
  kubectl delete "$REVISION"
fi

make local-deploy
