vmware:

clean:
	docker-compose down
	docker rmi server-img
	cd server && rm -f main

binary:
	cd server && make

run:
	docker-compose up

tests:
	@cd server/getUrls && go test -v
	@cd server/safeStack && go test -v
	@cd server/utils && go test -v