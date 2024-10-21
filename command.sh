# migration
migrate create -ext sql -dir migrations -seq create_user #sequence id
migrate create -ext sql -dir migrations create_user # timestamp id
migrate -path . -database "mysql://root:admin@tcp(127.0.0.1:3302)/go1" up

# jika ada error migrate,perbaiki sql and edit schema_migrations to latest stable version and change dirty to zero the run again

