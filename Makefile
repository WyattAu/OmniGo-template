# Thin wrapper over scripts/ — the same verbs in every Omni template.
.PHONY: bench bench-update build test lint fmt fmt-check coverage vuln fuzz contract ci clean

build:
	go build ./...

test:
	./scripts/test.sh

lint:
	./scripts/lint.sh

fmt:
	gofumpt -w .

fmt-check:
	@test -z "$$(gofumpt -l .)" || { echo 'gofumpt needed:'; gofumpt -l .; exit 1; }

coverage:
	./scripts/coverage.sh

vuln:
	./scripts/vuln.sh

fuzz:
	./scripts/fuzz.sh

contract:
	./scripts/check-contract.sh

## What CI gates before merge (mirror of .github/workflows/ci.yml):
ci: contract fmt-check lint test vuln

bench:
	./scripts/bench-budget.sh

bench-update:
	./scripts/bench-budget.sh --update

clean:
	go clean ./...
	rm -f coverage.out
