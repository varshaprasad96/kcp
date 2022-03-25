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
c "Note that it references our APIResourceSchema, more concretely today's version. In the future, we might want to have another version and roll that out to customers."

c "Let's see the APIExport status:"
pe "kubectl get apiexports"

c "Let's also register a physical cluster, to make sure that the APIs we need are available."
pe "kubectl apply -f contrib/demo/clusters/kind/us-east1.yaml"

c "Let's apply the CRDs required by cert-manager:"
pe "kubectl apply -f ${DEMO_DIR}/apibinding/cert-manager-crds.yaml"

# Run cert-manager on a kind cluster. Is this to be exposed to users?
cp ${KUBECONFIG} ${DEMO_DIR}/cm.kubeconfig
sed -i 's/current-context:.*/current-context: system:admin/' cm.kubeconfig
kubectl --kubeconfig=${CLUSTERS_DIR}/us-east1.kubeconfig ${DEMO_DIR}/apibinding/cert-manager-resources/cert-manager.yaml

clear
kubectl config use-context default
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
c "Let's start by creating a private key for ca issuer. This step requires openssl to be installed locally"
pe "openssl genrsa 2048 > ${DEMO_DIR}/ca.key"
pe "openssl req -new -x509 -nodes -days 365000 \
  -subj `"/C=US/ST=New York/L=New York City/O=Kubernetes/OU=KCP/CN=example.com/emailAddress=test@example.com"` \
  -key ${DEMO_DIR}/ca.key \
  -out ${DEMO_DIR}/ca.crt"

c "Now let's start by creating a certificate and a Certificate Authority issuer in the webapp workspace."
pe "kubectl apply -f ${DEMO_DIR}/apibinding/cert-manager-resources/cert-manager-test-resources.yaml"

sleep 2

c "Cert-manager should now be able to reconcile and create secret with the name `ca-cert-tls` as specified in the issuer."
pe "kubectl get secret ca-cert-tls"

c "Also, we can see that the certificate would now be in a ready condition with the secret generated from cert-manager."
pe "kubectl get certificate serving-cert"

