check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger_json: check_install
	swagger generate spec -o ./swagger.json

serve:
	swagger serve --no-open ./swagger.json

serve_swagger:
	swagger serve --flavor=swagger --no-open ./swagger.json	

jwt_run:
	JWT_SECRET=eUbP9shywUygMx7u MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" MONGO_DATABASE=demo go run *.go