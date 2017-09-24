# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope).

The `Add_Small` benchmark adds a single character to the front of a rope with varying initial size, from 1kb to 15kb.

The `Reader` benchmark uses the `io.Reader` and `io.WriterTo` interfaces through `io.Copy` to test reading the contents.

Initial results:

``` sh
% ./compare.sh
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         9319          9088          -2.48%
Benchmark_Add_Small/2500-8         16876         8523          -49.50%
Benchmark_Add_Small/5000-8         26668         8333          -68.75%
Benchmark_Add_Small/7500-8         45487         9112          -79.97%
Benchmark_Add_Small/10000-8        50474         8680          -82.80%
Benchmark_Add_Small/12500-8        65129         9375          -85.61%
Benchmark_Add_Small/15000-8        78159         9272          -88.14%
Benchmark_Reader/100000-8          4783          11650         +143.57%
Benchmark_Reader/125000-8          6151          12829         +108.57%
Benchmark_Reader/150000-8          7745          21360         +175.79%
Benchmark_Reader/175000-8          9347          24743         +164.72%
Benchmark_Reader/200000-8          11040         23237         +110.48%
Benchmark_Remove_Small/1000-8      7691          7989          +3.87%
Benchmark_Remove_Small/2500-8      16135         7919          -50.92%
Benchmark_Remove_Small/5000-8      25689         7841          -69.48%
Benchmark_Remove_Small/7500-8      44470         8659          -80.53%
Benchmark_Remove_Small/10000-8     49089         8355          -82.98%
Benchmark_Remove_Small/12500-8     64493         8609          -86.65%
Benchmark_Remove_Small/15000-8     76567         9480          -87.62%
```
