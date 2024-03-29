# otr-keychain

    go get github.com/juniorz/otr-keychain

This package extends `golang.org/x/crypto/otr` by allowing to import multiple
OTR private keys from the same file.

## Context

Althought the following libotr private key file (`otr.private_key`) is valid[1],

    (privkeys
     (account
    (name "1")
    (protocol libpurple-Jabber)
    (private-key 
     (dsa 
      (p #008D3EE66820824BDB81FCFEE853A2148C600D22C4D5D66ED7535A2F1F890A130239AFC3619FE2B5886EE2BEE109170D8D56B79AEB2217D490F92AA5A82F1B562CB19E51F3C38C717B4A0AA0306A9CBDF62977566AE041F3E2D60CE02C82ED9F350A38A016E9CB3C93C4E5F9A8B14ED738E9CB0D076EEB02009D0CBBF7AF36370B#)
      (q #00A7C84E6CB16B190F2FC962EDA3D77926490F902F#)
      (g #707219B8AE0783C234DF2426AA6E69F30738AC112FB0B53CC54776959AAF67626BB895E18AA8B735DD3E19BE780B5B7193DB2D64159B4757B0CB23AC555B70F25AF1474FE4D9C44059F8BFCA1BEBC845C4696A558ABA0C916368D3961EEBB8130E4651B618384E687CF7A9917637D97BE51D44217937F7D8A9A2C9ED850B0BAC#)
      (y #329D74BDCA8399977150519EDBC018988D43393834F77E36D577772B9434B2EE08B572BDBB7F4A23F5242FFCDBD90A4A8D405C8366CAB785455B431B888DB412BC37713648D6F5C919E501E949BA8B719BE3EF31B6244D483BCCAE042C1E6E6DF462E5A5064BED031A55BC1E3DAAF38895A1B92A20C6D82D78F90C24CABBC566#)
      (x #7A477B8980503487CEE95628EBAC52A19274BD57#)
      )
     )
     )
     (account
    (name "2")
    (protocol libpurple-jabber-gtalk)
    (private-key 
     (dsa 
      (p #00F24843F9447B62138AE49BF83188D1353ADA5CAC118890CFDEC01BF349D75E887B19C221665C7857CAD583AF656C67FB04A99FD8F8D69D09C9529C6C14D426F1E3924DC9243AF2970E3E4B04A23489A09E8A90E7E81EBA763AD4F0636B8A43415B6FC16A02C3624CE76272FA00783C8DB850D3A996B58136F7A0EB80AE0BC613#)
      (q #00D16B2607FCBC0EDC639F763A54F34475B1CC8473#)
      (g #00B15AFEF5F96EFEE41006F136C23A18849DA8133069A879D083F7C7AA362E187DAE3ED0C4F372D0D4E3AAE567008A1872A6E85D8F84E53A3FE1B352AF0B4E2F0CB033A6D34285ECD3E4A93653BDE99C3A8D840D9D35F82AC2FA8539DB6C7F7A1DAD77FEECD62803757FF1E2DE4CEC4A5A2AD643271514DDEEEF3D008F66FBF9DB#)
      (y #01F9BE7DA0E4E84774048058B53202B2704BF688A306092ED533A55E68EABA814C8D62F45AAD8FF30C3055DCA461B7DBA6B78938FC4D69780A830C6457CC107F3D275C21D00E53147C14162176C77169D3BCA586DC30F15F4B482160E276869AA336F38AF7FC3686A764AB5A02C751D921A42B8B9AE8E06918059CD73C424154#)
      (x #00943480B228FC0D3D7ADFC91F680FC415E1306333#)
      )
     )
     )
    )

the current otr implementation available at `golang.org/x/crypto/otr`'s
[PrivateKey.Import](https://godoc.org/golang.org/x/crypto/otr#PrivateKey.Import)
only process the first record.

1 - See `test_suite/otr.private_key` in git://git.otr.im/libotr.git
