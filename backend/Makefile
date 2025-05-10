Go = go

all: test

install:
	$(Go) install

test:
	$(Go) test -v ./register/...
	$(Go) test -v ./login/...

build:
	$(Go) build -o Blog