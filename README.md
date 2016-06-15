## showcode
Fancy way to show the code.

<br>

#### Usage
Use `-c` code flag to code, `-d` for decode back.

```
showcode -c ./sample/source.txt
```

Will produce the following file, but in `.ppm` format. *PNG and friends aint suitable for this purpose bescause of the compression.*

![output](https://github.com/pvlbzn/showcode/blob/master/sample/output.png "showcode output")

Use `-d` decode flag to decode `.ppm` image back to the code.

<br>

#### Issues
- Bug in column/row algorithm with huge source files (it produces only one row)
- Unicode unfriendly because of `.ppm`