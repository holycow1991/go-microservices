proto:
	protoc services/**/pkg/**/pb/*.proto --go_out=.
run-gateway:
	go run services/gateway-svc/cmd/main.go
run-auth:
	go run services/auth-svc/cmd/main.go
run-products:
	go run services/products-svc/cmd/main.go
run-orders:
	go run services/orders-svc/cmd/main.go
