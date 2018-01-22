# Catstamp

Thus named for one of my more common uses of `cat`, notetaking.
The idea is from a program that already exists, but that I couldn't remember what it was called.

The following use of `cat` is a reasonable method for capturing notes.
(Use Control+D (^D) to send an EOF and exit.)

```
$ cat >> somefile.txt
this is a test
oh, yeah.  the test.
oops.  forgot.
^D
$ cat somefile.txt 
this is a test.
oh, yeah.  the test.
oops.  forgot.
$
```

This, though, will also timestamp each captured line.

```
$ catstamp >> somefile.txt
this is a test
oh, yeah.  the test.
oops.  forgot.
^D
$ cat somefile.txt 
2018-01-22 19:24:10	this is a test
2018-01-24 02:03:05	oh, yeah.  the test.
2018-01-24 02:03:06	oops.  forgot.
$
```

## License

MIT License

Copyright (c) 2018 Luke Fritz

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
