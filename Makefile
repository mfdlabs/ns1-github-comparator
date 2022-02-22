IMAGE_NAME?=ns1-github-comparator
IMAGE_TAG=$(shell git rev-parse HEAD | cut -c1-7)

build:
	mkdir -p .ssh
	cp /opt/ssh/ops-user ./.ssh/ops-user
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) -f Dockerfile .
ifdef CI
	docker push -t $(IMAGE_NAME):$(IMAGE_TAG)
ifeq "$(CIRCLE_BRANCH)" "master"
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):latest
	docker push $(IMAGE_NAME):latest
endif
endif
.PHONY: build
