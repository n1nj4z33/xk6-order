# Summary
xk6 extension for convert orderID from int64 to string

# Install

```
go install go.k6.io/xk6/cmd/xk6@latest
```

# Compile

```
xk6 build --output ./bin/k6 --with xk6-order=.
```

Will produce `k6` executable with `order` extension.

# Tests

```
./bin/k6 run tests/test.js
```