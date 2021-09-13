.PHONY: generate test

default: generate test

generate:
	go generate -x internal/generator.go

clean-generate:
	 rm ./**/*_gen.go

test:
	go test ./... --cover
