dev-up-all:
	docker-compose up --build -d

dev-down-all:
	docker-compose down --rmi all -v

dev-prune-sys:
	docker system prune