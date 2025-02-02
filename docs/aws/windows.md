# Overview

This actions will handle provision Windows Server machines on dedicated hosts. This is a requisite to run nested virtualization on AWS.

Due to how qenvs checks the healthy state for the machine and due to some specific characteristics this action is intended for using within a [custom ami](https://github.com/adrianriobo/qenvs-builder). 

Some of the customizations this image includes:

* create user with admin privileges
* setup autologin for the user
* sshd enabled
* setup auth based on private key
* enable hyper-v
* setup specific UAC levels to allow running privileged without prompt

## Ami replication

Also the action is expecting the image exists with the name: `Windows_Server-2019-English-Full-HyperV-RHQE` at least on one region. If `--spot` option is enable and the image is not offered / created on the chosen region it will copy the AMI as part of the stack (As so it will delete it on destroy).

This process (replicate the ami) increase the overall time for spinning the machine, and can be avoided by running the replication cmd on the image to pre replicate the image on all regions.

## Create

```bash
qenvs aws windows create -h
create

Usage:
  qenvs aws windows create [flags]

Flags:
      --airgap                       if this flag is set the host will be created as airgap machine. Access will done through a bastion
      --ami-name string              name for the custom ami to be used within windows machine. Check README on how to build it (default "Windows_Server-2019-English-Full-HyperV-RHQE")
      --ami-owner string             alias name for the owner of the custom AMI (default "self")
      --ami-username string          name for de default user on the custom AMI (default "ec2-user")
      --conn-details-output string   path to export host connection information (host, username and privateKey)
  -h, --help                         help for create
      --spot                         if this flag is set the host will be created only on the region set by the AWS Env (AWS_DEFAULT_REGION)
      --tags stringToString          tags to add on each resource (--tags name1=value1,name2=value2) (default [])

Global Flags:
      --backed-url string     backed for stack state. Can be a local path with format file:///path/subpath or s3 s3://existing-bucket
      --project-name string   project name to identify the instance of the stack
```

### Outputs

* It will crete an instance and will give as result several files located at path defined by `--conn-details-output`:

  * **host**: host for the windows machine (lb if spot)
  * **username**: username to connect to the machine
  * **id_rsa**: private key to connect to machine
  * **bastion_host**: host for the bastion (airgap)
  * **bastion_username**: username to connect to the bastion (airgap)
  * **bastion_id_rsa**: private key to connect to the bastion (airgap)

* Also it will create a state folder holding the state for the created resources at azure, the path for this folder is defined within `--backed-url`, the content from that folder it is required with the same project name (`--project-name`) in order to detroy the resources.

### Container

When running the container image it is required to pass the authetication information as variables(to setup AWS credentials there is a [helper script](./../../hacks/aws_setup.sh)), following a sample snipped on how to create an instance with default values:  

```bash
podman run -d --name qenvs-rhel \
        -v ${PWD}:/workspace:z \
        -e AWS_ACCESS_KEY_ID=XXX \
        -e AWS_SECRET_ACCESS_KEY=XXX \
        -e AWS_DEFAULT_REGION=us-east-1 \
        quay.io/rhqp/qenvs:0.0.6-dev aws windows create \
            --project-name qenvs-windows \
            --backed-url file:///workspace \
            --conn-details-output /workspace
```