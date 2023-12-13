# PoC CosmosJS - Textual Mode

## Introduction
The main goal of this project was to improve cosmosjs, a JavaScript package used for interacting with Cosmos-based blockchains, by adding support for a new data representation mode known as 'textual.' 
This enhancement would allow cosmosjs to convert between different representation modes and enable the signing of transactions in wallets using this textual format. 
The focus was on practical functionality improvements to facilitate smoother interactions within the Cosmos blockchain network.

## Discussion
The project began with an extensive review of the Cosmos SDK's approach to textual encoding to understand the underlying mechanisms. 
The initial attempt to replicate this functionality in JavaScript involved using proto reflection and proto file descriptors. 
However, this approach was limited by the existing JavaScript protobuf support, where the protojs package could not fully match Golang's capabilities.

## Proof of Concept (POC)
The project then explored compiling the Cosmos SDK's textual encoding methods into WebAssembly (WASM) to integrate with a JavaScript library. 
While this was not feasible with tinygo, it was successful with full Golang, albeit resulting in a large binary size of 40 MB. 
This POC demonstrated the technical possibility of the intended functionality, with some limitations in terms of practicality and efficiency.

## Conclusion
Despite the technical success of the POC, the practical application of such a JavaScript library in encoding and signing transactions for Cosmos zones presented significant challenges. 
The need for each zone to develop custom handlers for their messages complicates the creation of a universally applicable library. 
Additionally, the task of reviewing and auditing each implementation for errors or potential security issues adds another layer of complexity to the process.

## Further Recommendations
To address these challenges, we propose developing a generic, message-agnostic application for transaction signing. 
This application would receive the transaction payload and the necessary data structure for interpretation, signing a combination of metadata identification and the transaction payload. 
The node, pre-configured with the metadata identification, would validate the signature. This approach simplifies the transaction signing process and delegates 
the responsibility of message-specific handling to the node, streamlining the workflow and enhancing the security and integrity of the transactions.