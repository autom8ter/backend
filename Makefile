.PHONY: help

help:   ## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: deploy-endpoints
deploy-endpoints: ## deploy to google endpoints
	gcloud endpoints services deploy vendor/github.com/autom8ter/api/descriptor.pb api_config.yaml

.PHONY: build
build: ## build and submit to google container registry
	gcloud builds submit --tag gcr.io/autom8ter-19/api:1.0 .

.PHONY: release
release: ## create binary release
	cd backend && gox -os=linux

.PHONY: deploy
deploy: ## deploy to Kubernetes
	kubectl apply -f deployment.yaml

clean:
	go fmt ./...
	go vet ./...
	go install ./...
