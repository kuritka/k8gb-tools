.PHONY: lint
lint:
	golangci-lint run
	go test  ./...

.PHONY: dns-tools
dns-tools:
	@kubectl -n k8gb get svc k8gb-coredns
	@kubectl -n k8gb run -it @

.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
