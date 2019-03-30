# Deployment

## Setup External Dependencies

### OAuth 2.0 Providers

Setup OAuth application on any of the supported providers which you intend to use. Supported providers are GitHub, Google, and LinkedIn. Generate keypairs for each oauth application you setup.

### MongoDB Cluster
Setup a MongoDB cluster on Mongo Atlas. The free M0 tier will be sufficient for development and small deployments. Setup a user for the microservices which has read and write permissions. By default all IP addresses are not allowed to connect to the cluster. You will want to whitelist and IP addresses which you will be manually connecting from. Additionally you should setup VPC peering to your API's VPC or whitelist the IP address of the NAT in your API's VPC. If you are on the M0 or other small tiers then you can not enable VPC peering and will need to whitelist the NAT's IP.

### Sparkpost
Setup an account with Sparkpost. You will also need to write / import all of the templates which the API attempts to send. Generate an Sparkpost API key for the API to use.

### AWS Services

If you are running the API on AWS then you should attach an IAM role to your API containers / services which will give the API access to the required AWS resources. If you are running the API outside of AWS, you should generate an AWS API keypair which can be loaded into the API.

#### S3
Setup an S3 bucket for uploads. The API will store resumes in this bucket and return presigned URLs to API consumers.

#### SNS
Setup an SNS application. You will need to generate certificates and keys for your iOS and Android mobile applications and load then into the SNS application. The API will send push notifications to iOS and Android devices via SNS using the ARN's generated during this setup.


## Setup Cloud Infrastructure
It is highly recommended that you use AWS to host your production deployment. However it should be possible to run the API elsewhere. You will need an S3 compatible storage service (Digital Ocean Spaces is a potential alternative) if you want to store resumes. And you will need AWS SNS, if you want to send push notifications. However the API can be run without either of these services.

### AWS Setup
Create an AWS account. Start by creating a VPC with 4 subnets. 2 subnets should be public and 2 subnets should be private. Deploy a NAT with an Elastic IP into the VPC which will allow microservices to connect to external websites, without being exposed to the public internet. Deploy an Elastic Load Balancer which does SSL termination and forwards traffic to the ECS cluster. You will need to load the SSL certificate into the load balancer.

Deploy a ECS cluster into the private subnets. The security group for this cluster should allow incoming and outgoing tcp connections with the subnet on ports 8000-8050 and on port 80 / 443. Once the ECS cluster is setup, then you can create task definitions for each microservice. The task definitions should use the API container on DockerHub and specify the service name in the command for the container. The `HI_CONFIG` environment variable should point to the location of the configuration json file. Additionally any secret configuration variables should be loaded via environment variables. Start a ECS service on the cluster for each microservice. When setting up the ECS services you will need to create AWS Route 53 DNS routes. You will first need to create an AWS 53 DNS Zone to put these routes into. Lastly when setting up the gateway you will need to attach the load balancer you created to the gateway.

### DNS Setup
You should already have an account with a DNS provider which is used for managing your website. Create an A record to setup a url which will resolve to the IP address of the load balancer for the API.

## Writing the Configuration File
Start with the provided example configuration json file and modify it as needed to satisfy your requirements. The urls for each service will depend on the dns routes that you setup when deploying the API. Any not secret configuration variables in the file should also be setup. For a full list of all required variables you can look at the example development configuration file. The definitions for registration and rsvp and the fields you will most likely need to configure. Any variable can be overwritten by setting an environment variable with the same key. You should use environment variables to load secret configuration variables.
