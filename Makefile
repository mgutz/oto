.PHONY: release test

release:
	goreleaser --rm-dist

test:
	go test github.com/pacedotdev/oto/...

