# Chapter 2 :: Inner Loop: Development Environments

On this chapter, we will be changing the Appointments Service for the MinSalus application. 
Developers are requested to extend the Appointment data structure to support appointment type. 

[Requirement](https://github.com/salaboy/devex-on-k8s/issues/1)

Readers can work inside the [go/begin](/chapter02/go/begin/) and [java/begin](/chapter02/java/begin/) directories, while the [go/end](/chapter02/go/end/) and [java/end](/chapter02/java/end/) directories can be used to validate the changes that are expected to be performed.

The idea is to go over the main activities performed by developers on the inner loop. 

## Pre-requisites

Install Podman, or the container runtime of your choice and VScode.

## Clone 

[Clone this repository](https://github.com/salaboy/devex-on-k8s/):

```bash
git clone https://github.com/devex-on-k8s/book
```

Then move to [chapter02/go/begin/appointments](/chapter02/go/begin/appointments/) where you can find the source code of the appointments service. 

```bash
cd chapter02/go/begin/appointments
```

## Run / Change

Open the `chapter02/go/begin/appointments` directory in VSCode.
