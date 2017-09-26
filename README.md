# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope), and an improved version that operates over raw byte arrays.

The `Add_Small` benchmark adds a single character to the front of a rope with varying initial size, from 1kb to 15kb.  `Remove_Small` removes a single character.

The `Reader` benchmark uses the `io.Reader` and `io.WriterTo` interfaces through `io.Copy` to test reading the contents.

Initial results:

``` sh
Comparing v1 and v2
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         103249        209292        +102.71%
Benchmark_Add_Small/2500-8         220088        202565        -7.96%
Benchmark_Add_Small/5000-8         427197        202324        -52.64%
Benchmark_Add_Small/7500-8         650595        282697        -56.55%
Benchmark_Add_Small/10000-8        741436        206815        -72.11%
Benchmark_Add_Small/12500-8        945524        243674        -74.23%
Benchmark_Add_Small/15000-8        1118746       282946        -74.71%
Benchmark_Reader/100000-8          4712          11751         +149.38%
Benchmark_Reader/125000-8          6176          12718         +105.93%
Benchmark_Reader/150000-8          7748          21232         +174.03%
Benchmark_Reader/175000-8          9303          22390         +140.68%
Benchmark_Reader/200000-8          10825         23462         +116.74%
Benchmark_Remove_Small/1000-8      175368        275906        +57.33%
Benchmark_Remove_Small/2500-8      385796        176372        -54.28%
Benchmark_Remove_Small/5000-8      778468        174487        -77.59%
Benchmark_Remove_Small/7500-8      1178242       265435        -77.47%
Benchmark_Remove_Small/10000-8     1362808       179598        -86.82%
Benchmark_Remove_Small/12500-8     1746724       224284        -87.16%
Benchmark_Remove_Small/15000-8     2053368       257834        -87.44%

Comparing v1 and v3
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         103249        51521         -50.10%
Benchmark_Add_Small/2500-8         220088        47706         -78.32%
Benchmark_Add_Small/5000-8         427197        50535         -88.17%
Benchmark_Add_Small/7500-8         650595        61165         -90.60%
Benchmark_Add_Small/10000-8        741436        54080         -92.71%
Benchmark_Add_Small/12500-8        945524        59661         -93.69%
Benchmark_Add_Small/15000-8        1118746       71715         -93.59%
Benchmark_Reader/100000-8          4712          9170          +94.61%
Benchmark_Reader/125000-8          6176          10044         +62.63%
Benchmark_Reader/150000-8          7748          18064         +133.14%
Benchmark_Reader/175000-8          9303          17557         +88.72%
Benchmark_Reader/200000-8          10825         18194         +68.07%
Benchmark_Remove_Small/1000-8      175368        106462        -39.29%
Benchmark_Remove_Small/2500-8      385796        75204         -80.51%
Benchmark_Remove_Small/5000-8      778468        76821         -90.13%
Benchmark_Remove_Small/7500-8      1178242       101510        -91.38%
Benchmark_Remove_Small/10000-8     1362808       80107         -94.12%
Benchmark_Remove_Small/12500-8     1746724       97797         -94.40%
Benchmark_Remove_Small/15000-8     2053368       108542        -94.71%
```
