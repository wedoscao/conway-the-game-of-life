.PHONY: build

dev: serve
	@go run cmd/server/main.go

serve: go_serve tailwind_serve

go_serve:
	@go mod tidy

tailwind_serve:
	@npx @tailwindcss/cli -i static/css/tailwind.css -o static/css/output.css -m

watch:
	@npx @tailwindcss/cli -i static/css/tailwind.css -o static/css/output.css -wm & air

migrate:
	@go run cmd/migration/main.go
