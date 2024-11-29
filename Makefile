# Variables
DOCKER_BUF := docker run --volume "$$(pwd):/workspace" --workdir /workspace bufbuild/buf

gen-java:
	mvn -T 1C clean install -Dmaven.test.skip -DskipTests

gen-go:
	$(DOCKER_BUF) build && $(DOCKER_BUF) generate

lint:
	$(DOCKER_BUF) lint

update:
	$(DOCKER_BUF) dep update