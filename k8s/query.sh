token=""
ip=""
phoneNumber=""

curl -H "DEVOPS-API-TOKEN:$token" http://$ip:8080/api/v1/queryphone?phone=$phoneNumber

curl -H "DEVOPS-API-TOKEN:$token" http://$ip:8080/api/v1/version
