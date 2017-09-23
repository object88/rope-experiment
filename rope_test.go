package ropeExperiment

import (
	"bytes"
	"flag"
	"io"
	"strings"
	"testing"
)

var version = flag.String("ver", "", "version of rope")

type ropeCreator func(init string) Rope

func Benchmark_Add_Small(b *testing.B) {
	tests := []struct {
		name string
		init string
	}{
		{"500", strings.Repeat("0123456789", 50)},
		{"750", strings.Repeat("0123456789", 75)},
		{"1000", strings.Repeat("0123456789", 100)},
		{"2500", strings.Repeat("0123456789", 250)},
		{"5000", strings.Repeat("0123456789", 500)},
		{"7500", strings.Repeat("0123456789", 750)},
		{"10000", strings.Repeat("0123456789", 1000)},
		{"12500", strings.Repeat("0123456789", 1250)},
		{"15000", strings.Repeat("0123456789", 1500)},
		{"17500", strings.Repeat("0123456789", 1750)},
		{"20000", strings.Repeat("0123456789", 2000)},
	}

	var rc ropeCreator
	if *version == "1" {
		rc = func(init string) Rope { return CreateV1(init) }
	} else if *version == "2" {
		rc = func(init string) Rope { return CreateV2(init) }
	} else {
		b.Fatalf("'%s' is not a valid version", *version)
	}

	for _, tc := range tests {
		testAdd(rc, tc.name, tc.init, b)
	}
}

func Test_Reader(t *testing.T) {
	init := RandStringBytesMaskImprSrc(1000)

	var buf bytes.Buffer
	buf.Grow(len(init))

	r := CreateV2(init)
	reader := r.NewReader()

	io.Copy(&buf, reader)

	result := string(buf.Bytes())
	if strings.Compare(result, init) != 0 {
		t.Fatalf("Read failed:\nExpected:\n'%s'\nGot:\n'%s'", init, result)
	}
}

func Benchmark_Reader(b *testing.B) {
	tests := []struct {
		name string
		init string
	}{
		{"100000", RandStringBytesMaskImprSrc(100000)},
		{"125000", RandStringBytesMaskImprSrc(125000)},
		{"150000", RandStringBytesMaskImprSrc(150000)},
		{"175000", RandStringBytesMaskImprSrc(175000)},
		{"200000", RandStringBytesMaskImprSrc(200000)},
	}

	var rc ropeCreator
	if *version == "1" {
		rc = func(init string) Rope { return CreateV1(init) }
	} else if *version == "2" {
		rc = func(init string) Rope { return CreateV2(init) }
	} else {
		b.Fatalf("'%s' is not a valid version", *version)
	}

	for _, tc := range tests {
		testReader(rc, tc.name, tc.init, b)
	}
}

func testAdd(creater ropeCreator, basename, init string, b *testing.B) {
	var err error
	b.Run(basename, func(b *testing.B) {
		rs := make([]Rope, b.N)
		for i := 0; i < b.N; i++ {
			rs[i] = creater(init)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			err = rs[i].Insert(0, "a")
		}
	})
	if err != nil {
		b.Fatal("Error during tests.")
	}
}

func testReader(creater ropeCreator, basename, init string, b *testing.B) {
	b.Run(basename, func(b *testing.B) {
		r := creater(init)

		b.StopTimer()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			buf.Grow(len(init))

			reader := r.NewReader()

			b.StartTimer()

			io.Copy(&buf, reader)

			b.StopTimer()

			if i == 0 {
				result := string(buf.Bytes())
				if strings.Compare(result, init) != 0 {
					b.Fatalf("Read failed:\nExpected:\n'%s'\nGot:\n'%s'", init, result)
				}
			}
		}
	})
}
