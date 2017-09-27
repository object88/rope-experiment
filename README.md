# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope), and an improved version that operates over raw byte arrays.

The `Add_Small` benchmark adds a single character to the front of a rope with varying initial size, from 1kb to 15kb.  `Remove_Small` removes a single character.

The `Reader` benchmark uses the `io.Reader` and `io.WriterTo` interfaces through `io.Copy` to test reading the contents.

Initial results:

``` sh
Comparing v1 and v2
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         104336        50643         -51.46%
Benchmark_Add_Small/2500-8         220362        49859         -77.37%
Benchmark_Add_Small/5000-8         417755        50748         -87.85%
Benchmark_Add_Small/7500-8         636626        65173         -89.76%
Benchmark_Add_Small/10000-8        731347        51496         -92.96%
Benchmark_Add_Small/12500-8        942344        61424         -93.48%
Benchmark_Add_Small/15000-8        1118221       73256         -93.45%
Benchmark_Reader/100000-8          4917          12108         +146.25%
Benchmark_Reader/125000-8          6207          13219         +112.97%
Benchmark_Reader/150000-8          7841          21182         +170.14%
Benchmark_Reader/175000-8          9792          22283         +127.56%
Benchmark_Reader/200000-8          11262         23359         +107.41%
Benchmark_Remove_Small/1000-8      185240        105762        -42.91%
Benchmark_Remove_Small/2500-8      384784        73937         -80.78%
Benchmark_Remove_Small/5000-8      767797        72738         -90.53%
Benchmark_Remove_Small/7500-8      1167559       99745         -91.46%
Benchmark_Remove_Small/10000-8     1432726       77285         -94.61%
Benchmark_Remove_Small/12500-8     1773000       86258         -95.13%
Benchmark_Remove_Small/15000-8     2053741       104725        -94.90%

Comparing v1 and v3
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         104336        55808         -46.51%
Benchmark_Add_Small/2500-8         220362        78773         -64.25%
Benchmark_Add_Small/5000-8         417755        84340         -79.81%
Benchmark_Add_Small/7500-8         636626        83649         -86.86%
Benchmark_Add_Small/10000-8        731347        80843         -88.95%
Benchmark_Add_Small/12500-8        942344        84138         -91.07%
Benchmark_Add_Small/15000-8        1118221       85474         -92.36%
Benchmark_Reader/100000-8          4917          10224         +107.93%
Benchmark_Reader/125000-8          6207          10930         +76.09%
Benchmark_Reader/150000-8          7841          18627         +137.56%
Benchmark_Reader/175000-8          9792          19493         +99.07%
Benchmark_Reader/200000-8          11262         20667         +83.51%
Benchmark_Remove_Small/1000-8      185240        97697         -47.26%
Benchmark_Remove_Small/2500-8      384784        111250        -71.09%
Benchmark_Remove_Small/5000-8      767797        103643        -86.50%
Benchmark_Remove_Small/7500-8      1167559       100216        -91.42%
Benchmark_Remove_Small/10000-8     1432726       115425        -91.94%
Benchmark_Remove_Small/12500-8     1773000       95439         -94.62%
Benchmark_Remove_Small/15000-8     2053741       102617        -95.00%
```
