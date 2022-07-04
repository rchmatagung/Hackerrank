package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    s = strings.ToLower(s)
    pm := strings.HasSuffix(s, "pm")
    am := strings.HasSuffix(s, "am")
    
    t := s[:len(s)-2]
    
    timeArr := strings.Split(t, ":")
    hh, mm, ss := timeArr[0], timeArr[1], timeArr[2]
    hhInt, err := strconv.Atoi(hh)
    
    if err != nil {
        panic(err.Error())
    }
    
    if am && hhInt == 12 {
        hhInt = 0
    }
    
    if pm && hhInt != 12 {
        hhInt += 12
    }
    
    hh = strconv.Itoa(hhInt)
    return fmt.Sprintf("%02s:%02s:%02s", hh, mm, ss)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
