webpay library for golang
================================================

This is a **unofficial** library for WebPay (https://webpay.jp/)

godoc : http://godoc.org/github.com/tsukinowasha/go-webpay


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

Difference
-----------------

ByCustomer
++++++++++++++++++

Webpay can charge not only by the Card but also Customer Id. The
official library can use same function to both.

However golang does not permit it. In this library, you should use
another function.

::

  client := webpay.NewWebPayClient("YOUR_AUTH_TOKEN")
  ret, err := client.Charge.CreateByCustomer(400.0, "jpy", "cus_45d3MV5xxxxxxv")


All
++++++++++++++++++

Customer or other list can be acquired with count, offset or created arguments.
These args can be empty.

In this library, arguments are map[string]int.

::

  ret, err := client.Customer.All(map[string]int{

  })


Current Status
------------------------

- Charge

  - Create: done
  - CreateByCustomer: done
  - Retrieve: done
  - Refund: notyet
  - Capture: notyet
  - all: done

- Customer

  - Create: done
  - Retrieve: done
  - Update: notyet
  - delete: done
  - all: done
  - delete_active_card: notyet

- Token

  - Create: done
  - Retrieve: done

- Recursion

  - Create: done
  - Retrieve: done
  - Resume: notyet
  - delete: notyet
  - all: done

- Account

  - Retrieve: done
  - Delete: notyet

- Event

  - Retrieve: done
  - All: done

- Shop ?

LICENSE
-----------

MIT



