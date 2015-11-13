# :clock4: yon

[![Build Status](https://travis-ci.org/mk2/yon.svg)](https://travis-ci.org/mk2/yon)

## :grey_question: About
- `yon` is the language built by Go.
- It will be intended to make multi-node tool easily.
- Its idea and grammar is basically based on a concatnative language, like forth and factor.
- But `yon` does not have a flexibility most concatnative languages have.

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

## Command line
```sh
# start yon REPL, and enter some codes
$ yon repl
```

## :books: Vocabularies

### :green_book: Prelude vocabulary
Prelude vocabulary contains all basic words (dup, over, rot, each, comparators, and arithmetic operators)

#### Prefix

```
prelude~
```

#### Table

word|stack effect|misc
:--:|:----------:|:--:
dup|(x -- x x)|duplicate the top word on stack

### :blue_book: PsUtil vocabulary
This vocabulary have not been implemented yet.

#### Prefix

```
psutil~
```

#### Table
TBD

### :notebook_with_decorative_cover User vocabulary
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
