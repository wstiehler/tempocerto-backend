has_java:=$(shell which java)

setup:
	npm i -g c4builder

new-project:
	c4builder new &&
	cd project &&
	c4builder

config:
	cd zpe-systems && npm run config

watch-standalone:
	cd zpe-systems && c4builder site --watch