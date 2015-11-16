# :clock4: yon

[![Build Status](https://travis-ci.org/mk2/yon.svg)](https://travis-ci.org/mk2/yon)

## About

### Summary
- `yon` is the interpreting language built by Go.
- Its idea and grammar is basically based on a concatnative language, like forth and factor.
- It will be intended to make multi-node tool easily.
- `yon` treats the whole significant literal as the word. (`2` is word, `"test"` is word, `true` is also word.)
- But `yon` does not have a flexibility most concatnative languages have.
- And `yon` is under development.

## How to try?
### Clone
```sh
$ go get -u github.com/mk2/yon
$ cd $GOPATH/src/github.com/mk2/yon
```

### Make
```sh
# automatically restore dependencies (require Godep)
$ make restore-deps
# build yon executable, this brings the `yon` binary in current directory
$ make release
```

### Start REPL
```sh
# Lets enjoy yon!
$ yon repl
```

## :notes: Code

```factor
"test" -- string literal
`test` -- string literal
2123   -- number literal
true   -- bool   literal
false  -- bool   literal
name   -- name   literal
{"test" 1 true} -- array literal
{"version":1 name:yon} -- dictionary literal
[1 dup] -- quoted function literal
```

## :books: Vocabularies

### :green_book: Prelude vocabulary
Prelude vocabulary contains all basic words (dup, over, rot, each, comparators, and arithmetic operators)

#### Prefix

```
prelude~
```

#### Table

word|fully qualified key|stack effect|misc
:--:|:-----------------:|:----------:|:--:
`dup`|`prelude~dup`|`(x -- x x)`|duplicate the top word on stack
`if`|`prelude~if`|`(b x y -- )`|take the bool word b, and execute x or y wheter b is true or false

### :blue_book: PsUtil vocabulary
This vocabulary have not been implemented yet.

#### Prefix

```
psutil~
```

#### Table
TBD

### :notebook_with_decorative_cover: User vocabulary
The vocabulary contains all user defined words.

#### Prefix

```
user~
```

#### Table
There are no initial words.

## :memo: TODOs
- [ ] dictionary literal
- [ ] more stable and usable quoted function
- [ ] comment support
- [ ] explicit function declaration (now only quoted function declaration)
- [ ] execute code from file
- [ ] error handling support
- [ ] better repl interface
- [ ] history recording

## :art: Future TODOs
- [ ] event handling
- [ ] stack-effect annotation
- [ ] serf integration (yon's most valuable feature, if realized)
