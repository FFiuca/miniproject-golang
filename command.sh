# migration
migrate create -ext sql -dir migrations -seq create_user #sequence id
migrate create -ext sql -dir migrations create_user # timestamp id
migrate -path . -database "mysql://root:admin@tcp(127.0.0.1:3302)/go1" up


