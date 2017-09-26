# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope).

The `Add_Small` benchmark adds a single character to the front of a rope with varying initial size, from 1kb to 15kb.  `Remove_Small` removes a single character.

The `Reader` benchmark uses the `io.Reader` and `io.WriterTo` interfaces through `io.Copy` to test reading the contents.

Initial results:

``` sh
Comparing v1 and v2
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         540588        209865        -61.18%
Benchmark_Add_Small/2500-8         1306477       208536        -84.04%
Benchmark_Add_Small/5000-8         2620812       210444        -91.97%
Benchmark_Add_Small/7500-8         3966476       293338        -92.60%
Benchmark_Add_Small/10000-8        5188767       211355        -95.93%
Benchmark_Add_Small/12500-8        6058574       244327        -95.97%
Benchmark_Add_Small/15000-8        7183792       287554        -96.00%
Benchmark_Reader/100000-8          4808          12423         +158.38%
Benchmark_Reader/125000-8          6365          13388         +110.34%
Benchmark_Reader/150000-8          9210          21469         +133.11%
Benchmark_Reader/175000-8          9386          23047         +145.55%
Benchmark_Reader/200000-8          11171         23381         +109.30%
Benchmark_Remove_Small/1000-8      558419        279104        -50.02%
Benchmark_Remove_Small/2500-8      1277251       176451        -86.19%
Benchmark_Remove_Small/5000-8      2663517       172185        -93.54%
Benchmark_Remove_Small/7500-8      4240808       297988        -92.97%
Benchmark_Remove_Small/10000-8     5170292       184936        -96.42%
Benchmark_Remove_Small/12500-8     6779202       237875        -96.49%
Benchmark_Remove_Small/15000-8     7472133       310256        -95.85%

Comparing v1 and v3
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         540588        51858         -90.41%
Benchmark_Add_Small/2500-8         1306477       48451         -96.29%
Benchmark_Add_Small/5000-8         2620812       51869         -98.02%
Benchmark_Add_Small/7500-8         3966476       62581         -98.42%
Benchmark_Add_Small/10000-8        5188767       54624         -98.95%
Benchmark_Add_Small/12500-8        6058574       60299         -99.00%
Benchmark_Add_Small/15000-8        7183792       72259         -98.99%
Benchmark_Reader/100000-8          4808          8863          +84.34%
Benchmark_Reader/125000-8          6365          9865          +54.99%
Benchmark_Reader/150000-8          9210          16119         +75.02%
Benchmark_Reader/175000-8          9386          17146         +82.68%
Benchmark_Reader/200000-8          11171         17790         +59.25%
Benchmark_Remove_Small/1000-8      558419        104552        -81.28%
Benchmark_Remove_Small/2500-8      1277251       77879         -93.90%
Benchmark_Remove_Small/5000-8      2663517       78132         -97.07%
Benchmark_Remove_Small/7500-8      4240808       101882        -97.60%
Benchmark_Remove_Small/10000-8     5170292       79707         -98.46%
Benchmark_Remove_Small/12500-8     6779202       98437         -98.55%
Benchmark_Remove_Small/15000-8     7472133       108643        -98.55%

```
