MIGRATIONS_FOLDER = $(PWD)/migrations

mig.create:
	migrate create -ext sql -dir migrations -seq create_init_tables

mig.up:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_FOLDER) up

mig.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down
