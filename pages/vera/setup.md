# Setup

Check for the [latest version](https://github.com/snackbag/vera/blob/main/gradle.properties). Vera currently only
supports Fabric 1.20.1, all other versions may behave unexpectedly.

```gradle
repositories {
    maven { url = "https://artifacts.snackbag.net/artifactory/libs-release/" }
}

dependencies {
    modImplementation "net.snackbag.mcvera:mcvera:${project.vera_version}"
}
```