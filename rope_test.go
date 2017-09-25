package ropeExperiment

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"strings"
	"testing"
)

var version = flag.String("ver", "", "version of rope")

type ropeCreator func(init string) Rope

func Test_Insert_To_Beginning(t *testing.T) {
	charSets := []struct {
		name      string
		generator func(int) string
	}{
		{"ASCII", GenerateASCIIString},
		{"Unicode", GenerateUnicodeString},
	}

	stringSizes := []struct {
		size int
	}{
		{100},
		{200},
		{300},
		{400},
		{500},
		{600},
		{700},
		{800},
		{900},
		{1000},
	}

	for _, charSet := range charSets {
		for _, stringSize := range stringSizes {
			t.Run(fmt.Sprintf("%s-Insert-%d", charSet.name, stringSize.size), func(t *testing.T) {
				init := charSet.generator(stringSize.size)
				r := Create(t, init)

				r.Insert(0, "a")

				result := r.String()
				expected := "a" + init
				if result != expected {
					t.Fatalf("Insert failed:\nExpected:\n'%s'\nGet:\n'%s'", init, result)
				}
			})
		}
	}
}

func Test_Insert_To_Middle(t *testing.T) {
	charSets := []struct {
		name      string
		generator func(int) string
	}{
		{"ASCII", GenerateASCIIString},
		{"Unicode", GenerateUnicodeString},
	}

	stringSizes := []struct {
		size int
	}{
		{100},
		{200},
		{300},
		{400},
		{500},
		{600},
		{700},
		{800},
		{900},
		{1000},
	}

	for _, charSet := range charSets {
		for _, stringSize := range stringSizes {
			t.Run(fmt.Sprintf("%s-Insert-%d", charSet.name, stringSize.size), func(t *testing.T) {
				init := charSet.generator(stringSize.size)
				i := len(init) / 2
				r := Create(t, init)

				r.Insert(i, "a")

				result := r.String()
				expected := init[0:i] + "a" + init[i:]
				if result != expected {
					t.Fatalf("Insert failed:\nExpected:\n'%s'\nGet:\n'%s'", init, result)
				}
			})
		}
	}
}

func Test_Reader(t *testing.T) {
	charSets := []struct {
		name      string
		generator func(int) string
	}{
		{"ASCII", GenerateASCIIString},
		{"Unicode", GenerateUnicodeString},
	}

	stringSizes := []struct {
		size int
	}{
		{100},
		{200},
		{300},
		{400},
		{500},
		{600},
		{700},
		{800},
		{900},
		{1000},
	}

	for _, charSet := range charSets {
		for _, stringSize := range stringSizes {
			t.Run(fmt.Sprintf("%s-Reader-%d", charSet.name, stringSize.size), func(t *testing.T) {
				init := charSet.generator(stringSize.size)

				var buf bytes.Buffer
				buf.Grow(len(init))

				r := Create(t, init)
				reader := r.NewReader()

				io.Copy(&buf, reader)

				result := string(buf.Bytes())
				if strings.Compare(result, init) != 0 {
					t.Fatalf("Read failed:\nExpected:\n'%s'\nGot:\n'%s'", init, result)
				}
			})
		}
	}
}

func Test_Remove_From_Beginning(t *testing.T) {
	charSets := []struct {
		name      string
		generator func(int) string
	}{
		{"ASCII", GenerateASCIIString},
		{"Unicode", GenerateUnicodeString},
	}

	stringSizes := []struct {
		size int
	}{
		{100},
		{200},
		{300},
		{400},
		{500},
		{600},
		{700},
		{800},
		{900},
		{1000},
	}

	for _, charSet := range charSets {
		for _, stringSize := range stringSizes {
			t.Run(fmt.Sprintf("%s-Remove-%d", charSet.name, stringSize.size), func(t *testing.T) {
				init := charSet.generator(stringSize.size)
				r := Create(t, init)

				r.Remove(0, 1)

				result := r.String()
				expected := init[1:]
				if result != expected {
					t.Fatalf("Remove failed:\nExpected:\n'%s'\nGet:\n'%s'", init, result)
				}
			})
		}
	}
}

func Test_Remove_From_Middle(t *testing.T) {
	charSets := []struct {
		name      string
		generator func(int) string
	}{
		{"ASCII", GenerateASCIIString},
		{"Unicode", GenerateUnicodeString},
	}

	stringSizes := []struct {
		size int
	}{
		{100},
		{200},
		{300},
		{400},
		{500},
		{600},
		{700},
		{800},
		{900},
		{1000},
	}

	for _, charSet := range charSets {
		for _, stringSize := range stringSizes {
			t.Run(fmt.Sprintf("%s-Remove-%d", charSet.name, stringSize.size), func(t *testing.T) {
				init := charSet.generator(stringSize.size)
				i := len(init) / 2
				r := Create(t, init)

				r.Remove(i, i+1)

				result := r.String()
				expected := init[0:i] + init[i+1:]
				if result != expected {
					t.Fatalf("Remove failed:\nExpected:\n'%s'\nGet:\n'%s'", init, result)
				}
			})
		}
	}
}

func Benchmark_Add_Small(b *testing.B) {
	tests := []struct {
		name string
		init string
	}{
		{"1000", GenerateASCIIString(1000)},
		{"2500", GenerateASCIIString(2500)},
		{"5000", GenerateASCIIString(5000)},
		{"7500", GenerateASCIIString(7500)},
		{"10000", GenerateASCIIString(10000)},
		{"12500", GenerateASCIIString(12500)},
		{"15000", GenerateASCIIString(15000)},
	}

	rc, err := getCreater()
	if nil != err {
		b.Error(err)
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
		{"100000", GenerateASCIIString(100000)},
		{"125000", GenerateASCIIString(125000)},
		{"150000", GenerateASCIIString(150000)},
		{"175000", GenerateASCIIString(175000)},
		{"200000", GenerateASCIIString(200000)},
	}

	rc, err := getCreater()
	if nil != err {
		b.Error(err)
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
		{"1000", GenerateASCIIString(1000)},
		{"2500", GenerateASCIIString(2500)},
		{"5000", GenerateASCIIString(5000)},
		{"7500", GenerateASCIIString(7500)},
		{"10000", GenerateASCIIString(10000)},
		{"12500", GenerateASCIIString(12500)},
		{"15000", GenerateASCIIString(15000)},
	}

	rc, err := getCreater()
	if nil != err {
		b.Error(err)
	}

	for _, tc := range tests {
		testRemove(rc, tc.name, tc.init, b)
	}
}

func getCreater() (ropeCreator, error) {
	var rc ropeCreator
	if *version == "1" {
		rc = func(init string) Rope { return CreateV1(init) }
	} else if *version == "2" {
		rc = func(init string) Rope { return CreateV2(init) }
	} else if *version == "3" {
		rc = func(init string) Rope { return CreateV3(init) }
	} else {
		return nil, fmt.Errorf("'%s' is not a valid version", *version)
	}

	return rc, nil
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

func Create(t *testing.T, init string) Rope {
	rc, err := getCreater()
	if err != nil {
		t.Fatal(err)
		return nil
	}
	r := rc(init)
	return r
}
