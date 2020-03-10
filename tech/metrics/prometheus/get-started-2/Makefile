#
# build
#

# build
build:
	env GOOS=linux GOARCH=amd64 go build -o bin/test-app .

#
# Docker
#

# network
create-net:
	docker network create prometheus

# using dockre-compose preferable
run-prom-srv:
	docker run -p 9090:9090 --net-alias=prom-srv --name test-prom-srv prom/prometheus

