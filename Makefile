GO_PATH ?= $(shell go env GOPATH)

static:
	CGO_ENABLED=0 go build -v -mod=vendor -tags='osusergo,netgo,gojay,static,static_build' -trimpath -gcflags="all=-trimpath=${GO_PATH}" -ldflags='all=-s -w "-extldflags=-fno-PIC -static"' -asmflags="all=-trimpath=${GO_PATH}" -installsuffix='netgo' github.com/zchee/ccictl
