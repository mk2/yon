# yon

[![Build Status](https://travis-ci.org/mk2/yon.svg)](https://travis-ci.org/mk2/yon)

## About
- `yon` is the language built by Go.
- It will be intended to make multi-node tool easily.
- Its idea and grammar is basically based on a concatnative language, like forth and factor.
- But `yon` does not have a flexibility most concatnative languages have.

## Code

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

## TODOs
- [ ] dictionary literal
- [ ] more stable and usable quoted function
- [ ] explicit function declaration (now only quoted function declaration)
- [ ] execute code from file
- [ ] error handling support
- [ ] better repl interface
- [ ] history recording

## Future TODOs
- [ ] event handling
- [ ] serf integration (yon's most valuable feature, if realized)
