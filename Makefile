# thie

.PHONY: contract
.PHONY: clear

urlRemote:=registry.cn-shenzhen.aliyuncs.com/owen_app/
urlContract:=$(urlRemote)contract:v$(version)
all: contract

define checkVersion
	@if [$(version) == ""]; then\
		echo "version is be empty";\
		exit 1;\
	fi
endef

contract:
	$(call checkVersion)
	docker build -t $(urlContract) -f ./cmd/contract/Dockerfile ./
	docker push $(urlContract)

clear:
	docker rmi $$(docker images | grep '<none>' | awk '{print $$3}')
