# Exercise: wordcount

Write a program called wordcount that reads words from stdin and writes each word and its count to stdout. Hint: `os.Stdin`, `ioutil.ReadAll`, `map[string]int`.

```
$ echo foo bar foo bar foo | wordcount
foo: 3 
bar: 2
```

Bonus:
* do it without buffering the stdin to memory. Hint: `bufio.NewScanner`, `bufio.ScanWords`.