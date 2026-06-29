GIT_HASH := $(shell git rev-parse --short HEAD)
DOCKER_USER := marie20767
IMAGE_NAME := keto-granola

up:
	docker compose up -d

up/build:
	docker compose up --build -d

down:
	docker compose down

down/vol:
	docker compose down -v

docker/local-build:
	DOCKER_BUILDKIT=1 docker buildx build \
	-t $(DOCKER_USER)/$(IMAGE_NAME):local .

# For CI caching
CACHE_FROM ?=
CACHE_TO ?=

docker/ci-build:
	DOCKER_BUILDKIT=1 docker buildx build \
	$(CACHE_FROM) \
	$(CACHE_TO) \
	--platform linux/amd64 \
	-t $(DOCKER_USER)/$(IMAGE_NAME):latest \
	-t $(DOCKER_USER)/$(IMAGE_NAME):$(GIT_HASH) .

docker/push:
	docker push --all-tags $(DOCKER_USER)/$(IMAGE_NAME)

docker/build-and-push:
	DOCKER_BUILDKIT=1 docker buildx build \
	--platform linux/amd64 \
	--push \
	-t $(DOCKER_USER)/$(IMAGE_NAME):latest \
	-t $(DOCKER_USER)/$(IMAGE_NAME):$(GIT_HASH) .