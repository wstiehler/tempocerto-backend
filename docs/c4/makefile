has_java:=$(shell which java)

setup:
	npm i -g c4builder

new-project:
	c4builder new &&
	cd project &&
	c4builder

config:
	cd tempocerto-systems && npm run config

watch-standalone:
	cd tempocerto-systems && c4builder site --watch
