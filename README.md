# IOSBlockchain - A Secure & Scalable Blockchain for Smart Services

The Internet of Services (IOS) offers a secure & scalable infrastructure for online service providers. Its high TPS, scalable and secure blockchain, and privacy protection scales social and economic cooperation to a new level. For more information about IOS Blockchain, please refer to our [whitepaper](https://github.com/iost-official/Documents)

The source codes released in this repo represent part of our early alpha-quality codes, with the purpose to demonstrate our development progress. Currently, there are still parts of our project have been intensively working on in private repos. We will gradually release our code.

## Overview

This is the first release version (v0.1.0) of IOST, including the prototype of protocol, blockchain, database and testing libs.

>“I could be bounded in a nutshell, and count myself a king of infinite space!”  - Hamlet II.ii

## Development Progress

v0.1.0 - MVP completed and preliminary tests conducted

v0.2.0 - Under development.

v0.5.0 (Everest) - Developed a complete blockchain framework that supports smart contract and validated PoB consensus mechanism

Our developers work in their own trees, then submit pull requests when they think their feature or bug fix is ready.

Although we have started to go open source on Github starting April 9th, 2018 and released certain parts of our code/repositories, please note that we have only released a portion of them at this moment.  Since the project is still at its very early stage and some codes are related to our core technology, we have decided to go open source gradually until the launch of the main net.

## Test net

Test net now is online, check further information below :

[Smart Contract Handbook](https://github.com/iost-official/Go-IOS-Protocol/wiki/Smart-Contract-Handbook)

[iwallet Handbook](https://github.com/iost-official/Go-IOS-Protocol/wiki/iwallet)

## Installation

*We recommended build and test with docker to avoid NAT issues as below*

[Docker Installation](https://github.com/iost-official/Go-IOS-Protocol/wiki/Docker-Installation)

### Building from source
#### Prerequisites
* Golang 1.10.1 (or newer) is required to build this project
* Redis 4.0.10 (or newer) is required, We recommend Redis stable version

#### Build
Get this repo:
```
go get -u -v github.com/iost-official/Go-IOS-Protocol
```
Change Directory into project:
```
cd $GOPATH/src/github.com/iost-official/Go-IOS-Protocol
```
Build iserver:
```
make build
```

### Run server
Start the redis
```
redis-server
```
Modify the `listener-addr` field of `iserver/iserver.yml` to the current server public IP address, then start the iserver
```
./build/iserver --config iserver/iserver.yml 
```

### Run wallet
We provides a wallet tool to send transaction and deploy contract. You can check the help by running the following command. The detailed document please check [Iwallet](https://github.com/iost-official/Go-IOS-Protocol/wiki/Iwallet)
```
./build/iwallet -h
```

## Contribution

Contribution of any forms is appreciated!

Currently, our core tech team is working intensively to develop the first stable version and core block chain structure. We will accept pull request after the first stable release published.

If you have any questions, please email to team@iost.io

## Community & Resources

Make sure to check out these resources as well for more information and to keep up to date with all the latest news about IOST project and team.

[/r/IOSToken on Reddit](https://www.reddit.com/r/IOStoken)

[Telegram](https://t.me/officialios)

[Twitter](https://twitter.com/IOStoken)

[Official website](https://iost.io)

## Disclaimer

- IOS Blockchain is unfinished and some parts are highly experimental. Use the code at your own risk.

- THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


