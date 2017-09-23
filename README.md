# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope).

Initial results:

``` sh
% ./compare.sh 
benchmark                       old ns/op     new ns/op     delta
Benchmark_Add_Small/500-8       304           370           +21.71%
Benchmark_Add_Small/750-8       345           404           +17.10%
Benchmark_Add_Small/1000-8      348           375           +7.76%
Benchmark_Add_Small/2500-8      738           542           -26.56%
Benchmark_Add_Small/5000-8      1143          840           -26.51%
Benchmark_Add_Small/7500-8      2784          830           -70.19%
Benchmark_Add_Small/10000-8     3526          3047          -13.58%
Benchmark_Add_Small/12500-8     10746         1014          -90.56%
Benchmark_Add_Small/15000-8     10253         4250          -58.55%
Benchmark_Add_Small/17500-8     15921         2617          -83.56%
Benchmark_Add_Small/20000-8     3812          999           -73.79%
```
