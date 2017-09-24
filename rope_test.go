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

func Test_Add(t *testing.T) {
	init := RandStringBytesMaskImprSrc(1000)
	r := CreateV2(init)

	r.Insert(0, "a")

	result := r.String()
	expected := "a" + init
	if result != expected {
		t.Fatalf("Insert failed:\nExpected:\n'%s'\nGet:\n'%s'", init, result)
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

func Test_Remove(t *testing.T) {
	init := RandStringBytesMaskImprSrc(1000)
	r := CreateV2(init)

	r.Remove(0, 1)

	result := r.String()
	expected := init[1:]
	if result != expected {
		t.Fatalf("Remove failed:\nExpected:\n'%s'\nGet:\n'%s'", init, result)
	}
}

func Benchmark_Add_Small(b *testing.B) {
	tests := []struct {
		name string
		init string
	}{
		{"1000", RandStringBytesMaskImprSrc(1000)},
		{"2500", RandStringBytesMaskImprSrc(2500)},
		{"5000", RandStringBytesMaskImprSrc(5000)},
		{"7500", RandStringBytesMaskImprSrc(7500)},
		{"10000", RandStringBytesMaskImprSrc(10000)},
		{"12500", RandStringBytesMaskImprSrc(12500)},
		{"15000", RandStringBytesMaskImprSrc(15000)},
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

func Benchmark_Remove_Small(b *testing.B) {
	tests := []struct {
		name string
		init string
	}{
		{"1000", RandStringBytesMaskImprSrc(1000)},
		{"2500", RandStringBytesMaskImprSrc(2500)},
		{"5000", RandStringBytesMaskImprSrc(5000)},
		{"7500", RandStringBytesMaskImprSrc(7500)},
		{"10000", RandStringBytesMaskImprSrc(10000)},
		{"12500", RandStringBytesMaskImprSrc(12500)},
		{"15000", RandStringBytesMaskImprSrc(15000)},
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
		testRemove(rc, tc.name, tc.init, b)
	}
}

func testAdd(creater ropeCreator, basename, init string, b *testing.B) {
	b.Run(basename, func(b *testing.B) {
		b.StopTimer()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var err error
			r := creater(init)

			b.StartTimer()

			for i := 0; i < 50; i++ {
				err = r.Insert(i, "a")
			}

			b.StopTimer()

			if err != nil {
				b.Fatal("Error during tests.")
			}
		}
	})
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

func testRemove(creater ropeCreator, basename, init string, b *testing.B) {
	b.Run(basename, func(b *testing.B) {
		b.StopTimer()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var err error
			r := creater(init)

			b.StartTimer()

			for i := 0; i < 50; i++ {
				err = r.Remove(i, i+1)
			}

			b.StopTimer()

			if err != nil {
				b.Fatal("Error during tests.")
			}
		}
	})
}
