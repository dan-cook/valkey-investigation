# valkey-investigation
Simple valkey client wrapper and basic set/get tests. 

There are two test implementations. One requires that you bring up Valkey explicitly with docker-compose (below), the other leverage the Testcontainers framework.

For the `cache_test.go` tests:
 - Bring valkey up first with `docker-compose up` then run the tests with `go test ./...` (clean test cache if necessary). This will run all tests including the testcontainers implementation.

For the `cache_no_docker_test.go` tests:
- This can be run independently of the prior tests requiring a manual start from docker by targeting it with the following: `go test '-run=^TestCacheWithTestContainers$' -v  ./...`.
