.PHONY: generate build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/give-help-server cmd/give-help-server/main.go

generate:
	rm -Rf generated
	mkdir generated
	swagger generate server -t generated  -P models.LoggedUser --exclude-main --skip-validation -f api/swagger.yml -r LICENSE

package: generate build
	cd ./bin/ && zip -r code.zip *

deploy: package
	sls deploy --verbose
	rm ./bin/code.zip