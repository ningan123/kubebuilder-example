```bash
# kubebuilder init --domain example.com --license apache2 --owner "The Kubernetes authors"
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.14.1
Update dependencies:
$ go mod tidy
Next: define a resource with:
$ kubebuilder create api


# kubebuilder create api --group webapp --version v1 --kind Frigate
Create Resource [y/n]
y
Create Controller [y/n]
y
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
api/v1/frigate_types.go
controllers/frigate_controller.go
Update dependencies:
$ go mod tidy
Running make:
$ make generate
mkdir -p /root/kubebuilder-example/bin
test -s /root/kubebuilder-example/bin/controller-gen && /root/kubebuilder-example/bin/controller-gen --version | grep -q v0.11.1 || \
GOBIN=/root/kubebuilder-example/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.11.1
/root/kubebuilder-example/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests


# make install
/root/kubebuilder-example/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
test -s /root/kubebuilder-example/bin/kustomize || { curl -Ss "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash -s -- 3.8.7 /root/kubebuilder-example/bin; }
{Version:kustomize/v3.8.7 GitCommit:ad092cc7a91c07fdf63a2e4b7f13fa588a39af4f BuildDate:2020-11-11T23:14:14Z GoOs:linux GoArch:amd64}
kustomize installed to /root/kubebuilder-example/bin/kustomize
/root/kubebuilder-example/bin/kustomize build config/crd | kubectl apply -f -
customresourcedefinition.apiextensions.k8s.io/frigates.webapp.example.com created
# kubectl get crd |grep frigate
frigates.webapp.example.com                           2023-03-27T07:25:20Z

```