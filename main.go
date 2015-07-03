package keychain

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s /path/to/otr.private_key\n\n", os.Args[0])
		fmt.Fprint(os.Stderr, "  Adium:   ~/Library/Application Support/Adium 2.0/Users/Default/otr.private_key\n")
		fmt.Fprint(os.Stderr, "  Pidgin:  ~/.purple/otr.private_key\n")
		os.Exit(1)
	}

	path := os.Args[1]
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open OTR private key at %q\n", path)
		os.Exit(1)
	}

	keys := ImportFromLibOTR(data)

	if len(keys) == 0 {
		fmt.Fprintf(os.Stderr, "Failed to import OTR private key at %q\n", path)
		os.Exit(1)
	}

	fmt.Printf("%d key(s) found in %s\n", len(keys), path)

	for _, k := range keys {
		fmt.Printf(" - key fingerprint:  %s\n", hex.EncodeToString(k.Fingerprint()))
	}
}