# Coin Market Cap

A Coin Market Cap API client.

Only functions I need have been added.


## Example Usage

A binary has been provided to show some example usage for the "Quotes Latest" call.

```
$ API_KEY=xxx go run cmd/coinmarketcap-client/main.go --symbols BSV,ETH --convert AUD

BSV_AUD = {LastUpdated:2021-08-08T23:54:15.000Z ...}
ETH_AUD = {LastUpdated:2021-08-08T23:54:15.000Z ...}
```

## References

- [Developer Login](https://pro.coinmarketcap.com/)
- [API Documentation](https://coinmarketcap.com/api/documentation/v1/)


## License

The MIT License (MIT)

Copyright (c) 2021 Scott Barr

See [LICENSE](LICENSE)
