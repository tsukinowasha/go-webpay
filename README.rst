webpay library for golang
================================================

This is a library for a WebPay(https://webpay.jp/)

Note:
  This is an unofficial library. 

godoc 


Install
--------

::

   go get github.com/tsukinowasha/go-webpay


Usage
--------

::

  import "github.com/tsukinowasha/go-webpay"

  client := webpay.NewWebPayClient("YOUR_AUTH_TOKEN")
  ret, err := client.Charge.Create(
    400.0,
    "jpy",
    Card{
        number:    "4242-4242-4242-4242",
        exp_month: 11,
        exp_year:  2014,
        cvc:       "123",
        name:      "KEI KUBO",
    },
  )

  chargeId, _ := GetId(ret)


Current Status
------------------------

- Charge

  - Create: done
  - Retrieve: done

- Customer

  - Create: done
  - Retrieve: done
  - Update: notyet
  - delete: notyet
  - all: notyet
  - delete_active_card: notyet

- Recursion

  - Create: done
  - Retrieve: done

- Token

  - Create: done
  - Retrieve: done

- Account

  not yet

- Event

  not yet

- Shop

  not yet

LICENSE
-----------

MIT



