# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope).

The `Add_Small` benchmark adds a single character to the front of a rope with varying initial size, from 1kb to 15kb.  `Remove_Small` removes a single character.

The `Reader` benchmark uses the `io.Reader` and `io.WriterTo` interfaces through `io.Copy` to test reading the contents.

Initial results:

``` sh
Comparing v1 and v2
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         538880        208572        -61.30%
Benchmark_Add_Small/2500-8         1298961       200680        -84.55%
Benchmark_Add_Small/5000-8         2540128       203213        -92.00%
Benchmark_Add_Small/7500-8         3861434       290402        -92.48%
Benchmark_Add_Small/10000-8        4833106       209279        -95.67%
Benchmark_Add_Small/12500-8        6040841       248040        -95.89%
Benchmark_Add_Small/15000-8        7200316       281155        -96.10%
Benchmark_Reader/100000-8          4745          12107         +155.15%
Benchmark_Reader/125000-8          6142          12451         +102.72%
Benchmark_Reader/150000-8          7737          21195         +173.94%
Benchmark_Reader/175000-8          9395          22284         +137.19%
Benchmark_Reader/200000-8          11684         23352         +99.86%
Benchmark_Remove_Small/1000-8      552115        274334        -50.31%
Benchmark_Remove_Small/2500-8      1275321       175403        -86.25%
Benchmark_Remove_Small/5000-8      2479431       172723        -93.03%
Benchmark_Remove_Small/7500-8      3754885       261538        -93.03%
Benchmark_Remove_Small/10000-8     4814577       176513        -96.33%
Benchmark_Remove_Small/12500-8     6391228       224494        -96.49%
Benchmark_Remove_Small/15000-8     7334311       258563        -96.47%
Comparing v1 and v3
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         538880        226119        -58.04%
Benchmark_Add_Small/2500-8         1298961       222158        -82.90%
Benchmark_Add_Small/5000-8         2540128       217458        -91.44%
Benchmark_Add_Small/7500-8         3861434       308512        -92.01%
Benchmark_Add_Small/10000-8        4833106       223415        -95.38%
Benchmark_Add_Small/12500-8        6040841       260910        -95.68%
Benchmark_Add_Small/15000-8        7200316       303280        -95.79%
Benchmark_Reader/100000-8          4745          9547          +101.20%
Benchmark_Reader/125000-8          6142          10615         +72.83%
Benchmark_Reader/150000-8          7737          17225         +122.63%
Benchmark_Reader/175000-8          9395          18146         +93.15%
Benchmark_Reader/200000-8          11684         19497         +66.87%
Benchmark_Remove_Small/1000-8      552115        106242        -80.76%
Benchmark_Remove_Small/2500-8      1275321       78939         -93.81%
Benchmark_Remove_Small/5000-8      2479431       84923         -96.57%
Benchmark_Remove_Small/7500-8      3754885       110284        -97.06%
Benchmark_Remove_Small/10000-8     4814577       79175         -98.36%
Benchmark_Remove_Small/12500-8     6391228       93307         -98.54%
Benchmark_Remove_Small/15000-8     7334311       96098         -98.69%
```
