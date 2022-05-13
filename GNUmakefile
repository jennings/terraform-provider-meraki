default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	go test ./... -v $(TESTARGS) -timeout 120m
