# yc-healthcheck
Yandex.Cloud sample deployments

To deploy everything to Yandex.Cloud:

1. Create an account and fill secrets in `tf/.env`

```
export YC_TOKEN=
export YC_CLOUD_ID=
export YC_FOLDER_ID=
```

2. Run scripts

```
cd tf

# init yandex provider
source secrets.env && terraform init

# run terraform to prepare infra
source secrets.env && terraform apply

# deploy all applications
./deploy.sh

# run tests
./test.sh

# destroy nat instance (optional)
yc compute instance delete nat-instance
```

3. Local tests

```
./integration_test_app.sh
./integration_test_nginx.sh
```
