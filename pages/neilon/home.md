# Neilon

Neilon is a Fabric text library to aid working with Minecraft text.\
View the roadmap on GitHub.

## Setup
Neilon requires any version of Vera over 1.0, as it uses `VColor` in the backend

```groovy
repositories {
    // ...
    maven { url = "https://artifacts.snackbag.net/artifactory/libs-release/" }
}

dependencies {
    // ...
    modImplementation "net.snackbag.mcvera:mcvera:<YOUR_VERA_VERSION>"
    modImplementation "net.snackbag.neilon-lib:neilon-lib:1.0.0" // latest version
}
```

You're good to go!
