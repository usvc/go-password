PROJECT_NAME=password
CMD_ROOT=password

-include ./makefile.properties

deps:
	go mod vendor -v
	go mod tidy -v
run:
	go run ./cmd/$(CMD_ROOT)
test:
	go test -json ./... -cover -coverprofile c.out
build:
	go build -o ./bin/$(CMD_ROOT) ./cmd/$(CMD_ROOT)_${GOOS}_${GOARCH}
build_production:
	CGO_ENABLED=0 \
	go build \
		-a -v \
		-ldflags "-X main.Commit=$$(git rev-parse --verify HEAD) \
			-X main.Version=$$(git describe --tag $$(git rev-list --tags --max-count=1)) \
			-X main.Timestamp=$$(date +'%Y%m%d%H%M%S') \
			-extldflags 'static' \
			-s -w" \
		-o ./bin/.$(CMD_ROOT)_$$(go env GOOS)_$$(go env GOARCH)${BIN_EXT} \
		./cmd/$(CMD_ROOT)
	rm -rf ./bin/$(CMD_ROOT)_$$(go env GOOS)_$$(go env GOARCH)${BIN_EXT}
	upx -9 -v \
		./bin/.$(CMD_ROOT)_$$(go env GOOS)_$$(go env GOARCH)${BIN_EXT} \
		-o ./bin/$(CMD_ROOT)_$$(go env GOOS)_$$(go env GOARCH)${BIN_EXT}

package:
	docker build --file ./deploy/Dockerfile --tag $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest .
save:
	mkdir -p ./build
	docker save --output ./build/themigrator.tar.gz $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest
load:
	docker load --input ./build/themigrator.tar.gz
publish_dockerhub:
	docker push $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME):latest

see_ci:
	xdg-open https://gitlab.com/zephinzer/themigrator/pipelines

.ssh:
	mkdir -p ./.ssh
	ssh-keygen -t rsa -b 8192 -f ./.ssh/id_rsa -q -N ""
	cat ./.ssh/id_rsa | base64 -w 0 > ./.ssh/id_rsa.base64
