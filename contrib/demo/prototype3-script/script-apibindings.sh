#!/usr/bin/env bash

# Copyright 2022 The KCP Authors.
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

DEMO_DIR="$(dirname "${BASH_SOURCE[0]}")"
source "${DEMO_DIR}"/../.setupEnv
source "${DEMO_DIR}"/script-preamble.sh

c "First we take the role of an API service provider. Our running example is cert-manager. We want to offer that to users of the kcp service."
c "We create a workspace for cert-manager to export the APIs from:"
pe "kubectl kcp workspace create cert-manager-service --use"

c "Then we need an APIResourceSchema. That's very similar to a CRD, actually like a snapshot of a CRD at a point in time:"
pe "cat ${DEMO_DIR}/apibinding/certificate-apiresourceschema.yaml"
pe "kubectl apply -f ${DEMO_DIR}/apibinding/certificate-apiresourceschema.yaml"

c "Now we export the API to make it available for users of the kcp service:"
pe "cat ${DEMO_DIR}/apibinding/cert-manager-apiexport.yaml"
pe "kubectl apply -f ${DEMO_DIR}/apibinding/cert-manager-apiexport.yaml"
c "Note that it references our APIResourceSchema, more concretely the March version. In the future, we might want to have an April version and roll that out to customers."

c "Let's see the APIExport status:"
pe "kubectl get apiexports"

clear
kubectl config use-config default
c "Now we swtich roles. We are the customer now, in need of certificate generation for our web application. We are lucky: kcp has a cert-manager service 🎉"
pe "kubectl kcp workspace create webapp --use"

c "We bind to the cert-manager API by creating an APIBinding:"
pe "cat ${DEMO_DIR}/apibinding/cert-manager-apibinding.yaml"
c "Note that we choose the stable version. There might be a cutting-edge APIExport too."
pe "kubectl apply -f ${DEMO_DIR}/apibinding/cert-manager-apibinding.yaml"

c "The kcp system will do the binding now in the background and make the APIs available:"
pe "kubectl get apibinding -o yaml"

sleep 2

pe "kubectl get apibinding -o yaml"
c "We see it is bound and ready. We should now see the Certificate API available 🤞:"
pe "kubectl api-resources"
pe "kubectl get certificates"

c "Note that this is much more than just a CRD: we have bound to the whole cert-manager service. So certificate requests will actually work 🎉"
c "Let's request a serving cert for our web application:"
