
define build_any
	@echo "\nbuilding $(PATH) ...\n\n"
	rm -rf ./dist/$(PATH) || true
	go build -mod vendor -v -o ./dist/	 ./$(PATH)/main.go
	# cp -rpv ./configs/*.yaml ./dist/$(1)/$(3)/$(2)
	# cp -rpv ./configs/$(1)/$(2)/*.yaml ./dist/$(1)/$(3)/$(2)
endef

build:
	$(call build_any)