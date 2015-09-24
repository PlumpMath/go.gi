
build-gir: gir/template_gen.go
	cd gir && go build
gir/template_gen.go: gir/template gir/cmd/static-template/static-template
	./gir/cmd/static-template/static-template -name gir < gir/template > $@
gir/cmd/static-template/static-template:
	cd `dirname $@` && go build

.PHONY: build-gir
