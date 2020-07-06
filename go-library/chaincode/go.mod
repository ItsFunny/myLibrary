module myLibrary/go-library/chaincode

        go 1.14

        require  (
        github.com/hyperledger/fabric-chaincode-go v0.0.0-20200511190512-bcfeb58dd83a
        myLibrary/go-library/go v0.0.0-00010101000000-000000000000
        )

        replace myLibrary/go-library/go => /Users/joker/go/src/myLibrary/go-library/go

        replace code.google.com/log4go => /Users/joker/Desktop/go-dependency/code.google.com/log4go