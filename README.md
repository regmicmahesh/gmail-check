# GMAIL CHECK (SUPERFAST)

This is a super-fast gmail checker implemented in go. The idea was originated from [@codedbrain](https://twitter.com/MRCodedBrain) and I decided to rewrite it in go.


Original Python Implementation: https://github.com/codedbrain/gmailify


# Installation

Make sure you have go installed on your system.


```bash
go version
```

The command above should return some output with description of installed Go version.

Now, you can build the tool.

```bash
go build .
```

Now, you should see an executable called gmail-check on your folder.


# Usage

Reading from a file and writing valid emails into another file.

```bash
./gmail-check -input=input.txt -output=output.txt
```

Reading from stdin and writing into a file.

```bash
./gmail-check -output=output.txt
test@gmail.com
another@gmail.com
```

Reading from file and echoing the valid email in the terminal (not very friendly way to do but ok)

```bash
./gmail-check -input=output.txt
test@gmail.com
another@gmail.com
```
