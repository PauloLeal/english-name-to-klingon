# english-name-to-klingon

This application translates an english name to klingon using hexadecimal numbers defined by a character table.

It also uses stapi.co to define the race of named character

Make sure you have \`go\` to run the game. If you don't, you can download from 

    https://github.com/moovweb/gvm


Do this if you didn't clone to your $GOPATH/src/ diretory (Linux/Mac only)
     
     ./createLinkToGOPATH.sh #creates a symlink to project

To run the application without compiling

    go run main.go
    
To compile and run the compiled binary

    go build
    ./english-name-to-klingon

To run the tests (I created a script to run tests with coverage reports)

    ./test.sh

You can also download using go get

    go get -v github.com/PauloLeal/english-name-to-klingon

Then run from your command line

    english-name-to-klingon


