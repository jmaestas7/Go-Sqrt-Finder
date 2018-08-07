# Go-Sqrt-Finder
A Simple CLI program that finds the square root of any flagged argument inputted using <b>"number"</b>

If not using the "main.exe", use the go build function to build the program.
``` Go
$ go build main.go
```
After the prgram has been built you can then optionally input a flagged <b>"number"</b> or use the defalt number 2
``` Bash
$ ./main
1.4142135623730951 is the solution found
exit
```
``` Bash
$ ./main -number=169
13 is the solution found
exit
```
This program uses an iteration of 
```
z -= (z*z - x) / (2*z)
```
where <b>z =1</b> and iterates till the soulition is found to match the square root of the inputted or default number
``` Go
for {
  z -= (z*z - x) / (2 * z)
  if z == math.Sqrt(x) {
    return z, "is the solution found"
  }
}
 ```
