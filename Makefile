run-little-mimic = docker run -d -p 8080:8080 little_mimic
kill-little-mimic = docker kill $1

test:
	go test

integration:
	go test -tags=integration

little-mimic:
	docker build -t little_mimic little_mimic/

test-openstack: little-mimic
	$(eval CTID:=$(shell ${run-little-mimic}))
	go test -tags=openstack
	$(call kill-little-mimic, ${CTID})
