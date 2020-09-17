dev-up-all:
	docker-compose up --build -d

dev-down-all:
	docker-compose down --rmi all -v

dev-prune-sys:
	docker system prune

dev-ps:
	docker-compose ps

dev-logs:
	docker-compose logs -f

dev-exec-restapi:
	docker-compose exec restapi bash

dev-exec-mariadb:
	docker-compose exec mariadb sh

go-tidy:
	go mod tidy

go-clean:
	go clean -modcache