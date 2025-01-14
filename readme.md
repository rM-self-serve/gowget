# GoWget

This program exists since the Remarkable Tablet's default wget binary does not
impliment TLS certificate validation. There is a statically compiled wget binary
hosted by [Toltec](https://toltec-dev.org/) at https://github.com/toltec-dev/bootstrap, but it is only 
compiled for arm32 bit. Since statically compiling wget is non-trivial,
this binary is used for arm64 bit.

## Functionality

At the current time, the only implimented functionality is -O,
a flag to specify where the downloaded content is written to.
It is equivelent to wget's -O flag.
