# PoC CosmosJS - Textual Mode
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

---

![zondax_light](docs/assets/zondax_light.png#gh-light-mode-only)
![zondax_dark](docs/assets/zondax_dark.png#gh-dark-mode-only)

_Please visit our website at [zondax.ch](https://www.zondax.ch)_

---

## About the repository :book:

### Introduction
The main goal of this project was to improve cosmosjs, a JavaScript package used for interacting with Cosmos-based blockchains, by adding support for a new data representation mode known as 'textual.' 
This enhancement would allow cosmosjs to convert between different representation modes and enable the signing of transactions in wallets using this textual format. 
The focus was on practical functionality improvements to facilitate smoother interactions within the Cosmos blockchain network.

### Discussion
The project began with an extensive review of the Cosmos SDK's approach to textual encoding to understand the underlying mechanisms. 
The initial attempt to replicate this functionality in JavaScript involved using proto reflection and proto file descriptors. 
However, this approach was limited by the existing JavaScript protobuf support, where the protojs package could not fully match Golang's capabilities.

### Proof of Concept (POC)
The project then explored compiling the Cosmos SDK's textual encoding methods into WebAssembly (WASM) to integrate with a JavaScript library. 
While this was not feasible with tinygo, it was successful with full Golang, albeit resulting in a large binary size of 40 MB. 
This POC demonstrated the technical possibility of the intended functionality, with some limitations in terms of practicality and efficiency.

### Conclusion
Despite the technical success of the POC, the practical application of such a JavaScript library in encoding and signing transactions for Cosmos zones presented significant challenges. 
The need for each zone to develop custom handlers for their messages complicates the creation of a universally applicable library. 
Additionally, the task of reviewing and auditing each implementation for errors or potential security issues adds another layer of complexity to the process.

### Further Recommendations
To address these challenges, we propose developing a generic, message-agnostic application for transaction signing. 
This application would receive the transaction payload and the necessary data structure for interpretation, signing a combination of metadata identification and the transaction payload. 
The node, pre-configured with the metadata identification, would validate the signature. This approach simplifies the transaction signing process and delegates 
the responsibility of message-specific handling to the node, streamlining the workflow and enhancing the security and integrity of the transactions.

### ADR for Reference
- [ADR20](https://docs.cosmos.network/main/build/architecture/adr-020-protobuf-transaction-encoding)
- [ADR50](https://docs.cosmos.network/main/build/architecture/adr-050-sign-mode-textual)


## About the repository :gear:
The current Proof of Concept (PoC) is implemented as a Go (Golang) project. In this project, we have developed a dedicated method for enabling textual signing of transactions. 
This capability is achieved through the integration of the Cosmos Software Development Kit (SDK), which has been included in our project's framework.

This method is then exported as part of the WebAssembly (WASM) binary during the compilation process. 
Consequently, this WASM module can be imported into a JavaScript environment, allowing the execution of the previously described transaction signing method.

To compile the WebAssembly (WASM) binary from the Go project, please use the following command:
```
make build
```

After completion, the compiled binary can be found in the output directory.

