# otr-keychain

    go get github.com/juniorz/otr-keychain

Improves `golang.org/x/crypto/otr` by allowing to import multiple OTR private
keys from the same file.

This format is currently supported by `libotr`[1] and
[guardianproject/keysync](https://github.com/guardianproject/keysync/tree/master/tests/adium).

1 - See `test_suite/otr.private_key` in git://git.otr.im/libotr.git
