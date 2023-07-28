dockerup:
	docker compose -f docker-compose.local.yml up -d

dockerdown:
	docker compose -f docker-compose.local.yml down

.PHONY: dockerup dockerdown