run: templ css build
	@./bin/erikjermanisme

deploy: templ css buildlinux
	scp ./bin/erikjermanismelinux erik@erikjermanis.me:/home/erik/erikjermanisme
	scp ./public/styles.css erik@erikjermanis.me:/home/erik/erikjermanisme/public
	scp ./public/favicon.png erik@erikjermanis.me:/home/erik/erikjermanisme/public

buildlinux:
	@GOOS=linux GOARCH=amd64 go build -o bin/erikjermanismelinux

build:
	@go build -o bin/erikjermanisme

css:
	@npx tailwindcss -i ./input.css -o ./public/styles.css

templ:
	@templ generate