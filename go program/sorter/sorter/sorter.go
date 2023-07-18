package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"strconv"
	"io"
	"time"

	"sorter/algorithms/bubblesort"
	"sorter/algorithms/qsort"
)

//func (f *FlagSet) String(name string, value string, usage string) *string
//String用指定的名称、默认值、使用信息注册一个string类型flag。返回一个保存了该flag的值的指针
var infile *string = flag.String("i","infile","File contains values for sorting")
var outfile *string = flag.String("o","outfile","File to receive sorted values")
var algorithm *string = flag.String("a","qsort","Sort algorithm")

func main() {
	//从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =",*infile,"outfile =",*outfile,"algorithm =",*algorithm)
	}
	values, err := readValues(*infile)
	fmt.Println("Read values:",values)
	if err == nil {
		t1 := time.Now()
		switch *algorithm{
			case "qsort":
				qsort.QuickSort(values)
			case "bubblesort":
				bubblesort.BubbleSort(values)
			default:
				fmt.Println("Sorting algorithm", *algorithm,"is either unkonw or unsupported.")

		}
		t2 := time.Now()

		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
		
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}

func readValues(infile string)(values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file",infile)
		return
	}

	defer file.Close()

	//func NewReader(rd io.Reader) *Reader
	//NewReader创建一个具有默认大小缓冲、从r读取的*Reader
	br := bufio.NewReader(file)

	values = make([]int,0)

	for {
		//func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
		line, isPrefix, errl := br.ReadLine()

		if errl != nil {
			if errl != io.EOF {
				err = errl
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line) //转换字符数组为字符串

		value, errl := strconv.Atoi(str)

		if errl != nil {
			err = errl
			return
		}

		values = append(values,value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}