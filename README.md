# geode-k8s

Initially utilized Kubebuilder to create a Kind for Locator and Server under the group GeodeCluster.

TODO:
1. Update API for proper configuration values
2. Update controller to have logic for actually doing things:)


# Testing

Install KIND (https://github.com/kubernetes-sigs/kind)
1. Open terminal and run: go get sigs.k8s.io/kind@v0.5.1
2. cd $GOPATH/bin
3. ./kind create cluster
4. kubectl cluster-info

Install CRD's into the cluster:
1. cd geode-k8s
2. make install

Run your controller (this will run in the foreground, so switch to a new terminal if you wnat to leave it runnning):
1. make run

Build and push your image to the location specified by IMG:
1. make docker-build docker-push IMG=<some-registry>/<project-name>:tag
2. make deploy IMG=<some-registry>/<project-name>:tag
