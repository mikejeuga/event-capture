#####################################################################################################################################################
#                             									Mike's Go Makefile                              										#
#                  										Copyright (C) 2022 - Michael Jeuga															#
#																																					#
#####################################################################################################################################################

repo=$(shell basename "`pwd`")

gopher:
	@git init
	@go mod init github.com/mikejeuga/$(repo)
	@go mod tidy
	@touch .gitignore


docker:
	@touch Dockerfile
	@touch docker-compose.yml


t: test
test:
	@docker-compose -f docker-compose.yml up -d
	@go test -v ./...
	@docker-compose down

ut: unit-test
unit-test:
	@go test -v -tags=unit ./...

at: acceptance-test
acceptance-test:
	@docker-compose -f docker-compose.yml up -d
	@go test -v -tags acceptance ./...
	@docker-compose down


pull:
	@git pull -r

container:
	@touch "$n".yml

ic: init
init:
	git add .
	git commit -m "Initial commit"
	git remote add origin git@github.com:mikejeuga/${repo}.git
	git branch -M main
	git push -u origin main

c: commit
commit:
	@git commit -am "$m"
	@git pull --rebase
	git push