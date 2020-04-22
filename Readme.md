# kubernetes-multinode-setup-example

The simple demonstration of creating multinode kuberenetes cluster with kubeadm.

Consist of simple golang server and client tried running on nodes with help of Pod.

There are definition files for pod, deployment and replicaset which can be used present inside __definitionfiles__ folder.

And verious service to enable communication between pods and node present inside __services__ folder


1. Setting up the __Virtual Machine__
	1. Each vm should have linux os 
	2. Each should have at least 2gb ram and 2 cpus assigned
	3. Each machine should have unique hostname and mac id
		If the hostname is same then change the name by following steps.
		1. Edit __/etc/hostname__ file and change the name of your choice.
		2. Edit __/etc/hosts__ file and change the old name and write the name used in above(step 1) step.
		3. Reboot the system to apply changes.
	4. Create a network for the nodes throught the vm (in case of __virtualbox Host Network manager and make sure the untick the DHCP property.__)
	5. Assign static IP to each vm interface created by the network from above(step 4) step. Starting from the next ip address.
		To assign the IP perform the following steps.
		1. Edit __/etc/network/interfaces__ file and add the static IP configuration
	         
		 For eg.<br/>
		 ~~~
	          auto enp0s8
	          iface enp0s8 inet static
	          address 192.168.56.2
	          netmask 255.255.255.0
		  ~~~
	reboot the machine to apply changes.
	
	6. Mark the swappoff flag by executing following command:
		~~~ 
		swapoff -a 
		~~~
	7. Comment the following lines in __/etc/fstab__
		~~~
		#swap was on /dev/sda3 during installation
		UUID=df4456d5-fd42-49cc-ba01-5fbf2fd87441 none  swap  sw          0  0
		~~~

2. Installing Docker. You can refer to the docker official  website: https://docs.docker.com/get-docker/

    1. Update the apt package index and install packages to allow apt to use a repository over HTTPS:
         ~~~
         sudo apt-get update

         sudo apt-get install \
         apt-transport-https \
         ca-certificates \
         curl \
         gnupg-agent \
         software-properties-common
         ~~~
    2. Add Docker’s official GPG key:
         ~~~
         curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
         ~~~    
    3. Use the following command to set up the stable repository. 
         ~~~
         sudo add-apt-repository \
         "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
         $(lsb_release -cs) \
         stable"
         ~~~
    4.   ~~~ 
    	  sudo apt-get update 
         ~~~
    5. To install a specific version of Docker Engine, list the available versions in the repo, then select and install:
        1. List the versions available in your repo:
            ~~~
             apt-cache madison docker-ce
            ~~~
        2. Install a specific version using the version string from the second column.
            ~~~    
             sudo apt-get install docker-ce=<VERSION_STRING> docker-ce-cli=<VERSION_STRING> containerd.io
            ~~~
        eg.  __sudo apt-get install docker-ce=5:19.03.7~3-0~ubuntu-bionic containerd.io__   
	
    6. Verify that Docker Engine is installed correctly by running the hello-world image.
    
    7. If you get error in above(step 6) step then execute following command:
         ~~~ 
          sudo service docker restart 
         ~~~
        and try running the above step again.


3. Insatlling kubeadm, kubectl and kubelet on each node. You can refer to the kubernetes official website for installation steps. https://kubernetes.io/docs/setup
    1. Execute following command on each node.
        ~~~
         cat <<EOF > /etc/sysctl.d/k8s.conf
         net.bridge.bridge-nf-call-ip6tables = 1
         net.bridge.bridge-nf-call-iptables = 1
         EOF
         sysctl --system
        ~~~
    2. Execute following command one by one:
      ~~~ 
       sudo apt-get update && sudo apt-get install -y apt-transport-https curl 
        
       curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
        
       cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
       deb https://apt.kubernetes.io/ kubernetes-xenial main
       EOF
        
       sudo apt-get update
        
       sudo apt-get install -y kubelet kubeadm kubectl    
      ~~~

4.  Cluster creation.
    1. Execute following command with neccesary changes.
      ~~~
       kubeadm init --pod-network-cidr=<network- cidr> --apiserver-advertise-address=<master node ip address>
      ~~~
        eg. __kubeadm init --pod-network-cidr=192.224.0.0/16 --apiserver-advertise-address=192.168.56.2__
    2. After the successful execution of above command there will be a token which is used to join the node to master copy that and keep somewhere to use in later steps.
    3. In same success message you will get following commands execute them as normal user
      ~~~
       mkdir -p $HOME/.kube
       sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
       sudo chown $(id -u):$(id -g) $HOME/.kube/config
      ~~~
    4. After that apply the Cluster network interface of your wish.
        
	For eg. __kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"__

       Execute following command to check wether master is ready or not.
      ~~~
       kubectl get nodes  
      ~~~
       Execute following command to check wether every kube-system pod is running.
      ~~~
       Kubectl get pod --all-namespaces
      ~~~         
    5.  If everything is fine just paste the token copied in step 2 in the other nodes to join it to cluster.

5. Optional. If you wish to access the dashboard for k8s.
    1. Execute following command for dashboard pod:
      ~~~
       kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-rc7/aio/deploy/recommended.yaml
      ~~~
    2. start proxy: 
      ~~~
       kubectl proxy&    
      ~~~
     open http://localhost:8001 link on master to access the dashboard.
     
    3. If want to access the dashbord remotly. Create your own secret and obtain the token:
      ~~~
       kubectl create serviceaccount <account name>
       kubectl create clusterrolebinding dashboard-admin --clusterrole=cluster-admin --serviceaccount=default:<account name>
       kubectl get secret
       kubectl describe secret <secret name>  
     ~~~ 
     copy the secret key and keep with you
    4. create ssh tunnel from a remote host outside of the cluster where you would access dashboard:
      ~~~
       ssh -L 9999:127.0.0.1:8001 -N -f -l <user name> <k8s master host name or ip>
      ~~~
    5. open a browser with the following api:
     http://localhost:9999/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/
  6. If you want to create different namespaces and deploy resources in it perform following steps:
     1. Create namespace you want
     ~~~
      kubectl create namespace <name>
     ~~~
      eg. __kubectl create namespace demo__
            
     2. Create context for the namespace
       ~~~
        kubectl config set-context <context-name> --namespace=<namespace> --user=<user> --cluster=<cluster name>
       ~~~
       eg.  __kubectl config set-context demo --namespace=demo --user=kubernetes-admin --cluster=kubernetes__
       
     3. Switch to the desired namespace context
       ~~~
        kubectl config use-context <context-name>
       ~~~
       eg. __kubectl config use-context demo__

__You can see the snapshots of all resources running on cluster inside snapshots folder.__
