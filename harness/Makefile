.PHONY: build
build:
	go build -o bin/app main.go

.PHONY: docker-build
docker-build:
	docker build -t redis_harness .


.PHONY: docker-publish
docker-publish:
	docker tag redis_harness:latest mattmacf98/redis_harness:latest
	docker push mattmacf98/redis_harness:latest