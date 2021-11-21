# Bloom Filter

[Bloom Filter](https://en.wikipedia.org/wiki/Bloom_filter) is a space-efficient probabilistic data structure, conceived by Burton Howard Bloom in 1970.

This implement is experimental to written Golang to prove the basic concept of Bloom Filter.

## Install

```
go get github.com/zhong-my/bloom-filter
```

## Usage

```
    // Create a counting bloom filter expect size 100, false detect rate 0.01.
    cbf := NewCountingBloomFilter(100, 0.01)

    // Add item
    cbf.Add([]byte("abc"))
    cbf.Add([]byte("efg"))
    cbf.Add([]byte("hijk"))

    // Test Add -> return "true"
    fmt.Println("Test cfb:", cfb.Test([]byte("abc")))

    // Remove item
    cbf.Remove([]byte("efg"))

    // Test Remove -> return "false"
    fmt.Println("Remove cfb:", cbf.Test([]byte("efg")))

```

## Inspired

- [https://github.com/kkdai/bloomfilter](https://github.com/kkdai/bloomfilter)
- [https://github.com/bits-and-blooms/bloom](https://github.com/bits-and-blooms/bloom)

## License

This package is licensed under MIT license. See LICENSE for details.