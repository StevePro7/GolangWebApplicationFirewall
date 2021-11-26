default: image

image:
	docker run --rm --privileged multiarch/qemu-user-static:register --reset
	docker build --cpuset-cpus 0 --pull -t calico/go-build:stevepro -f Dockerfile.amd64 .

