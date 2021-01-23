#!/bin/bash
# test.sh

JUMP_HOST=$(terraform output -raw external_ip_address_vm_nginx)
NGINX_HOST=$JUMP_HOST
POSTGRES_HOST=$(terraform output -raw internal_ip_address_vm_postgres)
APP_1_HOST=$(terraform output -raw internal_ip_address_vm_app_1)
APP_2_HOST=$(terraform output -raw internal_ip_address_vm_app_2)

echo '1. nginx and conf'
echo " ssh hw@$NGINX_HOST"
echo " cat nginx.conf"

echo '2.'

source secrets.env && echo "c. see https://console.cloud.yandex.ru/folders/$YC_FOLDER_ID/compute"
yc compute instances list

echo 'd. curl'
curl http://$NGINX_HOST/healthcheck 2> /dev/null | jq

echo 'e. postgres'
echo " ssh -J hw@$JUMP_HOST hw@$POSTGRES_HOST"
echo " docker-compose exec db psql -U postgres"
docker-compose exec db psql -U postgres -c 'SELECT * FROM healthchecks;'

echo '3. subnets at https://console.cloud.yandex.ru/folders/b1gsird0jdi6k8dibciv/vpc/network/enpgsihufkjr3mp3kqtk'
yc vpc subnet list

echo '4. see deploy at ./deploy.sh'
echo ' code ./deploy.sh'