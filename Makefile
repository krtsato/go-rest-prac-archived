up-all:
	docker-compose up --build -d

down-all:
	docker-compose down --rmi all -v
