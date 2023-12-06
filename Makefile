.PHONY: all
all: help

.PHONY: help
help:
	@echo "Usage: make gen"

.PHONY: gen-product
gen-product:
	cd app/product && cwgo server -I ../../idl --type RPC --service product --module github.com/baiyutang/gomall/app/product --idl ../../idl/product.proto

.PHONY: run-product
run-product:
	cd app/product && go run .

.PHONY: frontend
frontend:
	cd app/frontend && go run .