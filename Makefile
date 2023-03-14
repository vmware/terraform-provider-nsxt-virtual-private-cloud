TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=terraform.local
NAMESPACE=local
NAME=for-vmware-nsxt-virtual-private-cloud
BINARY=terraform-provider-${NAME}
VERSION=1.0.0
OS_ARCH=darwin_arm64

default: install

build:
  go build -o ${BINARY}

release:
  GOOS=darwin GOARCH=arm64 go build -o ./bin/${BINARY}_${VERSION}_darwin_arm64
  GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
  GOOS=freebsd GOARCH=arm64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm64
  GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
  GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
  GOOS=linux GOARCH=arm64 go build -o ./bin/${BINARY}_${VERSION}_linux_arm64
  GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
  GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
  GOOS=openbsd GOARCH=arm64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_arm64
  GOOS=solaris GOARCH=arm64 go build -o ./bin/${BINARY}_${VERSION}_solaris_arm64
  GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
  GOOS=windows GOARCH=arm64 go build -o ./bin/${BINARY}_${VERSION}_windows_arm64

install: build
  mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
  mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
  go test -i $(TEST) || exit 1
  echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
  TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
