# go-bitnuts

[![GoDoc](https://godoc.org/github.com/robbynshaw/go-bitnuts?status.svg)](https://godoc.org/github.com/robbynshaw/go-bitnuts)
[![Coverage Status](https://coveralls.io/repos/github/robbynshaw/go-bitnuts/badge.svg?branch=master)](https://coveralls.io/github/robbynshaw/go-bitnuts?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/robbynshaw/go-bitnuts)](https://goreportcard.com/report/github.com/robbynshaw/go-bitnuts)

> Bit twiddling, swapping, and maneuvering for go

Includes many functions useful if one is operating on a bit matrix.
Some functions include shifting bits in an array to the left or
right, shifting bits in a matrix up or down, etc.

## Warning

This library was yanked out of a test program and needs some TLC if it is
to be used in production. Particularly, many of the operations are not
optimized in the least, and were created simply to get the job done.
