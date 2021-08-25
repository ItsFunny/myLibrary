module myLibrary/go-library/go

go 1.14

replace (
	github.com/cloudflare/cfssl => /Users/joker/go/src/github.com/cloudflare/cfssl
	github.com/hyperledger/fabric => /Users/joker/go/src/github.com/hyperledger/fabric
	github.com/hyperledger/fabric-droplib => /Users/joker/Desktop/gopath/src/github.com/hyperledger/fabric-droplib
	github.com/hyperledger/fabric-protos-go => /Users/joker/go/src/github.com/hyperledger/fabric-protos-go/
	github.com/tjfoc/gmsm => /Users/joker/go/src/github.com/tjfoc/gmsm
	myLibrary/go-library/go/log => /Users/joker/go/src/myLibrary/go-library/go/log
)

require (
	code.google.com/log4go v0.0.0-00010101000000-000000000000
	github.com/FactomProject/basen v0.0.0 // indirect
	github.com/FactomProject/btcutilecc v0.0.0-20130527213604-d3a63a5752ec // indirect
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/SebastiaanKlippert/go-wkhtmltopdf v1.5.0
	github.com/astaxie/beego v1.12.2
	github.com/emirpasic/gods v1.12.0
	github.com/ethereum/go-ethereum v1.9.15
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/hbakhtiyor/schnorr v0.1.0
	github.com/hyperledger/fabric v0.0.1
	github.com/jinzhu/now v1.1.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/common v0.10.0
	github.com/satori/go.uuid v1.2.0
	github.com/signintech/gopdf v0.9.8
	github.com/sirupsen/logrus v1.4.2
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/smartystreets/goconvey v1.6.4
	github.com/stretchr/testify v1.6.1
	github.com/tealeg/xlsx v1.0.5
	github.com/tjfoc/gmsm v1.4.0
	github.com/tyler-smith/go-bip32 v0.0.0-20170922074101-2c9cfd177564
	github.com/tyler-smith/go-bip39 v1.0.2
	github.com/valyala/fasthttp v1.14.0
	github.com/wumansgy/goEncrypt v0.0.0-20190822060801-cf9a6f8787e4
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee
	golang.org/x/text v0.3.3
	myLibrary/go-library/go/log v0.0.0-00010101000000-000000000000
)

replace github.com/FactomProject/basen => /Users/joker/Desktop/go-dependency/github.com/FactomProject/basen

replace code.google.com/log4go => /Users/joker/Desktop/go-dependency/code.google.com/log4go
