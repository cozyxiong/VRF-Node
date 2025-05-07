VRF_ABI_ARTIFACT := ./abis/VRFManager.sol/VRFManager.json
BLS_ABI_ARTIFACT := ./abis/BLSApkRegistry.sol/BLSApkRegistry.json
BLS_FACTORY_ABI_ARTIFACT := ./abis/VRFFactory.sol/VRFFactory.json


vrf-node:
	env GO111MODULE=on go build  ./cmd/vrf-node

clean:
	rm vrf-node

test:
	go test -v ./...

lint:
	golangci-lint run ./...

bindings:  binding-vrf binding-bls binding-factory

binding-vrf:
	$(eval temp := $(shell mktemp))

	cat $(VRF_ABI_ARTIFACT) \
    	| jq -r .bytecode.object > $(temp)

	cat $(VRF_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg vrf \
		--abi - \
		--out bindings/vrf/vrfmanager.go \
		--type VRFManager \
		--bin $(temp)

		rm $(temp)

binding-bls:
	$(eval temp := $(shell mktemp))

	cat $(BLS_ABI_ARTIFACT) \
    	| jq -r .bytecode.object > $(temp)

	cat $(BLS_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg bls \
		--abi - \
		--out bindings/bls/blsapkregistry.go \
		--type BLSApkRegistry \
		--bin $(temp)

		rm $(temp)


binding-factory:
	$(eval temp := $(shell mktemp))

	cat $(BLS_FACTORY_ABI_ARTIFACT) \
    	| jq -r .bytecode.object > $(temp)

	cat $(BLS_FACTORY_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg vrf \
		--abi - \
		--out bindings/vrf/vrffactory.go \
		--type VRFFactory \
		--bin $(temp)

		rm $(temp)


.PHONY: \
	vrf-node \
	bindings \
	binding-vrf \
	binding-bls \
	binding-factory \
	clean \
	test \
	lint