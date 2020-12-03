check_install:
	which swagger || GO111MODULES=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger:check_install
	GO111MODULES=off swagger generate spec -o ./swagger.yaml --scan-models