build-gir: gir/template_gen.go
	cd gir && go build
gir/template_gen.go: gir/template cmd/static-template/static-template
	./cmd/static-template/static-template -name gir < gir/template > $@
cmd/static-template/static-template:
	cd `dirname $@` && go build

.PHONY: build-gir
