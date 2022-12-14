PREREQUISITE

* Ensure Docker is installed and running

INFORMATION

* A prebuilt binary file, `main`, is included
* When run in `debug mode`, some helpful messages are printed to screen. A delay is introduced to demonstrate the locking and unlocking of the `SafeStack` struct. To disable, set `debug` to `false` in the `server/Dockerfile.server`
* The Server will try to get a URL 3 times before returning an error

USER COMMANDS

* Run commands from the root directory
* Makefiles commands
    * run `make clean`
        * runs `docker-compose down`, removes containers and images, removes `main` binary
    * run `make run`
        * runs `docker-compose up`
    * run `make binary`
        * removes old go binaries and builds new ones
    * run `make tests`
        * runs unit tests
* Docker commands
    * run `docker-compose up`

SERVER

* The Server is listening on http://localhost:8000 and has endpoints:
    * /vmWare
    * /
* /vmWare
    * This endpoint takes parameters `sortKey` and `limit`
    * The parameter `sortKey` must be either `views` or `relevanceScore`
    * If `sortKey` is not one of the approved the results will be returned unsorted
    * The parameter `limit` must be an integer greater than `1` and less than `200`
    * If `limit` is less that `1` or greater than `200` all results will be returned
    * Requests with incorrect parameter names will fail*
* /
    * / and any other route will return a `PAGE_NOT_FOUND` message

EXAMPLE
* `curl http://localhost:8000/vmWare?sortKey=views&limit=6`

DOCKER 2 KUBERNETES
* This was ran and tested locally using a `docker-compose.yaml` template file. In order to meet the requirements, the tool `kompose` was used to create `server-deployment.yaml` and `server-service.yaml`
* To apply to to kubernetes cluster run `kubectl apply -f server-deployment.yaml,server-service.yaml`