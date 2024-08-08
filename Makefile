db_up:
	@goose --dir ~/projects/go_cart/sql/schema postgres postgres://yash:password@localhost:5432/go_cart up

dev:
	@air --build.cmd "go build -o bin/api cmd/go_cart/main.go" --build.bin "./bin/api" --build.exclude_dir "templates,build"
