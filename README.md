# goop, Go Object Print

This is a blatant copy of Rust oDB, a printer you use for debugging, sorta.

Remember writing printf incantations for structs, embedding them in your source, 
using them based on a debug flag, and then throwing them away?

This is a program that calls an http endpoint that returns a stuct in json, so the program can pretty-print it.

The interface with the program is trivial and localized, so you don't need to delet it and then reinvent it every time.


## Synopsis
goop [-opts] url command

## Description
Sends a command to a program via http, to print a named thing. Pretty-prints the returned json.

## Files

## See Also

## Bugs

The real oDB can do iterations: I'm not doing that for the MVP. They're easy to add.

## Diagnostics

## Compatability

## Changes

## Acknowlegements

oDB, by Dave Pacheco

## Author

Dave Collier-Brown
