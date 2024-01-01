### Generate templates

templ generate --path ./internal/templates

### Generate styles

tailwind -i ./internal/templates/global.css -o ./internal/assets/styles/styles.css

### Start Server

go run cmd/main.go