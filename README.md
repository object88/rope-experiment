# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope).

The `Add_Small` benchmark adds a single character to the front of a rope with varying initial size, from 1kb to 15kb.

The `Reader` benchmark uses the `io.Reader` and `io.WriterTo` interfaces through `io.Copy` to test reading the contents.

Initial results:

``` sh
Comparing v1 and v2
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         9486          9725          +2.52%
Benchmark_Add_Small/2500-8         17291         9548          -44.78%
Benchmark_Add_Small/5000-8         27719         10369         -62.59%
Benchmark_Add_Small/7500-8         46882         11459         -75.56%
Benchmark_Add_Small/10000-8        50978         10585         -79.24%
Benchmark_Add_Small/12500-8        68156         11184         -83.59%
Benchmark_Add_Small/15000-8        79952         11529         -85.58%
Benchmark_Reader/100000-8          4793          12066         +151.74%
Benchmark_Reader/125000-8          6344          13032         +105.42%
Benchmark_Reader/150000-8          7917          22732         +187.13%
Benchmark_Reader/175000-8          9516          25171         +164.51%
Benchmark_Reader/200000-8          11089         27225         +145.51%
Benchmark_Remove_Small/1000-8      8326          8100          -2.71%
Benchmark_Remove_Small/2500-8      17702         7822          -55.81%
Benchmark_Remove_Small/5000-8      26836         8620          -67.88%
Benchmark_Remove_Small/7500-8      52629         8689          -83.49%
Benchmark_Remove_Small/10000-8     56875         9258          -83.72%
Benchmark_Remove_Small/12500-8     65974         8547          -87.04%
Benchmark_Remove_Small/15000-8     93381         9819          -89.49%
Comparing v1 and v3
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         9486          16392         +72.80%
Benchmark_Add_Small/2500-8         17291         16894         -2.30%
Benchmark_Add_Small/5000-8         27719         16098         -41.92%
Benchmark_Add_Small/7500-8         46882         17153         -63.41%
Benchmark_Add_Small/10000-8        50978         16328         -67.97%
Benchmark_Add_Small/12500-8        68156         16272         -76.13%
Benchmark_Add_Small/15000-8        79952         17201         -78.49%
Benchmark_Reader/100000-8          4793          8926          +86.23%
Benchmark_Reader/125000-8          6344          9827          +54.90%
Benchmark_Reader/150000-8          7917          16306         +105.96%
Benchmark_Reader/175000-8          9516          17130         +80.01%
Benchmark_Reader/200000-8          11089         17844         +60.92%
Benchmark_Remove_Small/1000-8      8326          15460         +85.68%
Benchmark_Remove_Small/2500-8      17702         15667         -11.50%
Benchmark_Remove_Small/5000-8      26836         15919         -40.68%
Benchmark_Remove_Small/7500-8      52629         16160         -69.29%
Benchmark_Remove_Small/10000-8     56875         15358         -73.00%
Benchmark_Remove_Small/12500-8     65974         15956         -75.81%
Benchmark_Remove_Small/15000-8     93381         16877         -81.93%
```
