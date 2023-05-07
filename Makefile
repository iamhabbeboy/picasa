build:
	go build -o wallpaper

test:
	go test -tags integration -p 1 ./...