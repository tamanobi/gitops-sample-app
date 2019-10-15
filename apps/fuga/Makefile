NAME         := fuga-app
REVISION     := $(git rev-parse --short HEAD)
ORIGIN       := $(git remote get-url origin)
TAG          := $(REVISION)
REGISTRY     := your.docker.registry.local
BUILD_FILES  := .build

.PHONY: build
build: $(BUILD_FILES)

.PHONY: test
test: 
	@echo $(REGISTRY)/$(NAME):$(TAG)

.build: Dockerfile $(find . -type f)
	docker build \
        --build-arg GIT_REVISION=$(REVISION) \
        --build-arg GIT_ORIGIN=$(ORIGIN) \
        --build-arg IMAGE_NAME=$(REGISTRY)/$(NAME) \
		-t $(REGISTRY)/$(NAME):$(TAG) .
	docker inspect --format="{{range .RepoTags}}{{.}}{{end}}" $(REGISTRY)/$(NAME):$(TAG) > .build
	if [ -ne `cat .build` ]; then
		exit 1
	fi


.PHONY: clean
clean: $(BUILD_FILES)
	cat $^ | xargs -I {} docker rmi {}
	rm -f $^