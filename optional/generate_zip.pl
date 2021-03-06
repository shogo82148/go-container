#!/usr/bin/env perl

use v5.34.0;
use warnings;

open my $fh, ">", "zip_generated.go" or die "failed to open: $!";

print $fh <<'END';
// Code generated by generate_zip.pl; DO NOT EDIT.

package optional

import (
	"github.com/shogo82148/go-container/tuples"
)
END

for my $n(1..64) {
    say $fh "";
    my $types = join ", ", map { "T$_" } 1..$n;
    my $slice_args = join ", ", map { "v$_ Optional[T$_]" } 1..$n;
    say $fh "// Zip$n returns an optional of $n-tuples.";
    say $fh "func Zip${n}[$types any]($slice_args) Optional[tuples.Tuple${n}[$types]] {";
    say $fh "\tif ", join(" && ", map { "v$_.valid" } 1..$n ), " {";
    say $fh "\t\treturn Optional[tuples.Tuple${n}[$types]]{";
    say $fh "\t\t\tvalue: tuples.Tuple${n}[$types]{", join(", ", map { "V$_: v$_.value" } 1..$n ), "},";
    say $fh "\t\t\tvalid: true,";
    say $fh "\t\t}";
    say $fh "\t}";
    say $fh "\treturn Optional[tuples.Tuple${n}[$types]]{}";
    say $fh "}";
}

close($fh);
