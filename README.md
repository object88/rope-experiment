# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope).

The `Add_Small` benchmark adds a single character to the front of a rope with varying initial size, from 1kb to 15kb.

The `Reader` benchmark uses the `io.Reader` and `io.WriterTo` interfaces through `io.Copy` to test reading the contents.

Initial results:

``` sh
% ./compare.sh
benchmark                       old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8      445           580           +30.34%
Benchmark_Add_Small/2500-8      622           621           -0.16%
Benchmark_Add_Small/5000-8      855           529           -38.13%
Benchmark_Add_Small/7500-8      1048          565           -46.09%
Benchmark_Add_Small/10000-8     1233          475           -61.48%
Benchmark_Add_Small/12500-8     1452          496           -65.84%
Benchmark_Add_Small/15000-8     1661          527           -68.27%
Benchmark_Reader/100000-8       4773          11836         +147.98%
Benchmark_Reader/125000-8       6415          12731         +98.46%
Benchmark_Reader/150000-8       7790          21382         +174.48%
Benchmark_Reader/175000-8       9563          22518         +135.47%
Benchmark_Reader/200000-8       10929         23441         +114.48%
```
