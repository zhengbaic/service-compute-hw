# selpg
[selpg](https://github.com/zhengbaic/selpg) is a simple CLI implemented by GO

## wirte by writer
 I run it in Windows, so -d options is undebugable, and this part, I learn it from QQ group.

## Usage

```shell
Usage of selpg:
  -s int
  		start page(mandatory)(default -1)
  -e int
  		end page(mandatory)(default -1)
  -f bool
  		use \f to seperate pages
  -l int
  		length of page to seperate pages
  -d string
    	a destination command that receives the output, if it's not specified, output is printed to stdout)
```

## How I implement it

### a struct of selpg
```go
type selpg_args struct {
	start_page  int
	end_page    int
	in_filename string
	print_dest  string
	page_len    int
	f_type      bool
}
```

### the main process
```go
func main() {
	var sa selpg_args
	progname = os.Args[0]

	flag.Usage = Usage
	flag.IntVar(&sa.start_page, "s", -1, "sepcify start_page.")
	flag.IntVar(&sa.end_page, "e", -1, "sepcify end_page.")
	flag.IntVar(&sa.page_len, "l", -1, "specify page_lenth")
	flag.BoolVar(&sa.f_type, "f", false, "specify if using delimiter")
	flag.StringVar(&sa.print_dest, "d", "", "specify the des_ouput")
	flag.Parse()

	process_args(&sa)
	process_input(&sa)
}
```

### process_args
where i parse the input and panic the errors, it goes like:
```go
	if len(os.Args) < 3 {
		errPrint(progname + " : not enough arguments\n")
	}
```

### process_input
where i choose the output by different arg, and it goes like:
```go
	if flag.NArg() == 0 {
		fmt.Printf("\nUSAGE: %s Now the input will read from stdin, press (Ctrl+D)(Linux) | (Ctrl+C)(Windows) to exit\n", progname)
		ReadFromStdin(sa, stdin)
	} else {
		sa.in_filename = flag.Arg(0)
		input_stream, err := os.Open(sa.in_filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		input_buffer := bufio.NewReader(input_stream)
		if sa.f_type {
			type_f_process(sa, input_buffer, stdin)
		} else {
			type_l_process(sa, input_buffer, stdin)
		}
	}
```

## Example
the 1.txt is :
```shell
Line1
Line2
Line3
Line4
Line5
Line6
Line7
Line8
Line9
Line10
```

print 1.txt by -l:
```shell
>selpg.exe -s=1 -e=3 -l=1 1.txt
Line2
Line3
Line4
```

use terminal to input
```shell
>selpg.exe -s=1 -e=3 -l=1

USAGE: selpg.exe Now the input will read from stdin, press (Ctrl+D)(Linux) | (Ctrl+C)(Windows) to exit
q
q
2
2
4
4
```

As fot the rest, u can try.

#Bye 