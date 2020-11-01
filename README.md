# GTA Chaos Mod Discord Bot

This is a Discord bot for [this mod](https://github.com/gta-chaos-mod/Trilogy-ASI-Script) which allows users in a Discord server to vote for the next effect, similar to how Twitch viewers can vote using the original GUI app.

## Usage

This isn't a public bot because that would make it a lot more complicated. To use it yourself, register a bot and get a bot token then set the two environment variables:

```ini
DISCORD_TOKEN=xyz
DISCORD_CHANNEL=123
```

The app uses dotenv so you can just put these into a file named `.env` in the repository root directory.

Once that is done, just run the app:

```sh
go run .
```

Now, start the game and you should see:

```
channel: (the channel ID you entered in .env)
Connected!
```

It will just print out the effects to the terminal, so you'll see stuff like this while playing:

```sh
writing effect:oh_hey_tanks:60000:Oh Hey, Tanks!:N/A:0
WINNER! 2 Health, Armor, $250k
writing votes:BringMeToTheDocks;0;;gang_members_everywhere;0;;health_armor_money;1;;2
writing effect:health_armor_money:60000:Health, Armor, $250k:N/A:0
```
