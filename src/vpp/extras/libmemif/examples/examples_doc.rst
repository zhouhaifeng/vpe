.. _libmemif_examples_doc:

Examples
========

After build, root folder will contain scripts linking binary examples
with library (same name as example apps). These scripts can be executed
to run example apps without installing the library. Example apps
binaries can be found in *libs* folder. To run binaries directly, make
sure that libmemif library is installed.

Run in container
----------------

``ligato/libmemif-sample-service`` image contains built and installed
libmemf. To run different examples, override docker CMD to start
container in bash:

::

   # docker run -it --entrypoint=/bin/bash -i --rm --name icmp-responder --hostname icmp-responder --privileged -v "/run/vpp/:/run/vpp/" ligato/libmemif-sample-service

Current WORKDIR is set to root repository directory. Example apps can be
run from this directory (a script linking binary with library), or
browse to ``./.libs`` folder and execute binary directly.

* ``extras/libmemif/examples/icmp_responder``

Simplest implementation. Event polling is handled by libmemif.
Single memif connection in slave mode is created (id 0). Use Ctrl + C to exit app.
Memif receive mode: interrupt.

* ``extras/libmemif/examples/icmp_responder-epoll``  (run in container by default)

Supports multiple connections and master mode. User can create/delete connections, set ip addresses,
print connection information. :ref:`libmemif_example_setup_doc` contains instructions on basic
connection use cases setups. Memif receive mode: interrupt. App provides functionality to disable
interrupts for specified queue/s for testing purposes. Polling mode is not implemented in this example.

* ``extras/libmemif/examples/icmp_responder-mt``

Multi-thread example, very similar to icmpr-epoll. Packets are handled in threads assigned to specific queues. Slave mode only. Memif receive mode: polling (memif_rx_poll function), interrupt (memif_rx_interrupt function). Receive modes differ per queue.
