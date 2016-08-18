运行go test -bench=. 得到运行结果:

BenchmarkPopCount-8         300000000            5.00 ns/op
BenchmarkPopCountRewrite-8  100000000           13.1 ns/op

结论：PopCount()每次调用平均用时5.00ns, PopCountRewrite()每次调用平均用时13.1ns,重写前性能更好
