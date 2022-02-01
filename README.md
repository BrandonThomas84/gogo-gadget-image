# gogo-gadget-image

Image Processor that will do some pretty nifty stuff


## How to run it (before and after building)

```bash
go run cmd/gogo-gadget-image/main.go -d=/path/to/image-directory -o=/path/to/ouput-directory

# OR AFTER BUILD

./gogo-gadget-image -help -d=/path/to/image-directory -o=/path/to/ouput-directory
```

## Command Line Options Help (before and after building)

```bash
go run cmd/gogo-gadget-image/main.go -help 

# OR AFTER BUILD

./gogo-gadget-image -help
```

## How to Build

```bash
go build ./cmd/gogo-gadget-image
```

