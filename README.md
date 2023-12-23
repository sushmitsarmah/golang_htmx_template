### Generate templates

templ generate --path ./internal/templates && mv ./internal/templates/*.go ./internal/views/

### Generate styles

tailwind -o ./internal/styles/styles.css

### Start Server

go run cmd/main.go