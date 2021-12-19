# minus

The `minus` tool accepts data from stdin and filters _out_ any arguments passed.

It simplifies something like this:

```bash
some_command | grep -vF thing1 | grep -vF thing2 | grep -vF thing3
```

to look like this:

```bash
some_command | minus thing1 thing2 thing3
```

This can be done using a file with grep and the `-f` option, but sometimes that's a lot more work when experimenting manually.
