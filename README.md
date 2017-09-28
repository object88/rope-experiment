# rope-experiment

Experiment to evaluate different implementations of the rope data structure.  Conclusions of this (ongoing?) experiement are implemented in [rope](https://github.com/object88/rope).

The [rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure)) is a way to handle ad-hoc edits to largish blocks of text, as one might make in a text editor.  To understand the difference in CPU usage for an edit, the project compares bare string manipulation with a basic rope implementation [translated from JavaScript](https://github.com/component/rope), and an improved version that operates over raw byte arrays.

## Implementations

`v1` is a simple string.  Edits are made by reconstructing the whole string using `bytes.Buffer`.

`v2` is an actual rope implementation, using substrings as the data structure in each node in the B-Tree structure.

`v3` is similar to `v2`, except that it uses a `[]byte` as the data structure in each node.  This _seems_ to have little effect on adding or removing characters to the rope, but gives a slight performance boost when reading out via an `io.Reader`.

Further work may attempt to experiment further:

* May use a fixed-size `[]byte` for each leaf.  Theoretically, if changes do not reallocate the memory used even in a leaf, we may see better performance.  This may be of diminishing return for the complexity, especially as there would need to be some temporary buffer allocated to shift bytes within the fixed buffer for insertions and deletions.

* May have two implementations, one for ropes which contain just ASCII characters, and one which contains UFT8 or Unicode characters.  Operations on pure ASCII strings are much faster, as we don't need to map between rune offset and byte offset.  The rope may scan its initial content for any non-ASCII value, and if not found, start off with faster operations.  Once a non-ASCII rune is inserted, the three would need to be reassembled.

## Tests

The `Add_Small` benchmark adds a single character to the front of a rope with varying initial size, from 1kb to 15kb.  `Remove_Small` removes a single character.

The `Reader` benchmark uses the `io.Reader` and `io.WriterTo` interfaces through `io.Copy` to test reading the contents.

Initial results:

``` sh
Comparing v1 and v2
benchmark                          old ns/op     new ns/op     delta
Benchmark_Alter/1000-8             187914        85284         -54.62%
Benchmark_Alter/2500-8             413970        91478         -77.90%
Benchmark_Alter/5000-8             783128        95754         -87.77%
Benchmark_Alter/7500-8             1163277       97104         -91.65%
Benchmark_Alter/10000-8            1351197       101083        -92.52%
Benchmark_Alter/12500-8            1743066       109701        -93.71%
Benchmark_Alter/15000-8            2036076       106848        -94.75%
Benchmark_Insert_Small/1000-8      101709        50010         -50.83%
Benchmark_Insert_Small/2500-8      216811        48828         -77.48%
Benchmark_Insert_Small/5000-8      411261        47513         -88.45%
Benchmark_Insert_Small/7500-8      631217        66145         -89.52%
Benchmark_Insert_Small/10000-8     734110        51753         -92.95%
Benchmark_Insert_Small/12500-8     943049        59449         -93.70%
Benchmark_Insert_Small/15000-8     1110480       81927         -92.62%
Benchmark_Reader/100000-8          4700          11775         +150.53%
Benchmark_Reader/125000-8          6133          12889         +110.16%
Benchmark_Reader/150000-8          7688          21242         +176.30%
Benchmark_Reader/175000-8          9279          22491         +142.39%
Benchmark_Reader/200000-8          10849         23263         +114.43%
Benchmark_Remove_Small/1000-8      176937        99681         -43.66%
Benchmark_Remove_Small/2500-8      382343        71238         -81.37%
Benchmark_Remove_Small/5000-8      776255        72214         -90.70%
Benchmark_Remove_Small/7500-8      1177203       98011         -91.67%
Benchmark_Remove_Small/10000-8     1357515       76813         -94.34%
Benchmark_Remove_Small/12500-8     1742483       86207         -95.05%
Benchmark_Remove_Small/15000-8     2047698       103396        -94.95%

Comparing v1 and v3
benchmark                          old ns/op     new ns/op     delta
Benchmark_Alter/1000-8             187914        88435         -52.94%
Benchmark_Alter/2500-8             413970        127396        -69.23%
Benchmark_Alter/5000-8             783128        134560        -82.82%
Benchmark_Alter/7500-8             1163277       103130        -91.13%
Benchmark_Alter/10000-8            1351197       140595        -89.59%
Benchmark_Alter/12500-8            1743066       132609        -92.39%
Benchmark_Alter/15000-8            2036076       108254        -94.68%
Benchmark_Insert_Small/1000-8      101709        54512         -46.40%
Benchmark_Insert_Small/2500-8      216811        76864         -64.55%
Benchmark_Insert_Small/5000-8      411261        79241         -80.73%
Benchmark_Insert_Small/7500-8      631217        81506         -87.09%
Benchmark_Insert_Small/10000-8     734110        82649         -88.74%
Benchmark_Insert_Small/12500-8     943049        83155         -91.18%
Benchmark_Insert_Small/15000-8     1110480       84375         -92.40%
Benchmark_Reader/100000-8          4700          10209         +117.21%
Benchmark_Reader/125000-8          6133          10820         +76.42%
Benchmark_Reader/150000-8          7688          18725         +143.56%
Benchmark_Reader/175000-8          9279          19679         +112.08%
Benchmark_Reader/200000-8          10849         20630         +90.16%
Benchmark_Remove_Small/1000-8      176937        94508         -46.59%
Benchmark_Remove_Small/2500-8      382343        95370         -75.06%
Benchmark_Remove_Small/5000-8      776255        96544         -87.56%
Benchmark_Remove_Small/7500-8      1177203       98426         -91.64%
Benchmark_Remove_Small/10000-8     1357515       100405        -92.60%
Benchmark_Remove_Small/12500-8     1742483       101078        -94.20%
Benchmark_Remove_Small/15000-8     2047698       102076        -95.02%
```
