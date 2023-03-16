package sum_test

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/separovichabel/algorithms_datastructures/algorithms/sum"
	rand "github.com/separovichabel/algorithms_datastructures/observation"
)

func TestThreeNon(t *testing.T) {

	sum := sum.Three([]int{1, 2, 3})

	if sum != 0 {
		t.Error("Expected 0")
	}
}

func TestThreeEightInts(t *testing.T) {
	sum := sum.Three([]int{30, -40, -20, -10, 40, 0, 10, 5})

	if sum != 4 {
		t.Error("Expected 4")
	}
}

func benchmarkThree(len int, b *testing.B) {
	f := func() int {
		return rand.RandPositiveNegative(100, -100)
	}
	input := rand.RandSlice(len, f)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum.Three(input)
	}
}

func BenchmarkThree100(b *testing.B) {
	benchmarkThree(100, b)
}

func BenchmarkThree1000(b *testing.B) {
	benchmarkThree(1000, b)
}

func BenchmarkThree2000(b *testing.B) {
	benchmarkThree(2000, b)
}

func BenchmarkThree4000(b *testing.B) {
	benchmarkThree(4000, b)
}

func BenchmarkThree8000(b *testing.B) {
	benchmarkThree(8000, b)
}

func TestReadData(t *testing.T) {
	data := readData("test_data/rand3.txt")

	if len(data) != 3 {
		t.Error("Expected 3 itens")
	}
}

func readData(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var data []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), " ")

		for _, number := range numbers {
			if number == "" {
				continue
			}

			num, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			data = append(data, num)
		}
	}

	return data
}
