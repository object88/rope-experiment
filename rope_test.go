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

type alterSetIn struct {
	name          string
	initSize      int
	start         int
	end           int
	injectionSize int
}

type alterSet struct {
	x1        string
	x2        string
	x3        string
	start     int
	end       int
	injection string
}

type alterTestFunc func(t *testing.T, charSet charSet, alterSet *alterSet)

type charSet struct {
	name      string
	generator func(int) string
}

type ropeCreator func(init string) Rope

type stringSize struct {
	size int
}

func Test_Create(t *testing.T) {
	loopTest(t, "Create", func(t *testing.T, charSet charSet, stringSize stringSize) {
		init := charSet.generator(stringSize.size)
		r := create(t, init)

		if r == nil {
			t.Fatal("Got nil")
		}

		if r.Length() != stringSize.size {
			t.Fatalf("Incorrect length: expected %d, got %d", stringSize.size, r.Length())
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

func Test_Alter(t *testing.T) {
	alters := []alterSetIn{
		{"from-leaf-remove-0-9", 100, 0, 10, 0},
		{"from-leaf-remove-5-14", 100, 5, 15, 0},
		{"from-leaf-remove-5-94", 100, 5, 94, 0},
		{"from-leaf-remove-90-99", 100, 90, 100, 0},
		{"from-leaf-shrink-0-9", 100, 0, 10, 5},
		{"from-leaf-shrink-5-14", 100, 5, 15, 5},
		{"from-leaf-shrink-5-94", 100, 5, 94, 5},
		{"from-leaf-shrink-90-99", 100, 90, 100, 5},
		{"from-leaf-replace-0-9", 100, 0, 10, 10},
		{"from-leaf-replace-5-14", 100, 5, 15, 10},
		{"from-leaf-replace-5-94", 100, 5, 94, 90},
		{"from-leaf-replace-90-99", 100, 90, 100, 10},
		{"from-leaf-grow-0-9", 100, 0, 10, 15},
		{"from-leaf-grow-5-14", 100, 5, 15, 15},
		{"from-leaf-grow-5-94", 100, 5, 94, 95},
		{"from-leaf-grow-90-99", 100, 90, 100, 15},
		{"from-node-to-left-remove-0-9", 600, 0, 10, 0},
		{"from-node-to-left-remove-5-14", 600, 5, 15, 0},
		{"from-node-to-left-remove-5-294", 600, 5, 294, 0},
		{"from-node-to-left-remove-290-299", 600, 290, 300, 0},
		{"from-node-to-left-shrink-0-9", 600, 0, 10, 5},
		{"from-node-to-left-shrink-5-14", 600, 5, 15, 5},
		{"from-node-to-left-shrink-5-294", 600, 5, 294, 5},
		{"from-node-to-left-shrink-290-299", 600, 290, 300, 5},
		{"from-node-to-left-replace-0-9", 600, 0, 10, 10},
		{"from-node-to-left-replace-5-14", 600, 5, 15, 10},
		{"from-node-to-left-replace-5-294", 600, 5, 294, 90},
		{"from-node-to-left-replace-290-299", 600, 290, 300, 10},
		{"from-node-to-left-grow-0-9", 600, 0, 10, 15},
		{"from-node-to-left-grow-5-14", 600, 5, 15, 15},
		{"from-node-to-left-grow-5-294", 600, 5, 294, 95},
		{"from-node-to-left-grow-290-299", 600, 290, 300, 15},
		{"from-node-to-right-remove-300-309", 600, 300, 310, 0},
		{"from-node-to-right-remove-305-314", 600, 305, 315, 0},
		{"from-node-to-right-remove-305-594", 600, 305, 594, 0},
		{"from-node-to-right-remove-590-599", 600, 590, 600, 0},
		{"from-node-to-right-shrink-300-309", 600, 300, 310, 5},
		{"from-node-to-right-shrink-305-314", 600, 305, 315, 5},
		{"from-node-to-right-shrink-305-594", 600, 305, 594, 5},
		{"from-node-to-right-shrink-590-599", 600, 590, 600, 5},
		{"from-node-to-right-replace-300-309", 600, 300, 310, 10},
		{"from-node-to-right-replace-305-314", 600, 305, 315, 10},
		{"from-node-to-right-replace-305-594", 600, 305, 594, 90},
		{"from-node-to-right-replace-590-599", 600, 590, 600, 10},
		{"from-node-to-right-grow-300-309", 600, 300, 310, 15},
		{"from-node-to-right-grow-305-314", 600, 305, 315, 15},
		{"from-node-to-right-grow-305-594", 600, 305, 594, 95},
		{"from-node-to-right-grow-590-599", 600, 590, 600, 15},
		{"from-node-to-span-remove", 600, 295, 305, 0},
		{"from-node-to-span-shrink-left", 600, 295, 305, 3},
		{"from-node-to-span-shrink-right", 600, 295, 305, 7},
		{"from-node-to-span-replace", 600, 295, 305, 10},
		{"from-node-to-span-grow", 600, 295, 305, 15},
	}
	loopAlterTest(t, "WAT", alters, func(t *testing.T, charSet charSet, alterSet *alterSet) {
		init := alterSet.x1 + alterSet.x2 + alterSet.x3
		r := create(t, init)

		err := r.Alter(alterSet.start, alterSet.end, alterSet.injection)
		if err != nil {
			t.Fatal(err)
		}

		expectedString := alterSet.x1 + alterSet.injection + alterSet.x3
		expectedLength := utf8.RuneCountInString(expectedString)
		expectedByteLength := len(expectedString)

		if r.Length() != expectedLength {
			t.Fatalf("Incorrect length: expected %d, got %d", expectedLength, r.Length())
		}

		if r.ByteLength() != expectedByteLength {
			t.Fatalf("Incorrect byte length: expected %d, got %d", expectedByteLength, r.ByteLength())
		}

		s := r.String()
		if utf8.RuneCountInString(s) != expectedLength {
			t.Fatalf("Incorrect string length: expected %d, got %d", expectedLength, r.Length())
		}

		if len(s) != expectedByteLength {
			t.Fatalf("Incorrect byte string length: expected %d, got %d", expectedByteLength, r.ByteLength())
		}

		if strings.Compare(expectedString, s) != 0 {
			t.Fatalf("Alter failed:\nExpected:\n'%+q'\nGet:\n'%+q'", expectedString, s)
		}
	})
}

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

func Test_Insert_Small_To_Beginning(t *testing.T) {
	loopTest(t, "Insert-To-Middle", func(t *testing.T, charSet charSet, stringSize stringSize) {
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

func Test_Insert_Small_To_Middle(t *testing.T) {
	loopTest(t, "Insert-To-Middle", func(t *testing.T, charSet charSet, stringSize stringSize) {
		init := charSet.generator(stringSize.size)
		i := utf8.RuneCountInString(init) / 2
		r := create(t, init)

		r.Insert(i, "a")

		if r.Length() != stringSize.size+1 {
			t.Fatalf("Incorrect length: expected %d, got %d", stringSize.size+1, r.Length())
		}

		if r.ByteLength() != len(init)+1 {
			t.Fatalf("Incorrect byte length: expected %d, got %d", len(init)+1, r.ByteLength())
		}

		result := r.String()
		runes := []rune(init)
		expected := string(runes[0:i]) + "a" + string(runes[i:])
		if result != expected {
			t.Fatalf("Insert failed:\nExpected:\n'%s'\nGet:\n'%s'", init, result)
		}
	})
}

func Test_Insert_Small_To_End(t *testing.T) {
	loopTest(t, "Insert-Small-To-End", func(t *testing.T, charSet charSet, stringSize stringSize) {
		init := charSet.generator(stringSize.size)
		r := create(t, init)

		r.Insert(r.Length(), "a")

		if r.Length() != stringSize.size+1 {
			t.Fatalf("Incorrect length: expected %d, got %d", stringSize.size+1, r.Length())
		}

		if r.ByteLength() != len(init)+1 {
			t.Fatalf("Incorrect byte length: expected %d, got %d", len(init)+1, r.ByteLength())
		}

		result := r.String()
		expected := init + "a"
		if result != expected {
			t.Fatalf("Insert failed:\nExpected:\n'%s'\nGet:\n'%s'", init, result)
		}

	})
}

func Test_Insert_Large_To_Beginning(t *testing.T) {
	loopTest(t, "Insert-Large-To-Beginning", func(t *testing.T, charSet charSet, stringSize stringSize) {
		init := charSet.generator(stringSize.size)
		r := create(t, init)

		x := charSet.generator(100)
		r.Insert(0, x)

		if r.Length() != stringSize.size+100 {
			t.Fatalf("Incorrect length: expected %d, got %d", stringSize.size+100, r.Length())
		}

		if r.ByteLength() != len(init)+len(x) {
			t.Fatalf("Incorrect byte length: expected %d, got %d", len(init)+len(x), r.ByteLength())
		}

		result := r.String()
		expected := x + init
		if result != expected {
			t.Fatalf("Insert failed:\nExpected:\n'%+q'\nGet:\n'%+q'", expected, result)
		}
	})
}

func Test_Reader(t *testing.T) {
	loopTest(t, "Reader", func(t *testing.T, charSet charSet, stringSize stringSize) {
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
		i := stringSize.size / 2
		x1 := charSet.generator(i)
		x2 := charSet.generator(1)
		x3 := charSet.generator(i - 1)
		init := x1 + x2 + x3

		r := create(t, init)

		r.Remove(i, i+1)

		if r.Length() != stringSize.size-1 {
			t.Fatalf("Incorrect length: expected %d, got %d", stringSize.size-1, r.Length())
		}

		_, x2Size := utf8.DecodeRuneInString(x2)
		if r.ByteLength() != len(init)-x2Size {
			t.Fatalf("Incorrect byte length: expected %d, got %d", len(init)-x2Size, r.ByteLength())
		}

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
		expected := x1 + x3
		if result != expected {
			t.Fatalf("Remove failed:\nOriginal:\n%q\nExpected:\n%q\nGet:\n%q", init, expected, result)
		}
	})
}

func Benchmark_Alter(b *testing.B) {
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
		testAlter(rc, tc.name, tc.init, b)
	}
}

func Benchmark_Insert_Small(b *testing.B) {
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

func loopAlterTest(t *testing.T, name string, alters []alterSetIn, f alterTestFunc) {
	charSets := []charSet{
		{"ASCII", GenerateASCIIString},
		{"Unicode", GenerateUnicodeString},
	}
	for _, charSet := range charSets {
		for _, alterSetIn := range alters {
			t.Run(fmt.Sprintf("%s-%s", charSet.name, alterSetIn.name), func(t *testing.T) {
				x1 := charSet.generator(alterSetIn.start)
				x2 := charSet.generator(alterSetIn.end - alterSetIn.start)
				x3 := charSet.generator(alterSetIn.initSize - alterSetIn.end)
				injection := charSet.generator(alterSetIn.injectionSize)
				f(t, charSet, &alterSet{x1, x2, x3, alterSetIn.start, alterSetIn.end, injection})
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

func testAlter(creater ropeCreator, basename, init string, b *testing.B) {
	b.Run(basename, func(b *testing.B) {
		b.StopTimer()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var err error
			r := creater(init)

			b.StartTimer()

			for i := 0; i < 50; i++ {
				err = r.Alter(i, i+1, "aaaa")
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
