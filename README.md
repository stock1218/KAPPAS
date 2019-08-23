# Key Allocation for Protected PANs on AWS Services (KAPPAS)

<div align="center">
  <img src=https://2.bp.blogspot.com/-zGHcRWCH0AA/VNMEcaBW8qI/AAAAAAABNiU/Gyje0I1awOk/s1600/satori%2Bkappa%2Blarge.jpg>
</div>

## Overview
KAPPAS is a PCI DSS compliant PAN storage framework for Go. It's mainly developed for AWS services, however it's designed to be extensible and could be used in a variety of other environments. It allows services to store and retrieve PAN information quickly and securely.

## What is PCI DSS
Payment Card Industry Data Security Standard (PCI DSS) is described as:
> The Payment Card Industry Data Security Standard (PCI DSS) is a set of standards created by major payment card companies 
to protect consumers and avoid liability by forcing businesses involved in the payment card ecosystem to implement safety
 measures and processes.

## How to use
At the moment, this is meant to be used as a framework for developing a PAN storage solution. Everything is documented using [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc), and specification documents can be found in the `docs` folder.


## Contributions
Contributions are welcome, please make sure you write tests and documentation when adding new code 🙂
#### TODO
- Finish gRPC view
- Package everything up in a docker container
- Package archetecture into AWS Cloud Formation

## Why?
I built this framework while interning at a startup that wanted to process PAN information for their existing platform. They were using an AWS infrastructure, and they needed to become PCI compliant. I then designed this solution that sucessfuly allowed them to quickly spin up a secure system for processing sensitive information. This framework was also designed to be highly extensible and maintainable, so that it could provide buisness value further into the future.

Throughout this project I learned a lot about compliance, securty engeneering, and technologies including Go, Docker, gRPC, and AWS.


## Useful Links
- Making AWS services PCI compliant: https://cloudacademy.com/blog/how-to-make-your-infrastructure-compliant-with-pci-dss-on-aws/
