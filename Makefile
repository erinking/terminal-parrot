
ifndef DOCKER_IMAGE
  DOCKER_IMAGE=partyparrot
endif

SRC:=$(shell find ./ -name \*.go -print)

# Build the individual binaries based on source level dependencies
all: parrot

parrot: $(SRC)
	@echo "Building the parrot"
	@(go build -o parrot $(SRC))

docker: parrot
	docker build -t $(DOCKER_IMAGE) ./

push: docker
	docker push $(DOCKER_IMAGE)
