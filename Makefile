dockerup:
	docker compose -f docker-compose.local.yml up -d

dockerdown:
	docker compose -f docker-compose.local.yml down

migc:
	migrate create -ext sql -dir db/migration -seq $(name)

migup:
	migrate -path db/migration -database "postgresql://$(pguser):$(pgpwd)@localhost:5432/simplebank?sslmode=disable" -verbose up

migd:
	migrate -path db/migration  -database "postgresql://$(pguser):$(pgpwd)@localhost:5432/simplebank?sslmode=disable" -verbose down

sqlcini:
	sqlc init

sqlcg:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: dockerup dockerdown migc migup migd sqlcini sqlcg test