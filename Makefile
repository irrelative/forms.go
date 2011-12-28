include $(GOROOT)/src/Make.inc

TARG=forms
GOFMT=gofmt -s -spaces=true -tabindent=false -tabwidth=4

GOFILES=\
  forms.go\

include $(GOROOT)/src/Make.pkg

format:
	${GOFMT} -w ${GOFILES}

formtest:
	8g formtest.go
	8l -o formtest formtest.8
