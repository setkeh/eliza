Project Eliza
=============

[![Build Status](https://drone.io/github.com/thesetkehproject/eliza/status.png)](https://drone.io/github.com/thesetkehproject/eliza/latest)

Eliza is the Project Taking Kira's Place as the management controller for my internal network and Development enviroment.

Eliza's Purpse is to manage machines, updates, containers and hypervisors aswell as monitor services.


TODO:
----

  1. ~~Add Darthlukan's Gobot for IRC~~ See Notes.
  2. Add Docker integration.
  4. Add Machine management.
  5. Add a Logging Method.
  6. Add an API for interacting with Services.
  7. Add a Pretty web app for management.
  8. There is Plenty More to add to this list.

Todo's are in no Particular order.


Code i have Borrowed:
--------------------


Thanks to the Developers of all Borrowed Code.

Notes:
------

Now that its working ill be stripping the GoBot code back out of Eliza after Darthlukan applied some logic to the problem of notification integration.

This will Result in the Following Changes:

  * [GoBot Issue #17:](https://github.com/darthlukan/gobot/issues/17)
  * [GoBot Issue #18:](https://github.com/darthlukan/gobot/issues/18)
  * [GoBot Issue #19:](https://github.com/darthlukan/gobot/issues/19)

This will abstract the management of notification methods away from Eliza itself thus making code more manageable and removing a number of failure points.
It will Also allow for adding in support for other protocols i.e Tox or Telegram even Jabber without bulking the codebase too significantly.

Im also going to make each component its own library for example (Eliza-IPMI, Eliza-irc, and so on) and the main package will be what ties it all together housing the REST API and probably the Web interface.

Eliza's Management tasks are going to be very broad here is the short list:

  * Monitoring Network and Node health.
  * Generating alerts based on conditions relating to health monitors.
  * Monitoring Node and VM installation progress.
  * Configuring the network to allow isolation during Node and VM installation from the primary network.
  * Configuring Nodes and VM's during installation based on a Template (Likely JSON) for Post Installation setup of services and OS upgrades.
  * Keeping track of Recourses such as HDD, Bandwidth, Ram usage, CPU utilization for the entire network including networking equipment.
  * Node and VM Power management such as power on/off, hard reboot as well as soft reboot.

There are some things that wont be able to be completed because of limitations in my current network here is a list of things i will need to fully complete the project.

  * [Ubiquity Managed Gigabit Switch](http://www.scorptec.com.au/product/Networking_-_Wired/Gigabit_Switches/55583-ES-24-250W)

This list is likely to grow as the Project gets more and more complex.

I have a thing for Ubiquity hardware at the moment but i will likely try to integrate Cisco IOS and Arista hardware if i can get my hands on some hardware.


Contribute:
-----------

Please feel free to contribute by submitting pull requests.

You can also Donate to help contribute to the development and hardware upgrades/support you can hit me up via email or setkeh on freenode.
