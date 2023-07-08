.PHONY: migrate-up migrate-down

migrate-up:
	./script/migrate_up.sh

migrate-down:
	./script/migrate_up.sh down
