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

```