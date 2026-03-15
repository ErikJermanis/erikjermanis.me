# Run the application locally for development purposes
run: templ css build
	@./bin/erikjermanisme

# Deploy the application to the production server
deploy: templ css buildlinux
	scp ./bin/erikjermanismelinux erik@erikjermanis.me:/home/erik/sites/erikjermanis.me
	scp -r ./public erik@erikjermanis.me:/home/erik/sites/erikjermanis.me

# Helper to build the application for Linux (for deployment)
buildlinux:
	@GOOS=linux GOARCH=amd64 go build -o bin/erikjermanismelinux

# Helper to build the application for local development
build:
	@go build -o bin/erikjermanisme

# Helper to generate the CSS using Tailwind
css:
	@npx tailwindcss -i ./input.css -o ./public/styles.css

# Helper to render templ pages
templ:
	@templ generate

# Command to publish blog posts to the server without redeploying the entire application
publish:
	scp -r ./blogposts/* erik@erikjermanis.me:/home/erik/sites/erikjermanis.me/blogposts