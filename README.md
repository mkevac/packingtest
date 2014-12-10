Simple test for packing with BigEndian and LittleEndian
=======================================================

```
marko@fedora-server:~/goprojects/src/github.com/mkevac/packingtest $ go test -benchtime 10s -bench .
testing: warning: no tests to run
PASS
BenchmarkLittle    1000000000            20.9 ns/op
BenchmarkBig    1000000000            20.9 ns/op
ok      github.com/mkevac/packingtest    45.986s
```
