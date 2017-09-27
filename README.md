# rope-experiment

Experiment to evaluate different implementations of the rope data structure

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope), and an improved version that operates over raw byte arrays.

The `Add_Small` benchmark adds a single character to the front of a rope with varying initial size, from 1kb to 15kb.  `Remove_Small` removes a single character.

The `Reader` benchmark uses the `io.Reader` and `io.WriterTo` interfaces through `io.Copy` to test reading the contents.

Initial results:

``` sh
Comparing v1 and v2
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         105055        51304         -51.16%
Benchmark_Add_Small/2500-8         220795        50698         -77.04%
Benchmark_Add_Small/5000-8         419685        51533         -87.72%
Benchmark_Add_Small/7500-8         634878        65868         -89.63%
Benchmark_Add_Small/10000-8        739512        53706         -92.74%
Benchmark_Add_Small/12500-8        947487        61542         -93.50%
Benchmark_Add_Small/15000-8        1115115       72962         -93.46%
Benchmark_Reader/100000-8          4797          11798         +145.95%
Benchmark_Reader/125000-8          6228          12846         +106.26%
Benchmark_Reader/150000-8          7827          21358         +172.88%
Benchmark_Reader/175000-8          9298          22572         +142.76%
Benchmark_Reader/200000-8          10829         23346         +115.59%
Benchmark_Remove_Small/1000-8      176616        100926        -42.86%
Benchmark_Remove_Small/2500-8      383316        71681         -81.30%
Benchmark_Remove_Small/5000-8      777418        73519         -90.54%
Benchmark_Remove_Small/7500-8      1181402       98536         -91.66%
Benchmark_Remove_Small/10000-8     1360214       76367         -94.39%
Benchmark_Remove_Small/12500-8     1745512       85150         -95.12%
Benchmark_Remove_Small/15000-8     2050096       102419        -95.00%

Comparing v1 and v3
benchmark                          old ns/op     new ns/op     delta
Benchmark_Add_Small/1000-8         105055        55628         -47.05%
Benchmark_Add_Small/2500-8         220795        78493         -64.45%
Benchmark_Add_Small/5000-8         419685        78823         -81.22%
Benchmark_Add_Small/7500-8         634878        83109         -86.91%
Benchmark_Add_Small/10000-8        739512        84548         -88.57%
Benchmark_Add_Small/12500-8        947487        84668         -91.06%
Benchmark_Add_Small/15000-8        1115115       86248         -92.27%
Benchmark_Reader/100000-8          4797          10336         +115.47%
Benchmark_Reader/125000-8          6228          11102         +78.26%
Benchmark_Reader/150000-8          7827          18642         +138.18%
Benchmark_Reader/175000-8          9298          19791         +112.85%
Benchmark_Reader/200000-8          10829         20624         +90.45%
Benchmark_Remove_Small/1000-8      176616        94184         -46.67%
Benchmark_Remove_Small/2500-8      383316        95048         -75.20%
Benchmark_Remove_Small/5000-8      777418        95706         -87.69%
Benchmark_Remove_Small/7500-8      1181402       97128         -91.78%
Benchmark_Remove_Small/10000-8     1360214       95435         -92.98%
Benchmark_Remove_Small/12500-8     1745512       96838         -94.45%
Benchmark_Remove_Small/15000-8     2050096       104582        -94.90%
```
