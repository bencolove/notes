# Setup Testing `kube` Environments
* minikube
* microk8s
* k3s

---
## `microk8s`
>Features:
* single node env
* requires `snap`
* no VM needed for nodes

>Setup:
1. Create a multipass VM or run it locally
`multipass launch --name mk8svm --mem 4G --disk 5G`
1. Install `microku8s` on the vm
`multipass exec mk8svm -- sudo snap install microk8s --classic`
1. Shell into vm or alias command
    * `multipass shell mk8svm`  
    * `sudo microk8s.kubectl get nodes`  
    * `sudo snap alias microk8s.kubectl kubectl` and 
    * `sudo usermod -aG microk8s ubuntu`
1. Expose multipassed `kubectl` to host
    * `microk8s.config` expose config file for external `kubectl`
    * `multipass exec microk8s -- sudo microk8s.config > microk8s.yaml`
    * `export KUBECONFIG=$PWD/microk8s.yaml`  


>Under the hood:  
`microk8s` mimics the `snap` package(namespace and  isolation) as the single worker node

>Commands:  
* `microk8s.status`
* `sudo snap remove microk8s`

## Simple Testcase for `kuberctl`
1. a  
`kubetcl create deployment nginx --image=nginx`
1. b  
`kubectl get deployments`
1. c  
`kubectl get pods`
1. d  
`kubectl get all --all-namespaces`
