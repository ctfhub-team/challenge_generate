BUILD_ENV := CGO_ENABLED=0
export BUILD_TIME=`date +%s`
export GIT_COMMIT_ID=${DRONE_COMMIT_SHA:0:8}
export VERSION=${DRONE_TAG}
VAR_INJECT=-X util.GitCommitId=${GIT_COMMIT_ID} -X util.BuildTime=${BUILD_TIME} -X util.Version=${DRONE_TAG}
LDFLAGS=-v -a -ldflags '-s -w ${VAR_INJECT}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}"

TARGET_EXEC := cg

.PHONY: all setup build-linux build-osx

all: setup build-linux build-osx 

setup:
	mkdir -p build

build-osx:
	${BUILD_ENV} GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -o build/${TARGET_EXEC}_darwin_amd64
	${BUILD_ENV} GOARCH=arm64 GOOS=darwin go build ${LDFLAGS} -o build/${TARGET_EXEC}_darwin_arm64

build-linux:
	${BUILD_ENV} GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -o build/${TARGET_EXEC}_linux_amd64
	${BUILD_ENV} GOARCH=arm64 GOOS=linux go build ${LDFLAGS} -o build/${TARGET_EXEC}_linux_arm64
