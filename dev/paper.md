// This is a draft! not the final version!
# The Polar Protocol
## Decentralized object storage network for resilient services
Draft of January 29 2022
by [monoko](https://github.com/monok-o)

### I. Introduction
For years, governments of the world always tried to keep control over the citizens of their countries by creating new “security laws”, “anti-terrorist act”, “data protection law”. The years has passed by and now we are starting to experience a high level of surveillance that is strongly impacting the democracy in the world. Many example can be taken in dictatorships like China which is famous for its terrible repression and their surveillance technologies but also in India where the intelligence agencies are monitoring activists to arrest them or even more recently in France where the police does not hesitate about using anti-terrorism law to get information about climate activists.

Even in countries where human rights are known for “”being respected””, political censorship and control is becoming something almost normal, something hard to fight since everything is stored by big companies and can be easily censored.

To avoid this we need to makes people at risk anonymous by collecting the least metadata, store the data on swarm of nodes ran by many independent parties worldwide so then a shutdown of one server will not be a problem for the data on the network and all of this while staying accessible to everyone on the clearnet but also on others onion-routing networks likes Lokinet, Tor, I2P, etc.

### II. Why?
it’s been now a bit more than 3 years since an idea like Coldwire came in my mind. For me resiliency of the network is the most important because it would not have any sens to provide services for activists and then being shutdown by an authoritarian government. Also the network to be simple to setup and cheap in term of resources for running it and accessible, so technologies like the blockchains were immediately disqualified because of their resources need or the cost to just put data on it. So for my first prototype I started to build a kind of collaborative hosting service by connecting machines of users with a mesh networks using wireguard, but it was way too unstable and also hard to setup.

After some months thinking about this idea, I started to get more interest for federated services, being able to login to my account from any other instances using a software almost everywhere, that’s perfect! But the issue still the same, if my instance is down, my account and my data will be lost... But something I’m sure now, the provider of the services have to be everyone, every single organization or individual should be able to setup their own instance of their services, which could definitely help for resiliency, but the issue of the loss of data is still here... It’s here that the idea of polar came in my mind: When someone install an instance of a service of coldwire, they can be part of the polar network and take advantage of it (resiliency of data, decentralized entities system, etc) while participating to it by providing some disk space.

![figure-1](./fig-1.png)

As you can seen the polar nodes also act as “entry nodes” for the services hosted on the same infrastructure, in fact the polar nodes will expose a REST API over a unix socket that any software on the same machine/infrastructure will be able to use for interacting with the network.

### III. How it's working?
1. Nodes discovery<br>
When starting a node for the first time, a “seed server” can be specified in the command line if it’s not specified polar will use the public instance of coldwire and if polar detect there is already nodes registered on the local database it will just pick a random node from the list and send a “seed request” to it, because yes, what I call “seed servers“ are just ordinary servers.<br><br>
For an example of the discovery protocol like in the figure below, we will take two server, S which will act as a seed and N as a node whose willing to meet new nodes. First N send a “seed” request to S, S answer with a list of all the node it collected during its lifetime, then N will notify every nodes it don’t know by sending a “register” request.<br><br>
![fig-2](./fig-2.png)

2. Local socket<br>
3. Basic functions<br>
  a. Push: Will data in chunks and will dispatch them over the network and then respond with the three hash of the files<br>
  b. Pull: <br>
  c. Erase: <br>
  d. Search: <br>
4. Storage<br>
  a. Entities<br>
  b. Objects<br>
  c. Metadata<br>
### IV. Theats and solutions
### V. Conclusion