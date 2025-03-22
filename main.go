// main.go
package main

/*
#cgo CFLAGS: -I/home/<SEULOCAL>/zkpop-go/external/KEM-NIZKPoP/frodo-zkpop/frodo640/ -I/home/<SEULOCAL>/zkpop-go/external/KEM-NIZKPoP/kyber-zkpop/avx2/ -I/usr/include/ 
#cgo LDFLAGS: -L/home/<SEULOCAL>/zkpop-go/external/KEM-NIZKPoP/frodo-zkpop/frodo640/ -L/home/<SEULOCAL>/zkpop-go/external/KEM-NIZKPoP/kyber-zkpop/avx2/ -L/usr/lib/ -lfrodo -lpqcrystals_kyber512_avx2 -lpqcrystals_aes256ctr_avx2 -lpqcrystals_fips202_ref -lpqcrystals_fips202x4_avx2 -lssl -lcrypto 
*/
import "C"

import (
	"fmt"
	"log"
	"bytes"
	"zkpop-go/zkpop" 
)

//test FrodoKEM in N+1 iterations
func testFrodoKEM(N int){
        fmt.Println("Testing FrodoKEM...")

        //warmup
        pk, sk, err := zkpop.KeyPairFrodo640()
        if err != nil {
                log.Fatalf("Error generating keypair: %v", err)
        }

        //test N keygens
        for i := 0; i < N; i++  {
                pk, sk, err = zkpop.KeyPairFrodo640()
        }
	
	//warmup Encaps
	ct, ss, err := zkpop.EncapsFrodo640(pk)
	if err != nil {
                log.Fatalf("Failed FrodoKEM encapsulation: %v", err)
        }

	//test N Encaps
	for i := 0; i < N; i++  {
                ct, ss, err = zkpop.EncapsFrodo640(pk)
        }

	//warmup Decaps
	css, err := zkpop.DecapsFrodo640(ct, sk)
	if err != nil || !bytes.Equal(ss,css) {
                log.Fatalf("Failed FrodoKEM decapsulation.")
        }

	//test N Decaps
	for i := 0; i < N; i++  {
                css, err = zkpop.DecapsFrodo640(ct, sk)
        }

}


//test FrodoKEM-NIZKPoP in N+1 iterations
func testFrodoKEMNIZKPoP(N int){
  	fmt.Println("Testing FrodoKEM-NIZKPoP...")

	//warmup
        pk, _, zkpopProof, err := zkpop.KeyPairFrodo640NIZKPoP()
        if err != nil {
                log.Fatalf("Error generating keypair: %v", err)
        }

	//test N keygens
	for i := 0; i < N; i++  {
        	pk, _, zkpopProof, err = zkpop.KeyPairFrodo640NIZKPoP()
	}

        //fmt.Printf("Public Key: %x\n", pk)
        //fmt.Printf("Secret Key: %x\n", sk)
        //fmt.Printf("Zero-Knowledge Proof: %x\n", zkpopProof)

	//warmup
	valid := zkpop.VerifyFrodo640ZKPop(pk, zkpopProof)
        if !valid {
                log.Fatalf("Error verifying ZKPoP: %v", err)
	}

	//test N verifications
	for i := 0; i < N; i++  {
		valid = zkpop.VerifyFrodo640ZKPop(pk, zkpopProof) 
	}
}

//test Kyber512-NIZKPoP in N+1 iterations
/*func testKyberNIZKPoP(N int){
	//warmup
        _, _, _, err := zkpop.KeyPairKyber512NIZKPoP()
        if err != nil {
                log.Fatalf("Error generating keypair: %v", err)
        }

}*/

func testKyber(N int){
	fmt.Println("Testing Kyber...")
	 //warmup
        pk, sk, err := zkpop.KeyPairKyber512()
        if err != nil {
                log.Fatalf("Error generating keypair: %v", err)
        }

	//test N keygens
        for i := 0; i < N; i++  {
                pk, sk, err = zkpop.KeyPairKyber512()
        }

        //warmup Encaps
        ct, ss, err := zkpop.EncapsKyber512(pk)
        if err != nil {
                log.Fatalf("Failed Kyber512 encapsulation: %v", err)
        }

        //test N Encaps
        for i := 0; i < N; i++  {
                ct, ss, err = zkpop.EncapsKyber512(pk)
        }

	//warmup Decaps
        css, err := zkpop.DecapsKyber512(ct, sk)
        if err != nil || !bytes.Equal(ss,css) {
                log.Fatalf("Failed Kyber512 decapsulation.")
        }

        //test N Decaps
        for i := 0; i < N; i++  {
                css, err = zkpop.DecapsKyber512(ct, sk)
        }
}


func main() {
	N := 10
	fmt.Printf("Testing %d iterations for each algorithm...\n", N)

	//Frodo-KEM
	testFrodoKEM(N)
	testFrodoKEMNIZKPoP(N)

	//Kyber
	testKyber(N)
	//testKyberNIZKPoP(N)

	fmt.Println("End of testing.")
}

