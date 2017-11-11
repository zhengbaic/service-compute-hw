package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type selpg_args struct {
	start_page  int
	end_page    int
	in_filename string
	print_dest  string
	page_len    int
	f_type      bool
}

var progname string

func errPrint(err string) {
	fmt.Fprintln(os.Stderr, err)
	flag.Usage()
	os.Exit(1)
}
func Usage() {
	fmt.Printf("\nUSAGE: %s -s=start_page -e=end_page [ -f | -l=lines_per_page ][filename]\n", progname)
}

func process_args(sa *selpg_args) {
	if len(os.Args) < 3 {
		errPrint(progname + " : not enough arguments\n")
	}

	if os.Args[1][0] != '-' || os.Args[1][1] != 's' {
		errPrint(progname + " : 1st arg should be -s=start_page\n")
	}
	if os.Args[2][0] != '-' || os.Args[2][1] != 'e' {
		errPrint(progname + " : 2nd arg should be -e=end_page\n")
	}
	if sa.start_page > sa.end_page || sa.start_page < 0 || sa.end_page < 0 {
		errPrint(progname + " : your arguments are invalid, please check\n")
	}
	if sa.f_type == true && sa.page_len != -1 {
		errPrint(progname + " : you can't set the -l and -f at the same time\n")
	}
	if sa.f_type == false {
		if sa.page_len < 0 {
			errPrint(progname + " : the length of page can not < 0\n")
		} else if sa.page_len == -1 {
			sa.page_len = 72
		}
	}
}
func process_input(sa *selpg_args) {
	var stdin io.WriteCloser
	var err error
	var cmd *exec.Cmd
	if sa.print_dest != "" {
		cmd = exec.Command("cat", "-n")
		stdin, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		stdin = os.Stdout
	}
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
	if sa.print_dest != "" {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func type_f_process(sa *selpg_args, reader *bufio.Reader, stdin io.WriteCloser) {
	for pageNum := 0; pageNum <= sa.end_page; pageNum++ {
		Output, err := reader.ReadString('\f')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if pageNum >= sa.start_page {
			OutputProcess(sa, string(Output), stdin)
		}
	}
}
func type_l_process(sa *selpg_args, reader *bufio.Reader, stdin io.WriteCloser) {
	num := 0
	for {
		Output, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if num/sa.page_len >= sa.start_page {
			if num/sa.page_len <= sa.end_page {
				OutputProcess(sa, string(Output), stdin)
			} else {
				break
			}
		}
		num++
	}
}
func ReadFromStdin(sa *selpg_args, stdin io.WriteCloser) {
	StdinInput := bufio.NewScanner(os.Stdin)
	num := 0
	Output := ""
	temp := ""
	for StdinInput.Scan() {
		temp = StdinInput.Text() + "\n"
		if num/sa.page_len >= sa.start_page && num/sa.page_len < sa.end_page {
			Output += temp
		}
		num++
	}
	OutputProcess(sa, string(Output), stdin)
}

func OutputProcess(sa *selpg_args, Output string, stdin io.WriteCloser) {
	if sa.print_dest != "" {
		stdin.Write([]byte(Output + "\n"))
	} else {
		fmt.Println(Output)
	}
}
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
