build:
	# build all components
	go get -t ./
	go build

run:
	# run the user-api component
	go build
	./copier-panic-repro
