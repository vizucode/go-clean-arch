docker compose --file ./docker-compose-dev.yaml down
docker rmi ecommerceapp
docker compose --file ./docker-compose-dev.yaml --env-file .env up -d