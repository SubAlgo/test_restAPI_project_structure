docker run -d \
--name=db_for_test \
--restart=always -p 8100:5432 \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=test123456 \
-e POSTGRES_DB=test_api \
postgres:13.1