# Setup

This part of the wiki will guide you through the simple process of installing TT20 on your server.

## Basics

**Server Installation**

1. Head over to the [Modrinth page](https://modrinth.com/mod/tt20)
2. Download the version that matches your server's Minecraft version and mod loader
3. Drop the `.jar` file into your server's `mods` folder

That's it! You've got TT20 running on your server.

BEGIN NOTE
**Note:** Players do **do not** need to install TT20 on their clients. In fact, **they shouldn't**. Even vanilla clients can connect and benefit from TT20's optimizations.
END NOTE

---

**Singleplayer Installation**

1. Drop TT20 into your client's `mods` folder like any other mod
2. Start the game at least once to let TT20 generate its config files
3. Navigate to your game's config folder and select the `tt20` folder within
4. Open `config.json` and set `"singleplayer-warning": false,`\
   ⚠️ Don’t forget the comma at the end of the line!
5. Either restart your game or run the `/tt20 reload` command in-game 

## Modpacks

In a modpack, it's usually best to avoid maintaining separate versions for the server and the client. That's why,
starting with version 0.8, TT20 gained the ability to automatically disable itself in singleplayer. To enable this
behavior, set `singleplayer-enabled` to `false` in the `config.json` file. 
