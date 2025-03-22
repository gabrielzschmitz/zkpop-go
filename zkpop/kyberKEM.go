// Kyber original bindings, just for comparison purposes
package zkpop

/*
#include "kyber/api_kyber.h"
#include <stdint.h>
#include <stdlib.h>
#include <openssl/evp.h> //from OpenSSL 1.1.1
#include <openssl/aes.h>
*/
import "C"

import (
        "fmt"
        "unsafe"
)

//binding for pqcrystals_kyber512_avx2_keypair (might need to change to 'ref' instead of avx2.)
func KeyPairKyber512()([]byte, []byte, error){
        pk := make([]byte, C.pqcrystals_kyber512_PUBLICKEYBYTES)
        sk := make([]byte, C.pqcrystals_kyber512_SECRETKEYBYTES)

        ret := C.pqcrystals_kyber512_avx2_keypair((*C.uint8_t)(unsafe.Pointer(&pk[0])),
                (*C.uint8_t)(unsafe.Pointer(&sk[0])))

        if ret != 0 {
                return nil, nil, fmt.Errorf("failed to generate keypair")
        }
        return pk, sk, nil
}


//Encapsulation for a given public key pk
//Returns a ciphertext ct and a 16-byte shared secret ss
func EncapsKyber512(pk []byte)([]byte, []byte, error){
//      crypto_kem_enc_Frodo640
//      (unsigned char *ct, unsigned char *ss, const unsigned char *pk)
        ss := make([]byte, C.pqcrystals_kyber512_BYTES)
        ct := make([]byte, C.pqcrystals_kyber512_CIPHERTEXTBYTES)

        ret := C.pqcrystals_kyber512_avx2_enc((*C.uint8_t)(unsafe.Pointer(&ct[0])),
                                           (*C.uint8_t)(unsafe.Pointer(&ss[0])),
                                           (*C.uint8_t)(unsafe.Pointer(&pk[0])))
        if ret != 0 {
                return nil, nil, fmt.Errorf("failed to encaps")
        }
        return ct, ss, nil
}


//Given a ciphertext ct and a private key sk
//returns a candidate shared secret css
func DecapsKyber512(ct []byte, sk []byte)([]byte, error){
        css := make([]byte, C.pqcrystals_kyber512_BYTES)

        ret := C.pqcrystals_kyber512_avx2_dec((*C.uint8_t)(unsafe.Pointer(&css[0])),
                                           (*C.uint8_t)(unsafe.Pointer(&ct[0])),
                                           (*C.uint8_t)(unsafe.Pointer(&sk[0])))
        if ret != 0 {
                return nil, fmt.Errorf("failed to perform decapsulation")
        }
        return css, nil
}






