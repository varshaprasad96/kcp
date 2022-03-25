#!/usr/bin/env bash

# Copyright 2021 The KCP Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

if [[ -n "${REUSE_KIND_CLUSTERS:-}" ]]; then
    exit
fi

DEMO_DIR="$(dirname "${BASH_SOURCE[0]}")"
source "${DEMO_DIR}"/../.setupEnv

kind delete clusters us-west1 us-east1 > /dev/null || true

# delete private keys and certs created for cert-manager
rm -rf ${DEMO_DIR}/ca.key
rm -rf ${DEMO_DIR}/ca.crt

# delete the temprory kubeconfig created to deploy cert-manager
rm -rf ${DEMO_DIR}/cm.kubeconfig
