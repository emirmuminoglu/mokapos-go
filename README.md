[![Go Report Card](https://goreportcard.com/badge/github.com/emirmuminoglu/mokapos-go)](https://goreportcard.com/report/github.com/emirmuminoglu/mokapos-go)

# mokapos-go


A client library for *[Moka Pos](https://www.moka.com/moka-pos/)*

**Warning:** It's not tested and still in development! Please check the project and test before use. I don't take any liability.

# Examples


 - Payment without 3D Secure
```go

import (
    "github.com/emirmuminoglu/mokapos-go"
)

var (
    //Your Dealer Username
    username = "test"

    //Your Dealer Password
    password = "test"

    //If you want to use test environment set this value true
    isTest = true

)

func main(){
    m := moka.New(username, password, isTest)

    request := moka.PaymentDealerRequest{
        CardHolderFullName: "Emir Muminoglu",
        CardNumber: "5555666677778888",
        ExpMont: "12",
        ExpYear:"2020",
        CvcNumber: "123",
        Amount: 123.123,
        Currency:"TL",
        InstallmentNumber:"1",
    }

    resp, err := m.PaymentNonThreeD(request)
    if err != nil {
        fmt.Fatalf("An error occured: %v \n", err.Error())
    }

    fmt.Println(resp)
    return
}
```

 - Payment with 3D Secure

 ```go
import (
    "github.com/emirmuminoglu/mokapos-go"
)

var (
    //Your Dealer Username
    username = "test"

    //Your Dealer Password
    password = "test"

    //If you want to use test environment set this value true
    isTest = true

)

func main(){
    m := moka.New(username, password, isTest)

    customerInfo := moka.CustomerInformation{
        DealercustomerID: 1234,
        CustomerCode: "1234",
        FirstName: "emir",
        LastName:"muminoglu",
    }

    request := moka.PaymentDealerRequest{
        CardHolderFullName: "Emir Muminoglu",
        CardNumber: "5555666677778888",
        ExpMont: "12",
        ExpYear:"2020",
        CvcNumber: "123",
        Amount: 123.123,
        Currency:"TL",
        InstallmentNumber:"1",
        RedirectURL: "https://pos.testmoka.com/DealerPayment/PayResult?MyTrxId=1A2B3CD456",
        RedirectType: 0,
        CustomerInformation: customerInfo,
    }

    resp, err := m.PaymentWithThreeD(request)
    if err != nil {
        fmt.Fatalf("An error occured: %v \n", err.Error())
    }

    fmt.Println(resp)
    return
}

 ```
