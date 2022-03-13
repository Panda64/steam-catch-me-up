# Steam Catch Me Up

This simple program takes the current top 10 games on steam (by player count) and displays their relevant info and stats, including the most viewed youtube video about the game within the last week.

## Usage

In order to properly run this program, you must have a [YouTube Data API](https://developers.google.com/youtube/v3/getting-started) developer key. After putting your key in a `.env` file under the variable name `YOUTUBE_API_KEY`, you can then run the program by

```bash
$ go run app.go
```

Your ouput should look similar to the following:

```
Getting real-time steam stats...
--------------------------------------------------------------------------------
Let's catch you up on what people are playing today...

Rank:  1
Game: Lost Ark
Current Players: 818,822
Peak Today: 908,335
Store Link: https://store.steampowered.com/app/1599340/Lost_Ark/
Trending Video: 'Asmongold Quits Elden Ring.' ---> by Asmongold TV
https://www.youtube.com/watch?v=4HF9lYsLQN4


Rank:  2
Game: ELDEN RING
Current Players: 704,918
Peak Today: 944,190
Store Link: https://store.steampowered.com/app/1245620/ELDEN_RING/
Trending Video: 'I&#39;m An Elden Ring God' ---> by penguinz0
https://www.youtube.com/watch?v=BnH5tedjgsQ


Rank:  3
Game: Counter-Strike: Global Offensive
Current Players: 509,814
Peak Today: 963,573
Store Link: https://store.steampowered.com/app/730/CounterStrike_Global_Offensive/
Trending Video: 'LEGO Full-Auto P90 | Asiimov [Blowback Rubber Band Gun] - Counter-Strike: Global Offensive' ---> by Kevin183
https://www.youtube.com/watch?v=qTbU4t9C6QI


Rank:  4
Game: Dota 2
Current Players: 388,846
Peak Today: 663,120
Store Link: https://store.steampowered.com/app/570/Dota_2/
Trending Video: 'Dota 2 WTF Moments 7.31' ---> by Dota Watafak
https://www.youtube.com/watch?v=1EAj17sfL_o


Rank:  5
Game: Destiny 2
Current Players: 154,493
Peak Today: 222,695
Store Link: https://store.steampowered.com/app/1085660/Destiny_2/
Trending Video: 'Destiny 2: The Witch Queen - Vow of the Disciple - World First Raid Race' ---> by Destiny 2
https://www.youtube.com/watch?v=WiAd15wfVRo


Rank:  6
Game: Apex Legends
Current Players: 134,805
Peak Today: 313,383
Store Link: https://store.steampowered.com/app/1172470/Apex_Legends/
Trending Video: 'PLAYING MASTERS RANKED LIKE IT&#39;S PUBS w/aceu &amp; sYnceDez' ---> by iiTzTimmy
https://www.youtube.com/watch?v=4wZB4n2UCvY


Rank:  7
Game: Rust
Current Players: 105,634
Peak Today: 131,026
Store Link: https://store.steampowered.com/app/252490/Rust/
Trending Video: 'КРЕПОСТЬ! МЫ ПОСТРОИЛИ НЕПРОБИВАЕМЫЙ ДОМ в RUST/РАСТ' ---> by ДЕРЖИ ДВЕРЬ
https://www.youtube.com/watch?v=gWEoDdFcQ8w


Rank:  8
Game: PUBG: BATTLEGROUNDS
Current Players: 101,283
Peak Today: 503,148
Store Link: https://store.steampowered.com/app/578080/PUBG_BATTLEGROUNDS/
Trending Video: '🔴Live สด!  “LEO PUBG Thailand Series Season 7” Day 15' ---> by PUBG: BATTLEGROUNDS THAILAND
https://www.youtube.com/watch?v=9EvrqXRN8AM


Rank:  9
Game: Team Fortress 2
Current Players: 83,016
Peak Today: 86,893
Store Link: https://store.steampowered.com/app/440/Team_Fortress_2/
Trending Video: 'TF2 sniper this weapon needs a nerf' ---> by Zhain 2 Remastered
https://www.youtube.com/watch?v=ggzfD58PpEA


Rank:  10
Game: Grand Theft Auto V
Current Players: 64,296
Peak Today: 135,725
Store Link: https://store.steampowered.com/app/271590/Grand_Theft_Auto_V/
Trending Video: 'MEJORANDO de la PEOR a la MEJOR MOTO en GTA 5! 🏍️🛠️ (Mods)' ---> by E-MasterSensei
https://www.youtube.com/watch?v=oGqsbcb7eGA


--------------------------------------------------------------------------------
```

The YouTube videos returned have a language preference for english, but videos in english are not guaranteed.
