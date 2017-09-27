# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope), and an improved version that operates over raw byte arrays.

The `Add_Small` benchmark adds a single character to the front of a rope with varying initial size, from 1kb to 15kb.  `Remove_Small` removes a single character.

The `Reader` benchmark uses the `io.Reader` and `io.WriterTo` interfaces through `io.Copy` to test reading the contents.

Initial results:

``` sh
Comparing v1 and v2
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         104942        51148         -51.26%
Benchmark_Add_Small/2500-8         220855        47480         -78.50%
Benchmark_Add_Small/5000-8         418908        51563         -87.69%
Benchmark_Add_Small/7500-8         638515        65619         -89.72%
Benchmark_Add_Small/10000-8        753561        53250         -92.93%
Benchmark_Add_Small/12500-8        976434        61515         -93.70%
Benchmark_Add_Small/15000-8        1137186       73093         -93.57%
Benchmark_Reader/100000-8          4987          11717         +134.95%
Benchmark_Reader/125000-8          6285          12831         +104.15%
Benchmark_Reader/150000-8          7832          21112         +169.56%
Benchmark_Reader/175000-8          9317          22305         +139.40%
Benchmark_Reader/200000-8          11270         23927         +112.31%
Benchmark_Remove_Small/1000-8      189820        105098        -44.63%
Benchmark_Remove_Small/2500-8      407168        74523         -81.70%
Benchmark_Remove_Small/5000-8      777083        73651         -90.52%
Benchmark_Remove_Small/7500-8      1200555       98245         -91.82%
Benchmark_Remove_Small/10000-8     1353792       76699         -94.33%
Benchmark_Remove_Small/12500-8     1743920       86217         -95.06%
Benchmark_Remove_Small/15000-8     2048785       101931        -95.02%

Comparing v1 and v3
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         104942        58558         -44.20%
Benchmark_Add_Small/2500-8         220855        83797         -62.06%
Benchmark_Add_Small/5000-8         418908        78743         -81.20%
Benchmark_Add_Small/7500-8         638515        81214         -87.28%
Benchmark_Add_Small/10000-8        753561        85171         -88.70%
Benchmark_Add_Small/12500-8        976434        85154         -91.28%
Benchmark_Add_Small/15000-8        1137186       85334         -92.50%
Benchmark_Reader/100000-8          4987          10319         +106.92%
Benchmark_Reader/125000-8          6285          10944         +74.13%
Benchmark_Reader/150000-8          7832          18445         +135.51%
Benchmark_Reader/175000-8          9317          19529         +109.61%
Benchmark_Reader/200000-8          11270         20583         +82.64%
Benchmark_Remove_Small/1000-8      189820        93785         -50.59%
Benchmark_Remove_Small/2500-8      407168        95026         -76.66%
Benchmark_Remove_Small/5000-8      777083        102738        -86.78%
Benchmark_Remove_Small/7500-8      1200555       97547         -91.87%
Benchmark_Remove_Small/10000-8     1353792       96375         -92.88%
Benchmark_Remove_Small/12500-8     1743920       96555         -94.46%
Benchmark_Remove_Small/15000-8     2048785       104609        -94.89%
```
