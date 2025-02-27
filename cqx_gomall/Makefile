.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I=../../idl --type HTTP --service frontend --module github.com/MrLittle05/Gomall/app/frontend --idl ../../idl/frontend/order_page.proto

.PHONY: gen-cart
gen-user:
	#  生成rpc客户端
  @cd rpc_gen && cwgo client --type RPC --service order --module github.com/MrLittle05/Gomall/rpc_gen -I=../idl --idl ../idl/order.proto
	#  基于已生成的客户端,生成rpc服务端,并使用生成的客户端作为依赖
	@cd app/order && cwgo server --type RPC --service order --module github.com/MrLittle05/Gomall/app/order --pass "-use github.com/MrLittle05/Gomall/rpc_gen/kitex_gen" -I=../../idl --idl ../../idl/order.proto
	

.PHONY: start-docker
start-docker:
	docker-compose up -d