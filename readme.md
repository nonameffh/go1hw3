Результаты:
```
$ go test -bench . -benchmem

goos: darwin
goarch: amd64
BenchmarkSlow-16              52          23134388 ns/op        19131893 B/op     195822 allocs/op
BenchmarkFast-16             831           1456990 ns/op          492526 B/op       6410 allocs/op
```

Пример результатов с которыми будет сравниваться:
```
$ go test -bench . -benchmem

goos: windows
goarch: amd64
BenchmarkSlow-8            10        142703250 ns/op     336887900 B/op      284175 allocs/op
BenchmarkSolution-8       500          2782432 ns/op        559910 B/op       10422 allocs/op

PASS

ok coursera/hw3 3.897s
```