#!/bin/bash
# deploy.sh

SECRET_DB_PASSWORD=trololo123

JUMP_HOST=$(terraform output -raw external_ip_address_vm_nginx)
NGINX_HOST=$JUMP_HOST
POSTGRES_HOST=$(terraform output -raw internal_ip_address_vm_postgres)
APP_1_HOST=$(terraform output -raw internal_ip_address_vm_app_1)
APP_2_HOST=$(terraform output -raw internal_ip_address_vm_app_2)

COMPOSE_POSTGRES=$(cat ../deploy/postgres.yml)
COMPOSE_APP=$(cat ../deploy/app.yml)
COMPOSE_NGINX=$(cat ../deploy/nginx.yml)

# postgres
ssh -J hw@$JUMP_HOST hw@$POSTGRES_HOST << EOF
    echo '$COMPOSE_POSTGRES' > docker-compose.yml

cat > .env << ABC
DB_PASSWORD=$SECRET_DB_PASSWORD
ABC

    docker-compose up -d
EOF

# app-1
ssh -J hw@$JUMP_HOST hw@$APP_1_HOST << EOF
    echo '$COMPOSE_APP' > docker-compose.yml

cat > .env << ABC
MY_IP=$APP_1_HOST
DB_HOST=$POSTGRES_HOST
DB_PASSWORD=$SECRET_DB_PASSWORD
ABC

    docker-compose up -d
EOF

# app-2
ssh -J hw@$JUMP_HOST hw@$APP_2_HOST << EOF
    echo '$COMPOSE_APP' > docker-compose.yml

cat > .env << ABC
MY_IP=$APP_2_HOST
DB_HOST=$POSTGRES_HOST
DB_PASSWORD=$SECRET_DB_PASSWORD
ABC

    docker-compose up -d
EOF

# nginx
ssh hw@$NGINX_HOST << EOF
    echo '$COMPOSE_NGINX' > docker-compose.yml

cat > nginx.conf << final-line
events {}
http {
    upstream balancer {
        server $APP_1_HOST:3000;
        server $APP_2_HOST:3000;
    }

    server {
        listen 80;

        location / {
            proxy_pass http://balancer;
        }
    }
}

final-line

    docker-compose up -d
EOF

