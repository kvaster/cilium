To install Cilium on `Azure Kubernetes Service (AKS) <https://docs.microsoft.com/en-us/azure/aks/>`_,
perform the following steps:

**Default Configuration:**

=============== =================== ==============
Datapath        IPAM                Datastore
=============== =================== ==============
Direct Routing  Azure IPAM          Kubernetes CRD
=============== =================== ==============

.. tip::

   If you want to chain Cilium on top of the Azure CNI, refer to the guide
   :ref:`chaining_azure`.

**Requirements:**

* The AKS cluster must be created with ``--network-plugin azure`` for
  compatibility with Cilium. The Azure network plugin will be replaced with
  Cilium by the installer.

* User node pools must be tainted with ``node.cilium.io/agent-not-ready=true:NoSchedule``
  using the ``--node-taints`` option to ensure application pods will only be
  scheduled once Cilium is ready to manage them.

* To ensure applications pods are not deployed on system node pools and are
  properly managed by Cilium, it is recommended that:

  * The initial system node pool be replaced with another system node pool
    tainted with ``CriticalAddonsOnly=true:NoSchedule`` using the ``--node-taints``
    option. Replacement is necessary because it is not possible to assign taints
    to the initial AKS node pool at this time, cf. `Azure/AKS#1402 <https://github.com/Azure/AKS/issues/1402>`_

  * Any additional system node pool also be tainted with ``CriticalAddonsOnly=true:NoSchedule``
    using the ``--node-taints`` option.

**Limitations:**

* All VMs and VM scale sets used in a cluster must belong to the same resource
  group.
