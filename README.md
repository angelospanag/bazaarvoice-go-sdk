# bazaarvoice-go-sdk

A wrapper around the BazaarVoice API written in Go (WIP)


## Install

```go get github.com/angelospanag/bazaarvoice-go-sdk```

## Add your API key

In order to use the BazaarVoice API you need to create an account and request API keys:

https://developer.bazaarvoice.com/

Then add your API keys to the `config.toml` file located at the root of the project.

Example:
```
[production]
conversations_api_key = "YOUR_API_KEY_HERE"
```
