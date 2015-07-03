package keychain

import (
	"reflect"
	"testing"
)

var otrPrivateKeyFile = `
(privkeys
 (account
(name otrtest1)
(protocol prpl-aim)
(private-key 
 (dsa 
  (p #00BD114F05B275A8A94954047983C5CD96ED95C782D2ED65A18E78C98E8EAFBAF58BBD046BE9895AD55FD0FF95907E7EBD6ACA2688D24779BDE9F0AAB13924CE65F597F9C9B9953DDBACF51DA7113FBAB9BE1DF6C6EA836DEB48983CCDCFC4125B5013D0CE52F890D0C391A035D30BCD5169A3451FD7023685274576DCB5F8FA47#)
  (q #00D1DA3915346A704EB2D2F2A48CD48F3DCC4CF25D#)
  (g #501BCFB989AD2C346BBD7782CA0230551F976B1A07EE3AEE27E4B63B7B00B1ACA712AD85784986411278163156D4DBA9DF75C8560F9C2E02C02AEC830EC403A56B6F64432869D6CA9314A648076511343507629BF4FC96F8FDBB9797258DDF11F437B1450BA23F1AA7E885EC6A33D37B7D7EC384A004420DB238E140B94AAAFE#)
  (y #7C9CB7732164787DD1931BB58257665EB60D6AA72B8D64D634530A61BE93D5AF01427962646542F18401B73032B12B9CBCAE8E3CF080DAD55C6612A97D6D8776CF2CBDD3AAC75D302B60E6956E5B3C60B39E171A2D5F150A924C6E22981EFDF052D5C6507B2DEC15E96CB6CAF7B260D5386BBDD7D7F69B4BF14451D64D847AEB#)
  (x #00AB1E941176D94505911118AC799A504ADCCE88F8#)
  )
 )
 )
 (account
(name otrtest3)
(protocol prpl-aim)
(private-key 
 (dsa 
  (p #00BB4C57669E50E4C35F8E4CA84855CF2C83EE75C4F44B4BB4A7E88590D394D7A738E82EE97892E5051CE45E200741E18D423137AA8E6679B1CFAB4FF11D45D8C9CBDE388D30FC800B4879713E3C57BA48A92FE135BB9AF265F770B706FB9A04802244D12CBFFD97ACE5C73FCE88C2B716B4B22B994CD6429A7E16D9B6D1874137#)
  (q #00C40DA63B679A80FC31BF49A68503BB39754D0A45#)
  (g #6C0A48BEA859587D6677306D1777A2A0635470F149A86EB64EA62EAAA4C21ECE4375ACD016B776E3AD3411C18BB3FF37F963FCEBB8820FF8838AFA6FCD1B39558DAB78450AE2ED9457DEDBDCE13DF5A6B20A738D2973D375D360C044AF7F0204CCC372098F0B6460963274B1EA0B5FEC93571A15F5C03DCDF54EE83BB198F363#)
  (y #00AB2C8A82F020DB99EF5B7A8330EC43E0D5EBD623FEB67D1B046D88FACA01D8E31E4D7865DC62D4DA58CF8BC7FF4B57C203A9F7F5C85DAB1B63D63299EF13AD89AAA7E6638C9DBC42D096408936C9F0382224CFB5C1528DCC8C7F2554CB4CA2FF3C3239BC921F1C690295DD9AE69C8EF5BBD8E58A8FAA8BB9D5F88463CAECEE7B#)
  (x #7824B713A4E5FA6D6C69172196648CD4657A1ED1#)
  )
 )
 )
)
`

var expectedFingerprints = [][]byte{
	[]byte{
		0xff, 0x6d, 0xda, 0xdb,
		0xab, 0x39, 0xce, 0xc0,
		0x9f, 0x5b, 0x7c, 0x54,
		0x11, 0x9f, 0xf6, 0xe3,
		0xcb, 0x9e, 0x13, 0x59,
	},
	[]byte{
		0x55, 0xb7, 0xb5, 0x05,
		0xe0, 0x8e, 0x4f, 0x5d,
		0x24, 0x73, 0xdf, 0x28,
		0x29, 0x7e, 0xa6, 0xa5,
		0x23, 0x05, 0xe9, 0x3a,
	},
}

func TestImportFromLibOTRFileFormat(t *testing.T) {
	in := []byte(otrPrivateKeyFile)
	keys := ImportFromLibOTR(in)

	if len(keys) != 2 {
		t.Errorf("failed to import all private keys in file")
	}

	for i, k := range keys {
		exp := expectedFingerprints[i]
		if fp := k.Fingerprint(); !reflect.DeepEqual(fp, exp) {
			t.Errorf("unexpected fingerprint: %#v", fp)
		}
	}
}
