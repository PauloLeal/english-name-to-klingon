# english-name-to-klingon

This application translates an english name to klingon's pIqaD alphabet unicode points.

It also uses [http://stapi.co]() to define the specie of named character. It'll say `Unknown` if it can't determine the specie for any reason.

#### The translation was done using these tables

| ![alt text](https://raw.githubusercontent.com/PauloLeal/english-name-to-klingon/master/resources/translation-table.png "Translation table") | ![alt text](https://raw.githubusercontent.com/PauloLeal/english-name-to-klingon/master/resources/unicode-points.png "Unicode points") |
|---|---|

Any character not in image 1 will cause an error: `invalid input`

#### Examples

```bash
$ ./english-name-to-klingon Nyota
0xF8DB 0xF8E8 0xF8DD 0xF8E3 0xF8D0
Human
```

```bash
$ ./english-name-to-klingon B\'Etor
0xF8D1 0xF8E9 0xF8D4 0xF8E3 0xF8DD 0xF8E1
Klingon
```

```bash
$ ./english-name-to-klingon "Ch'Targh"
0xF8D2 0xF8E9 0xF8E3 0xF8D0 0xF8E1 0xF8D5
Klingon
```

```bash
$ ./english-name-to-klingon Spock
invalid input
```

```bash
$ ./english-name-to-klingon no name
0xF8DB 0xF8DD 0x0020 0xF8DB 0xF8D0 0xF8DA 0xF8D4
Unknown
```

```bash
$ ./english-name-to-klingon bad name
0xF8D1 0xF8D0 0xF8D3 0x0020 0xF8DB 0xF8D0 0xF8DA 0xF8D4
Unknown
```

---

Make sure you have \`go\` to run the application. If you don't, you can download from 

[https://golang.org/dl/]() or [https://github.com/moovweb/gvm]()

If you didn't clone to ```$GOPATH/src/``` do this to create a symbolic link in ```$GOPATH/src/github.com/PauloLeal``` diretory (Linux/Mac only)
     
     ./createLinkToGOPATH.sh

To run the application without compiling

    go run main.go
    
To compile and run the compiled binary

    go build
    ./english-name-to-klingon

To run the tests (I created a script to run tests with coverage reports)

    ./test.sh

### You can also download using \`go get\`

    go get -v github.com/PauloLeal/english-name-to-klingon

Then run from your command line

    english-name-to-klingon <some_character_name>


