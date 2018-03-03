go-recaptcha
============

About
-----

This package handles [reCaptcha](https://www.google.com/recaptcha) ([API version 2.0](https://developers.google.com/recaptcha/intro)) form submissions in [Go](http://golang.org/).

Usage
-----

Install the package in your environment:

```
go get github.com/jeremybower/go-recaptcha
```

To use it within your own code, import "<tt>github.com/jeremybower/go-recaptcha</tt>" and create a client:

```
client := recaptcha.NewClient(recaptchaPrivateKey)
```

Then for each form POST:

```
client.Confirm (clientIpAddress, recaptchaResponse)
```

use the form's "<tt>g-recaptcha-response</tt>"as the "<tt>recaptchaResponse</tt>" parameter.

The client's Confirm() function returns either true (i.e., the captcha was completed correctly) or false, and any error.


See the [instructions](example/README.md) for running the example for more details.
