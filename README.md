# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope), and an improved version that operates over raw byte arrays.

The `Add_Small` benchmark adds a single character to the front of a rope with varying initial size, from 1kb to 15kb.  `Remove_Small` removes a single character.

The `Reader` benchmark uses the `io.Reader` and `io.WriterTo` interfaces through `io.Copy` to test reading the contents.

Initial results:

``` sh
Comparing v1 and v2
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         102669        50815         -50.51%
Benchmark_Add_Small/2500-8         221716        47013         -78.80%
Benchmark_Add_Small/5000-8         422432        50599         -88.02%
Benchmark_Add_Small/7500-8         636135        65228         -89.75%
Benchmark_Add_Small/10000-8        737612        53000         -92.81%
Benchmark_Add_Small/12500-8        954379        60657         -93.64%
Benchmark_Add_Small/15000-8        1112471       71466         -93.58%
Benchmark_Reader/100000-8          4762          11741         +146.56%
Benchmark_Reader/125000-8          6098          12801         +109.92%
Benchmark_Reader/150000-8          7670          21119         +175.35%
Benchmark_Reader/175000-8          9328          22477         +140.96%
Benchmark_Reader/200000-8          10994         23310         +112.02%
Benchmark_Remove_Small/1000-8      176368        100031        -43.28%
Benchmark_Remove_Small/2500-8      381990        72034         -81.14%
Benchmark_Remove_Small/5000-8      768943        73278         -90.47%
Benchmark_Remove_Small/7500-8      1166440       97591         -91.63%
Benchmark_Remove_Small/10000-8     1353408       76516         -94.35%
Benchmark_Remove_Small/12500-8     1736040       87001         -94.99%
Benchmark_Remove_Small/15000-8     2040696       102887        -94.96%

Comparing v1 and v3
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         102669        52062         -49.29%
Benchmark_Add_Small/2500-8         221716        47900         -78.40%
Benchmark_Add_Small/5000-8         422432        51818         -87.73%
Benchmark_Add_Small/7500-8         636135        62199         -90.22%
Benchmark_Add_Small/10000-8        737612        56091         -92.40%
Benchmark_Add_Small/12500-8        954379        58341         -93.89%
Benchmark_Add_Small/15000-8        1112471       72622         -93.47%
Benchmark_Reader/100000-8          4762          8920          +87.32%
Benchmark_Reader/125000-8          6098          9803          +60.76%
Benchmark_Reader/150000-8          7670          16184         +111.00%
Benchmark_Reader/175000-8          9328          16945         +81.66%
Benchmark_Reader/200000-8          10994         17764         +61.58%
Benchmark_Remove_Small/1000-8      176368        105628        -40.11%
Benchmark_Remove_Small/2500-8      381990        77169         -79.80%
Benchmark_Remove_Small/5000-8      768943        77529         -89.92%
Benchmark_Remove_Small/7500-8      1166440       101795        -91.27%
Benchmark_Remove_Small/10000-8     1353408       79207         -94.15%
Benchmark_Remove_Small/12500-8     1736040       98278         -94.34%
Benchmark_Remove_Small/15000-8     2040696       115327        -94.35%
```
