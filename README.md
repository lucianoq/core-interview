# core-interview-test

This repository includes the YOTI Core team interview test, and its resources.

## The YOTI Core test

Implement an 'encryption-server' with two endpoints:

1) store endpoint - accepts some bytes and an id

  * Generates an AES encryption key
  * It encrypts the provided bytes using the generated key
  * Stores the encrypted bytes.
  * Returns the key used to encrypt the bytes

2) retrieve endpoint - accepts an id and an AES key

  * Retrieves the bytes stored under the provided data store key
  * Returns the decrypted bytes

A client interface has also been provided, please implement a client which 
satisfies the interface in order to interact with the above server.

### Extra

It would be desirable for the data store key to be difficult to derive from
the original id provided when storing the data  

### Notes

All work should be committed to a GIT repository where the commit history can be
reviewed. Your solution should also be easy to verify, to that end please feel
free to provide any further instructions, documentation, etc on how to go about
the verification process.
