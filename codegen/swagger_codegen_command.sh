java -jar swagger-codegen-cli.jar generate -c swagger_codegen_config.json -i swagger_spec.json -l go -o client
rm client/defs_*.go
gofmt -s -w client/*.go
