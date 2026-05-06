start-server:
	@export database_path="postgres://postgres:0404@172.25.96.1:5432/postgres" && \
	cd cmd && go run main.go