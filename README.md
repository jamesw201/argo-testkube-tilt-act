# ArgoCD, TestKube, Act and Tilt
This repository contains a demo of ArgoCD, TestKube, Act and Tilt.

The idea is to have ArgoCD deploy a simple application with TestKube managing its tests and being called by Github Actions managed by Act locally.

### tasks
- [ ] use Act to build and publish the docker image
- [ ] have ArgoCD deploy the application to the local cluster
- [ ] create a testkube test
- [ ] run the testkube test from Act

## Create a K3D cluster with local registry for testing
`$ brew install k3d`

`$ brew install tilt` or use the latest install script on the tilt.dev site

`$ k3d registry create testkube-registry --port 6000`

`$ k3d cluster create testkube-cluster --registry-use k3d-testkube-registry:6000 -p "8080:8080@loadbalancer"` 

`$ tilt up`

You should now be able to curl the registry catalogue to find your images:   
`$ curl http://k3d-testkube-registry:6000/v2/_catalog`


### Tearing down the local cluster and registry
`$ k3d cluster delete testkube-cluster && k3d registry delete k3d-testkube-registry`


### ArgoCD
Source material: 
https://github.com/kubeshop/testkube-argocd

Get the login password for the `admin` user by running:
`% kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d`

Patch the argocd-repo-server
`% kubectl patch deployments.apps -n argocd argo-cd-argocd-repo-server --type json --patch-file testkube-argocd/customization/patch.yaml`

`% kubectl patch configmaps -n argocd argocd-cm --patch-file testkube-argocd/customization/argocd-plugins.yaml`

