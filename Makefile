.PHONY: help

help:   ## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

deploy: ## deploy to google endpoints
	gcloud endpoints services deploy vendor/github.com/autom8ter/api/descriptor.pb api_config.yaml

build: ## build and submit to google container registry
	gcloud builds submit --tag gcr.io/autom8ter-19/api:1.0 .