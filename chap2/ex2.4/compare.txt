运行go test -bench=. 得到运行结果:

BenchmarkPopCount-8         300000000            5.95 ns/op
BenchmarkPopCountRewrite2-8  100000000            79.1 ns/op

结论：PopCount()每次调用平均用时5.95ns, PopCountRewrite2()每次调用平均用时79.1ns,重写前性能更好
