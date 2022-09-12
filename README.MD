PREREQUISITE

Ensure Docker is installed and running

INFORMATION

A prebuilt binary file, `main` is included
When run in debug mode, some messages are printed to screen and a deplay is introduced to demonstrate the locking/unlocking of the `SafeStack` struct
The Server will try to get a URL 3 times before returning an error

USER COMMANDS:

From the vmWare directory

A user can
    run `docker-compose up`

Or use the Makefiles commands

run `make clean`
    runs `docker-compose down`, removes containers and images, removes `main` binary

run `make run`
    runs `docker-compose up`

run `make binary`
    removes old go binaries and builds new ones

run `make tests`
    runs unit tests


SERVER

The Server is listening on http://localhost:8000
and has endpoints:
        /vmWare
        /

/vmWare
    This endpoint takes parameters `sortKey` and `limit`
    The parameter `sortKey` must be either `views` or `relevanceScore`
    If `sortKey` is not one of the approved the results will be returned unsorted
    The parameter `limit` must be an integer greater than `1` and less than `200`
    If `limit` is less that `1` or greater than `200` all results will be returned
    Requests with incorrect parameter names will fail

/
    / and any other route will return a `PAGE_NOT_FOUND` message

EXAMPLE
    `curl http://localhost:8000/vmWare?sortKey=views&limit=6`