
CLOUD_BUILD_LOCAL ?= cloud-build-local
DRYRUN           ?= false
PUSH			  ?= false


testlocal:
	$(CLOUD_BUILD_LOCAL) --dryrun=$(DRYRUN) $(PWD)