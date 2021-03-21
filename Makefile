GOPATH=$(HOME)/buildpath
OUTDIR=$(HOME)/go/src/project

generate_go:
	mkdir -p $(GOPATH) && mkdir -p $(OUTDIR) && chmod 777 $(GOPATH)
	rm -rf $(OUTDIR)/*
	echo "module project" > $(OUTDIR)/go.mod
	go build -o bootstrap && \
		./bootstrap new -i schema.sample.json -o $(OUTDIR) -p project
	export GOPATH=$(HOME)/go && goimports -l -w $(OUTDIR)/

	cp Makefile.template $(OUTDIR)/Makefile
	cp run.sh $(OUTDIR)/run.sh
	cp docker-compose.yml $(OUTDIR)/docker-compose.yml

go_import:
	goimports -l -w $(OUTDIR)