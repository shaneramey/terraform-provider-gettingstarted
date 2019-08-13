RELEASE_VERSION=v0.1.0
RELEASE_DIR=releases
GOARCH?=amd64
BINARY=terraform-provider-gettingstarted_${RELEASE_VERSION}_x4
LINUX_RELEASE=terraform-provider-gettingstarted_${RELEASE_VERSION}_linux_${GOARCH}
MAC_RELEASE=terraform-provider-gettingstarted_${RELEASE_VERSION}_darwin_${GOARCH}

default: deps linux mac stage

deps:
	go get -d

clean:
	rm -rf "${RELEASE_DIR}"
	rm ${LINUX_RELEASE} ${MAC_RELEASE} ${LINUX_RELEASE}.zip ${MAC_RELEASE}.zip

linux:
	GOOS=linux GOARCH=${GOARCH} go build -o "${BINARY}"
	zip "${LINUX_RELEASE}.zip" "${BINARY}"
	rm "${BINARY}"

mac:
	GOOS=darwin GOARCH=${GOARCH} go build -o "${BINARY}"
	zip "${MAC_RELEASE}.zip" "${BINARY}"
	rm "${BINARY}"

stage:
	mkdir -p "${RELEASE_DIR}"
	mv "${LINUX_RELEASE}.zip" "${RELEASE_DIR}"
	mv "${MAC_RELEASE}.zip" "${RELEASE_DIR}"
