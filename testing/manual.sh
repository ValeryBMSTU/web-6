
set -eux
set -o pipefail

SERVERPORT=3333
SERVERADDR=localhost:${SERVERPORT}

#1 тест
#curl -iL -w "\n" -X GET ${SERVERADDR}/get/

#2 тест
#curl -iL -w "\n" -X GET ${SERVERADDR}/api/user?name=Golang

#3 тест
curl -iL -w "\n" -X POST -H "Content-Type: application/json" --data '{"count":"3KM:12MIN"}' ${SERVERADDR}/
curl -iL -w "\n" -X POST -H "Content-Type: application/json" --data '{"count":"1488"}' ${SERVERADDR}/
curl -iL -w "\n" -X POST -H "Content-Type: application/json" --data '{"count":"228"}' ${SERVERADDR}/
curl -iL -w "\n" -X GET ${SERVERADDR}/
