# zkpop-go

`zkpop-go` is a Go wrapper around the 
[KEM-NIZKPoP](https://github.com/Chair-for-Security-Engineering/KEM-NIZKPoP)
project. It provides functions to generate keypairs, create Zero-Knowledge
Proofs of Possession (ZKPoP), and verify proofs.

## Getting Started

### Supported algorithms

Currently, our binding supports functions like Keygen, Keygen-zkpop, Encaps, Decaps, and Verify-zkpop for:
- Kyber512_avx2; and
- FrodoKEM-640.

### Prerequisites

Ensure you have the following installed on your system:
- GCC with AVX2, AES, and BMI2 support.
- OpenSSL development libraries:
  ```bash
  sudo apt install libssl-dev
  ```
- Go programming language (version 1.15+ recommended).

### Clone the Repository

First, clone the `zkpop-go` repository:

```bash
git clone --recurse-submodules https://github.com/gabrielzschmitz/zkpop-go.git
cd zkpop-go
```

If you have already cloned the repository without submodules, you can initialize and update them using:

```bash
git submodule update --init --recursive
```

### Build the External Library

1. Frodo-KEM
To build the library, navigate to `external/KEM-NIZKPoP/frodo-zkpop/`:

`make clean && make OPT_LEVEL=FAST USE_OPENSSL=FALSE GENERATION_A=SHAKE128 ZKPOP_N=65536 ZKPOP_TAU=8 && frodo640/test_KEM`

If you are going to use openssl, just do a `make` instead.

2. Kyber

It was tricky to make the binding work for kyber; You will need to:

- navigate to `/KEM-NIZKPoP/kyber-zkpop/avx2/` (In a hope that you have avx2 instructions!)
- change the Makefile using some editor:
```
#Comment the rule libpqcrystals_kyber512_avx2.so (below)
#libpqcrystals_kyber512_avx2.so: $(SOURCES) $(HEADERS) symmetric-shake.c
#        $(CC) -shared -fpic $(CFLAGS) -DKYBER_K=2 $(SOURCES) \
#          symmetric-shake.c -o libpqcrystals_kyber512_avx2.so
#and replace with
libpqcrystals_kyber512_avx2.so: $(SOURCES) $(HEADERS) symmetric-shake.c
        $(CC) -shared -fpic $(CFLAGS) -Wl,--allow-multiple-definition -DKYBER_K=2 $(SOURCES) \
          symmetric-shake.c zkpop.c zkpop.h  -o libpqcrystals_kyber512_avx2.so 

```
which actually do a export for zkpop symbols to kyber512_avx2.so library.
- compile with `make shared`
- It should generate a lot of `.so` files; you should copy them with `sudo cp *.so /usr/lib/`.
- If you want, check with `nm -D libpqcrystals_kyber512_avx2.so` whether `pqcrystals_kyber512_avx2_crypto_kem_keypair_nizkpop` is actually there. If not, the change in the Makefile was not effective (maybe `make clean && make shared`  can help).


### Build the Go Project

Before building the project, you need to change `main.go` because (sadly!) `CFLAGS` and `LDFLAGS` contains absolute paths. Change it to your path.

Now you can build the project (navigate to `zkpop-go/` directory):

```bash
go build -o zkpop`
```

### Execution

```bash
./zkpop
```

it should execute 10 times for each Frodo640 and Kyber512 operations.


## License

This project is licensed under the MIT License.
