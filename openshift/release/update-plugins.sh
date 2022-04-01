#!/bin/bash

# A standalone script to update embedded plugin dependencies
# It can be combined with environment variables to override default values:
#
# - FAAS_VERSION - specify a version for FaaS plugin 
# - EVENT_VERSION - specify a version for event plugin
#
# Usage: FAAS_VERSION=value update-plugins.sh

ROOT_DIR="$(dirname "${BASH_SOURCE[0]:-$0}")/../.."
# shellcheck source=./common.sh
source "$ROOT_DIR/openshift/release/common.sh"

update_faas_plugin
