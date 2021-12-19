#!/usr/bin/env perl

use v5.34.0;
use warnings;

open my $fh, ">", "zip-generated.go" or die "failed to open: $!";

print $fh <<'END';
// Code generated by generate-zip.pl; DO NOT EDIT.

package slices

import (
	"fmt"

	"github.com/shogo82148/go-container/tuples"
)
END

for my $n(2..10) {
    say $fh "";
    my $types = join ", ", map { "T$_" } 1..$n;
    my $slice_args = join ", ", map { "s$_ []T$_" } 1..$n;
    my $slices = join ", ", map { "s${_}[i]" } 1..$n;
    say $fh "// Zip$n returns a slice of $n-tuples.";
    say $fh "// All slices must have same length. the lengths are different, Zip$n panics.";
    say $fh "func Zip${n}[$types any]($slice_args) []tuples.Tuple${n}[$types] {";
    say $fh "	l := len(s1)";
    for my $i(2..$n) {
        say $fh "	if len(s$i) != l {";
        say $fh "		panic(fmt.Errorf(\"s$i have different length from s1. len(s$i) = %d, len(s1) = %d\", len(s$i), l))";
        say $fh "	}";
    }
    say $fh "	ret := make([]tuples.Tuple${n}[$types], l)";
    say $fh "	for i := 0; i < l; i++ {";
    say $fh "		ret[i] = tuples.Tuple${n}[$types]{$slices}";
    say $fh "	}";
    say $fh "	return ret";
    say $fh "}";
}

close($fh);
