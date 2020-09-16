dev-up-all:
	docker-compose up --build -d

dev-down-all:
	docker-compose down --rmi all -v

dev-prune-sys:
	docker system prune

dev-logs:
	docker-compose logs -f

go-tidy:
	go mod tidy

go-clean:
	go clean -modcache