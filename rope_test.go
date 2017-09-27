package ropeExperiment

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"strings"
	"testing"
	"unicode/utf8"
)

var version = flag.String("ver", "", "version of rope")

type ropeCreator func(init string) Rope

func Test_Insert(t *testing.T) {
	initial := "🐿🐿🐿🐿🐿"
	r := create(t, initial)
	r.Insert(1, "a")

	actual := r.String()
	expected := "🐿a🐿🐿🐿🐿"

	if expected != actual {
		t.Fatalf("Failed to properly insert:\nexpected %s\ngot %s\n", expected, actual)
	}
}

func Test_Create(t *testing.T) {
	charSets := []struct {
		name      string
		generator func(int) string
	}{
		{"ASCII", GenerateASCIIString},
		{"Unicode", GenerateUnicodeString},
	}

	testCases := []struct {
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
		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%s-Create-%d", charSet.name, tc.size), func(t *testing.T) {
				init := charSet.generator(tc.size)
				r := create(t, init)

				if r == nil {
					t.Fatal("Got nil")
				}

				if r.Length() != tc.size {
					t.Fatalf("Incorrect length: expected %d, got %d", tc.size, r.Length())
				}

				if r.ByteLength() != len(init) {
					t.Fatalf("Incorrect byte length: expected %d, got %d", len(init), r.ByteLength())
				}

				actual := r.String()
				if actual != init {
					t.Fatalf("Did not get same string back.\nexpected:\n%+q\ngot:\n%+q\n", init, actual)
				}
			})
		}
	}
}

func Test_Insert_Small_To_Beginning(t *testing.T) {
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
				r := create(t, init)

				r.Insert(0, "a")

				if r.Length() != stringSize.size+1 {
					t.Fatalf("Incorrect length: expected %d, got %d", stringSize.size+1, r.Length())
				}

				if r.ByteLength() != len(init)+1 {
					t.Fatalf("Incorrect byte length: expected %d, got %d", len(init)+1, r.ByteLength())
				}

				result := r.String()
				expected := "a" + init
				if result != expected {
					t.Fatalf("Insert failed:\nExpected:\n'%+q'\nGet:\n'%+q'", expected, result)
				}
			})
		}
	}
}

func Test_Insert_Small_To_Middle(t *testing.T) {
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
				i := utf8.RuneCountInString(init) / 2
				r := create(t, init)

				r.Insert(i, "a")

				result := r.String()
				runes := []rune(init)
				expected := string(runes[0:i]) + "a" + string(runes[i:])
				if result != expected {
					t.Fatalf("Insert failed:\nExpected:\n'%s'\nGet:\n'%s'", init, result)
				}
			})
		}
	}
}

func Test_Insert_Large_To_Beginning(t *testing.T) {
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
			t.Run(fmt.Sprintf("%s-Insert-Large-%d", charSet.name, stringSize.size), func(t *testing.T) {
				init := charSet.generator(stringSize.size)
				r := create(t, init)

				x := charSet.generator(100)
				r.Insert(0, x)

				result := r.String()
				expected := x + init
				if result != expected {
					t.Fatalf("Insert failed:\nExpected:\n'%+q'\nGet:\n'%+q'", expected, result)
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

				r := create(t, init)
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

func Test_Remove_Small_From_Beginning(t *testing.T) {
	loopTest(t, "Remove-From-Beginning", func(t *testing.T, charSet charSet, stringSize stringSize) {
		init := charSet.generator(stringSize.size)
		r := create(t, init)

		r.Remove(0, 1)

		if r.Length() != stringSize.size-1 {
			t.Fatalf("Incorrect length: expected %d, got %d", stringSize.size-1, r.Length())
		}

		_, rSize := utf8.DecodeRuneInString(init)
		if r.ByteLength() != len(init)-rSize {
			t.Fatalf("Incorrect byte length: expected %d, got %d", len(init)-rSize, r.ByteLength())
		}

		result := r.String()
		if !utf8.ValidString(result) {
			t.Fatal("Invalid UTF8 string")
		}
		expected := string([]rune(init)[1:])
		if result != expected {
			t.Fatalf("Remove failed:\nOriginal:\n%q\nExpected:\n%q\nGet:\n%q", init, expected, result)
		}
	})
}

func Test_Remove_Small_From_Middle(t *testing.T) {
	loopTest(t, "Remove-From-Middle", func(t *testing.T, charSet charSet, stringSize stringSize) {
		init := charSet.generator(stringSize.size)
		i := utf8.RuneCountInString(init) / 2
		r := create(t, init)

		r.Remove(i, i+1)

		result := r.String()
		if !utf8.ValidString(result) {
			b := []byte(result)
			for i := 0; i < len(result); {
				ru, n := utf8.DecodeRune(b)
				if ru == utf8.RuneError {
					t.Fatalf("Invalid UTF8 string; first instance at %d\n%s", i, result)
				}
				i += n
			}
			t.Fatal("Invalid UTF8 string")
		}
		expected := string([]rune(init)[0:i]) + string([]rune(init)[i+1:])
		if result != expected {
			t.Fatalf("Remove failed:\nOriginal:\n%q\nExpected:\n%q\nGet:\n%q", init, expected, result)
		}
	})
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

func create(t *testing.T, init string) Rope {
	rc, err := getCreater()
	if err != nil {
		t.Fatal(err)
		return nil
	}
	r := rc(init)
	return r
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

type charSet struct {
	name      string
	generator func(int) string
}

type stringSize struct {
	size int
}

func loopTest(t *testing.T, name string, f func(t *testing.T, charSet charSet, stringSize stringSize)) {
	charSets := []charSet{
		{"ASCII", GenerateASCIIString},
		{"Unicode", GenerateUnicodeString},
	}

	stringSizes := []stringSize{
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
			t.Run(fmt.Sprintf("%s-%s-%d", charSet.name, name, stringSize.size), func(t *testing.T) {
				f(t, charSet, stringSize)
			})
		}
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
