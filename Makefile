build:
	go build -o kubectx-git

test:
	bats e2e-test/test.bats

.PHONY: build, test
