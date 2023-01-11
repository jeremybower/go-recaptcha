.PHONY: test

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
COVERAGE_DIR=${ROOT_DIR}/coverage

test:
	@mkdir -p ${COVERAGE_DIR}
	@go test \
		-tags="integration unit" \
		-race \
		-timeout 120s \
		-cover \
		-covermode=atomic \
		-coverprofile ${COVERAGE_DIR}/coverage.out \
		-count=1 \
		-failfast \
		${ROOT_DIR}
	@${GO_ENV} go tool cover \
		-html=${COVERAGE_DIR}/coverage.out \
		-o ${COVERAGE_DIR}/coverage.html