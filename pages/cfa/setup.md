# Setup

Let's guide you through the basic setup of your Create Mod addon. There's multiple base templates you can choose from.
We'll look at two of them.

### The Official Template

This template is maintained by the Create Fabric team themselves and contains various features such as premade pieces of
code inside the main class, as well as three optional recipe viewers and parchment mappings. You can find
it [here](https://github.com/Fabricators-of-Create/create-fabric-addon-template). Make sure to update all dependencies
first. You may come across repository issues when building for the first time.

### Our Template

This template is what we'll use throughout this tutorial and wiki. It uses Yarn mappings, has three uses Yarn mappings,
comes with EMI & JEI and no default code. We make sure to regularly update all dependencies ourselves and continue
potential legacy support. You can find it [here](https://github.com/snackbag/create-fabric-addon-template).

## Let's start

Use whatever template you like the most. In this tutorial, we're using our own template. All code we'll write is
available inside
the [snackbag/create-fabric-addon-wiki-example-mod](https://github.com/JXSnack/create-fabric-addon-wiki-example-mod)
repository.

1. Review your `fabric.mod.json` and adjust all values necessary (don't forget the LICENSE file itself!)
2. Adjust your package name (`/com/example/modid/...`) and `maven_group` in `gradle.properties`
3. Update your `archives_base_name` in `gradle.properties`
4. Rename your `modid.mixins.json`
5. Rename your `ExampleMod` and `ExampleModDataGenerator` class names
6. Finally, update the `MOD_ID` in `ExampleMod`

Restart IntelliJ and hit the `Minecraft Client` run configuration.