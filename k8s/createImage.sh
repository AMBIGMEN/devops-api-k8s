docker pull centos:latest
docker build -t devops-api:v1.0.0 .
docker save devops-api:v1.0.0 -o devops-api-v1.0.0.tar
