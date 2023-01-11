[![Build Status](https://img.shields.io/travis/jeremybower/go-recaptcha/master.svg?style=flat-square)](https://travis-ci.org/jeremybower/go-recaptcha)
[![Coverage Status](https://img.shields.io/codecov/c/github/jeremybower/go-recaptcha/master.svg?style=flat-square)](https://codecov.io/gh/jeremybower/go-recaptcha)

# go-recaptcha

### About

This package handles [reCaptcha](https://www.google.com/recaptcha) ([API version 2.0](https://developers.google.com/recaptcha/intro)) form submissions.

### Installing

Install the package in your environment:

```
go get github.com/jeremybower/go-recaptcha
```

### Usage

To use it within your own code, import `github.com/jeremybower/go-recaptcha` and create a client:

```
client := recaptcha.NewClient(recaptchaPrivateKey)
```

Then for each form POST:

```
client.Confirm (clientIpAddress, recaptchaResponse)
```

use the form's `g-recaptcha-response` as the `recaptchaResponse` parameter.

The client's Confirm() function returns either true (i.e., the captcha was completed correctly) or false, and any error.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details