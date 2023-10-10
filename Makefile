BUILD_ENV := CGO_ENABLED=0
export BUILD_TIME=$(shell date +%Y-%m-%d)
export GIT_COMMIT_ID=$(wordlist 1,9,${DRONE_COMMIT_SHA})
export VERSION=${DRONE_TAG}
export LDFLAGS="\
	-X 'cg/pkg/util.GitCommitId=${GIT_COMMIT_ID}' \
	-X 'cg/pkg/util.BuildTime=${BUILD_TIME}' \
	-X 'cg/pkg/util.Version=${DRONE_TAG}' \
	"

TARGET_EXEC := cg

.PHONY: all setup linux osx

all: setup linux osx finish

setup:
	mkdir -p build

osx:
	${BUILD_ENV} GOARCH=amd64 GOOS=darwin go build -v -a -ldflags ${LDFLAGS} -o build/${TARGET_EXEC}_darwin_amd64
	${BUILD_ENV} GOARCH=arm64 GOOS=darwin go build -v -a -ldflags ${LDFLAGS} -o build/${TARGET_EXEC}_darwin_arm64

linux:
	${BUILD_ENV} GOARCH=amd64 GOOS=linux go build -v -a -ldflags ${LDFLAGS} -o build/${TARGET_EXEC}_linux_amd64

clean:
	rm -rf build

finish:
	cp usage.md build/ && cd build && zip -r cg_${VERSION}.zip ./
