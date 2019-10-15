MAKE   := /usr/bin/make
SUBDIR := $(wildcard */)

.PHONY:build $(SUBDIR) clean
build: $(SUBDIR)
$(SUBDIR):
	$(MAKE) -C $@

clean: $(SUBDIR)
	# TODO: 本当は各サブディレクトリにcleanを任せたい
	rm -f $(addsuffix "/.build",$(SUBDIR))