build:
	docker-compose build

up:
	docker-compose up -d
	
down:
	docker-compose down

logs:
	docker logs -f go-api 

rebuild-up:
	make down && make build && make up && make logs