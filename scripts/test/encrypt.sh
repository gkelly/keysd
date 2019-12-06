#!/usr/bin/env bash

set -e -u -o pipefail # Fail on error

tmpfile=`mktemp /tmp/XXXXXXXXXXX`
tmpfile2=`mktemp /tmp/XXXXXXXXXXX`
tmpfile3=`mktemp /tmp/XXXXXXXXXXX`

head -c 500000 </dev/urandom > "$tmpfile"

encfile=`mktemp /tmp/XXXXXXXXXXX`
encfile2=`mktemp /tmp/XXXXXXXXXXX`

kid=`keys list | head -1 | cut -d ' ' -f 1`

echo "encrypt"
keys encrypt -recipients $kid -in "$tmpfile" -out "$encfile"
echo "decrypt"
keys decrypt -in "$encfile" -out "$tmpfile2"
diff "$tmpfile2" "$tmpfile2"

echo "encrypt"
cat "$tmpfile2" | keys encrypt -recipients $kid > "$encfile2"
echo "decrypt"
cat "$encfile2" | keys decrypt > "$tmpfile3"
diff "$tmpfile2" "$tmpfile3"