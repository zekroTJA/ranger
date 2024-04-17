# ranger

This Go package is a collection of useful iterator functions.

Currently, this is only for some shenanigans with range functions, but over time, I will add more utilities to make this a useful and well documented package.

> [!WARNING]  
> Iterator functions are an [experimental feature in Go 1.22.0](https://go.dev/wiki/RangefuncExperiment). To use them, you need to set `rangefunc` in the `GOEXPERIMENT` environment variable.  
> Also, that means that **the API of this package might change a lot** until I find it stable and useful enough to release a `v1.0`.