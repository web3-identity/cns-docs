# CNS-Docs

The Conflux Name Service (CNS) is a distributed, open, and extensible naming system based on the Conflux blockchain.

CNS’s job is to map human-readable names like ‘alice.web3’ to machine-readable identifiers such as Conflux addresses, other cryptocurrency addresses, content hashes, and metadata. CNS also supports ‘reverse resolution’, making it possible to associate metadata such as canonical names or interface descriptions with Conflux addresses.

CNS has similar goals to DNS, the Internet’s Domain Name Service, but has significantly different architecture due to the capabilities and constraints provided by the Conflux blockchain. Like DNS, CNS operates on a system of dot-separated hierarchical names called domains, with the owner of a domain having full control over subdomains.

Top-level domains, like ‘.web3’ and ‘.test’, are owned by smart contracts called registrars, which specify rules governing the allocation of their subdomains. Anyone may, by following the rules imposed by these registrar contracts, obtain ownership of a domain for their own use.

Because of the hierarchal nature of CNS, anyone who owns a domain at any level may configure subdomains - for themselves or others - as desired. For instance, if Alice owns 'alice.web3', she can create 'pay.alice.web3' and configure it as she wishes.

CNS is deployed on the Conflux main network and test networks.

You can try CNS out for yourself now by using the CNS Manager App, or by using any of the many CNS enabled applications on our homepage.

CNS contracts is forked from ENS, and is fully compatible with ENS.

Currently CNS's top domain is `.web3`

## CNS Architecture

CNS architecture is same as [ENS's architecture](https://docs.ens.domains/#ens-architecture).
