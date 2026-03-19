#!/bin/bash
# Use a local build of the upstream terraform provider

set -e

if [[ -f "$1" ]]; then
  TF_PROVIDER_ZIP_IN="$1"
else
  echo "usage: source $0 path/to/terraform-provider/dist.zip"
  return 1
fi

mkdir -p ".cache"
cp "$TF_PROVIDER_ZIP_IN" ".cache/"
TF_PROVIDER_ZIP="$(basename "$TF_PROVIDER_ZIP_IN")"
OS_ARCH="$(echo "$TF_PROVIDER_ZIP" | sed -E 's/.*_([^_]+_[^_]+)\.zip/\1/')"

PROVIDER_BIN_DIR=".cache/provider"
rm -rf "$PROVIDER_BIN_DIR"
mkdir -p "$PROVIDER_BIN_DIR"
unzip -o ".cache/$TF_PROVIDER_ZIP" -d "$PROVIDER_BIN_DIR"
PROVIDER_BIN_PATH="$(find "$PROVIDER_BIN_DIR" -maxdepth 1 -type f -executable | head -n1)"
PROVIDER_BIN="$(basename "$PROVIDER_BIN_PATH")"

# Use local provider for make generate
mkdir -p ".work"
cat > .work/local_provider.tfrc <<EOF
provider_installation {
  filesystem_mirror {
    path    = "$(realpath .)/$PROVIDER_BIN_DIR"
    include = ["registry.terraform.io/stackitcloud/stackit"]
  }
  direct {
    exclude = ["registry.terraform.io/stackitcloud/stackit"]
  }
}
EOF
export TF_CLI_CONFIG_FILE=".work/local_provider.tfrc"

# Terraform Registry is queried, so we need an existing version
TF_PROVIDER_VERSION="$(grep "export TERRAFORM_PROVIDER_VERSION" Makefile | awk '{print $4}')"
TF_PROVIDER_NAME="$(grep "export TERRAFORM_PROVIDER_DOWNLOAD_NAME" Makefile | awk '{print $4}')"

# Copy to local mirror
TF_PROVIDER_PATH="$PROVIDER_BIN_DIR/registry.terraform.io/stackitcloud/stackit/$TF_PROVIDER_VERSION/$OS_ARCH/"
mkdir -p "$TF_PROVIDER_PATH"
cp "$PROVIDER_BIN_PATH" "$TF_PROVIDER_PATH/"

# Use local provider for make build
DOCKER_BIN_DIR="_output/bin"
mkdir -p "$DOCKER_BIN_DIR"
cp ".cache/$TF_PROVIDER_ZIP" "$DOCKER_BIN_DIR/${TF_PROVIDER_NAME}_${TF_PROVIDER_VERSION}_$OS_ARCH.zip"
export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX="bin"
export TERRAFORM_NATIVE_PROVIDER_BINARY="$PROVIDER_BIN"
