# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope).

Initial results:

``` sh
% ./compare.sh 
benchmark                       old ns/op     new ns/op     delta
Benchmark_Add_Small/500-8       348           424           +21.84%
Benchmark_Add_Small/750-8       364           483           +32.69%
Benchmark_Add_Small/1000-8      402           356           -11.44%
Benchmark_Add_Small/2500-8      832           683           -17.91%
Benchmark_Add_Small/5000-8      1133          762           -32.74%
Benchmark_Add_Small/7500-8      3646          816           -77.62%
Benchmark_Add_Small/10000-8     5530          3358          -39.28%
Benchmark_Add_Small/12500-8     5080          1184          -76.69%
Benchmark_Add_Small/15000-8     3662          1230          -66.41%
Benchmark_Add_Small/17500-8     4079          6828          +67.39%
Benchmark_Add_Small/20000-8     5718          6891          +20.51%
Benchmark_Reader/100000-8       4992          11756         +135.50%
Benchmark_Reader/125000-8       6320          12838         +103.13%
Benchmark_Reader/150000-8       8087          21383         +164.41%
Benchmark_Reader/175000-8       9964          22397         +124.78%
Benchmark_Reader/200000-8       11078         23248         +109.86%
```
