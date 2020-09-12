dev-up-all:
	docker-compose up --build -d

dev-down-all:
	docker-compose down --rmi all -v
