# words_counter

### The program reads the file, and outputs the 20 most frequently used words in the file in order, along with their frequency.

In accordance with the task constraints the program does not use 'string' and 'map' types.

The program has concurrently running functions that are syncronized with channels

## Usage

To run the program do the following steps:

1. Clone this repo to your machine

2. Navigate to the cloned repo

3. In the root of the project run the command:

```
go run cmd/main.go <fileName>.txt
```

The <fileName> should be replaced with the actual name of the file that you want to analyse.

Example: 

```
go run cmd/main.go mobydick.txt
```

If you want to use the program as a compiled binary do the following:

1. Build the binary file

```
go build -o wordCounter cmd/main.go
```

2. Run the program with the binary file

```
./wordCounter mobydick.txt
```
