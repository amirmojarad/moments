docker run -p 5432:5432 -d \
    -e POSTGRES_PASSWORD=postgres_test \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_DB=moments_test \
	-v pgdata:/var/lib/postgres/data \
	postgres