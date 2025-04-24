# Technical Details

This page is explicitly made for other modders and contributors. If you're just a normal user, this is not important
information for you.

## What the port is this?

There are two separate versions of TT20, both maintained by two separate
people. [TT20 (Fabric/Main/Original version)](https://github.com/snackbag/TT20), maintained by JX_Snack (aka JX or
Snack) and [TT20 (Forge version)](https://github.com/snackbag/TT20-Forge), maintained by Lemonnik6468 (aka Lemon). The
Forge version is meant to be as close to a letter-to-letter port of the Fabric version as possible, to make parity as
good as possible.

> Why don't you just use Artifactory?

Because JX_Snack doesn't want to maintain the mess Artifactory creates when combined with the multi-version
system [Stonecutter](https://stonecutter.kikugie.dev/)

## Base configuration

Both the Fabric and Forge versions use the Stonecutter versioning system made by Kikugie. For the automatic update
notifier we use ShitLib (Super Helpful and Intuitive Toolkit). It's a little library without any extra dependencies
which allows for things like easier configuration files, web requests and Discord webhooks.

## Math & Classes

Our main formula, the one the mod is named after, is `ticks*tps/20`. `ticks` being the amount of ticks a task needs,
`tps` the server's current tps and `20` the maximum possible tps the server can have. This way, if the tps would be 12
and the eating task would take 40 ticks, it would now only take 24 ticks. This shortens the eating time, but takes
physical time into account, allowing the user to have the eating speeds they're used to. You can find this formula in
the `util.TPSUtil` class file. See the `tt20` methods.

The `util.TPSCalculator` calculates the current TPS. It gives out four values:

* applicableMissedTicks -> All ticks that should have happened during the time of the last tick + the ones that weren't
  applicable (less than 1.0) of ticks before
* getTPS -> The exact current TPS
* getAverageTPS -> Average TPS of the last 40 ticks
* getMostAccurateTPS() -> If `getTPS` is more than `getAverageTPS` return `getAverageTPS`, otherwise return `getTPS`. We
  do this because `getTPS` sometimes has huge performance/lag spikes

