all: docker

docker: nextcloud-kcintegrate
	docker-compose build
	docker-compose up -d
	./run.sh docker

nextcloud-kcintegrate:
	go build

bin: nextcloud-kcintegrate
