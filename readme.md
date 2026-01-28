[![GitHub](https://img.shields.io/github/license/jkulvich/cotltracker)](https://github.com/jkulvich/COTLTracker/blob/master/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jkulvich/cotltracker)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/jkulvich/cotltracker)](https://github.com/jkulvich/COTLTracker/releases)
[![GitHub issues](https://img.shields.io/github/issues/jkulvich/cotltracker)](https://github.com/jkulvich/COTLTracker/issues)
[![GitHub last commit](https://img.shields.io/github/last-commit/jkulvich/cotltracker)](https://github.com/jkulvich/COTLTracker/commits/master)

> **Note**
> 
> Dear users, thank you for your stars and contributions. It's really so important for me. This project was one of first which goal automate musical instrument playing for Sky: Children of The Lights. After time, we have several better alternatives (see links). And I have no time to maintain this tool anymore. It's sad, but this tool isn't actual anymore.
>
> Thank you all for your support!  
> You all are breathtaking ;)
> 
> Special thanks to [@MapleStudio](https://www.youtube.com/@MapleStudio) for my inspiration and great tool to learn and play COTL instruments [SkyStudio](https://play.google.com/store/apps/details?id=com.Maple.SkyStudio)

> **Warning**
> 
> - Stable work guaranteed only if you **play on Android** ~~and you **have a Linux (Ubuntu based) device**~~
> - ~~For **Windows** users it's possible to **run Linux inside virtual machine** (like Virtual Box)~~
> - ~~For **Windows** users it's not possible to run this tool under CMD, PowerShell or WSL (due WSL USB restrictions)~~
> - At the moment, this project **have not active maintenance**, use similar tools like:
>     - [Sky Auto Music: Clicker Studio](https://play.google.com/store/apps/details?id=com.zhukovartemvl.skyautomusic)
>     - [Sky Music Studio: Auto](https://play.google.com/store/apps/details?id=com.edegrangames.skyMusic)

> :sparkles: COTLTracker can run Sky Music Sheet Maker format!
> 
> :sparkles: COTLTracker now works on Windows, Mac, and Linux without ADB installed!

# :joystick::notes: COTLTracker :: Players' assistant
"Sky: Children of The Light" musical assistant tool to automatic play on in-game musical instruments

![Tool Proof](./assets/proof2.gif)

:eyes: YouTube Demos:
1. [Love Scenario | Piano](https://youtu.be/ejYJq7mixME)
2. [Girls Like You | Piano](https://youtu.be/8W7AQtnZh0k)
3. [Counting Stars | Piano](https://youtu.be/JMDFZYuwwz8)
4. [Way Back Home | Piano](https://youtu.be/OMZEtMOoTOI)
5. [River Flows in You | Horn](https://www.youtube.com/watch?v=-RD3mvBv8M8) - :-1: Bad Record
6. [Sparkle | Piano](https://www.youtube.com/watch?v=9vW_sGyi8EE) - :-1: Bad Record
7. [Zen Zen Zense | Piano](https://www.youtube.com/watch?v=WTTuqxaN5xg) - :-1: Bad 

My source of inspiration for music is [Maple on YouTube](https://www.youtube.com/channel/UCDckPUJKSo9UeVtlY31p3Ag)

# :v: Contact info
For legal issues, tech questions and chatting.  
Please, feel free to text me anytime :)

- [Telegram | @jkulvich](https://t.me/jkulvich) - :star: Priority
- [Instagram | @ijkulvich](https://instagram.com/ijkulvich)
- [Twitter | @jkulvich](https://twitter.com/jkulvich)
- [VKontakte | @jkulvich](https://vk.com/jkulvich)
- [EMail | jkulvichi@gmail.com](mailto:jkulvichi@gmail.com)

# :fast_forward: Fast Start To Play

## :iphone::left_right_arrow::computer: Prepare it
Please, prepare your android phone and plug it with your PC.  
**IMPORTANT**: Using of the tool suitable only for android players.
1. [**Enable USB debugging**](https://www.phonearena.com/news/How-to-enable-USB-debugging-on-Android_id53909) on your phone, It is in developer options.
~~2. [**Install ADB**](https://www.xda-developers.com/install-adb-windows-macos-linux/) on your PC.~~
    ~~1. If you are a Windows user, make sure that [**ADB is in %PATH% variable**](https://nerdschalk.com/set-adb-fastboot-path-windows)~~
3. **Plug your phone** with your PC and accept debug permissions if required.

## :package::arrow_down: Configure it
See latest release with [prebuilds and tracks here](https://github.com/jkulvich/COTLTracker/releases).
1. **Download** one of these **prebuilt binary** app for your PC OS.
2. **Download an archive** with musical tracks for player and unpack it near the app.

> Installation folder does not matter. You can drop prebuilt binary and tracks folder in any folder.

## :computer::arrow_forward: Run it
1. **Run a terminal** on your PC. If you're a Windows user, just press RMB when Shift pressed and select "Open command window here" or "Open PowerShell here". You _should be in same directory_ where the app located.
2. **Execute next command**: `player --test` when the game running on your phone. _Don't forget take a musical instrument in your hands!_
3. **Run a lovely track**: When the short test passed type: `player --track tracks/sparkle.txt` and press Enter.

You can stop the app by CTRL+C.

## :musical_note: Prepared tracker files

You can find it in the [tracks folder](./tracks)

# :checkered_flag: Flags

- track - Path to track file (musical file)
- delay - Delay in ms between taps (default is 80, increment it if your device can't catch all tones)
- start - Number of block where to start (default is 0)
- test - Run taps test for all musical instrument buttons (Check it before real usage)

Simple example:
```bash
./player --track ./tracks/zen_zen_zense.txt
```

# :1234: Block in tracker file

There are 3 blocks' types:
- **Delay** - several presentation types are available:
    - **Number** - milliseconds: `200 500` - :-1: Deprecated
    - **tN** - 200ms * N: `t t5 t10` - :-1: Deprecated
    - **Dash** - 200ms * dash_count: `- -- ---` - :star: Modern Variant
- **Note** - note in char notation like: `C4 A2 G5`
- **Chord** - chord like: `Am E G` (Not all notes available, use **Note** instead) - :-1: Deprecated

Please, use dash or t for timings, so you can change
the track speed by changing the `timing` comment.

```
#!TIMING:200
#!SHIFT:2
```

# :arrow_up_small::arrow_down_small: Tones transposing

The tool has an auto transpose mechanism.
So, if you are using `A4 B5 C4` then all notes for 4 octaves will be
transposed to 1 octave
```
[C4] [  ] [  ] [  ] [  ]
[A4] [  ] {  } {  } {  }
{  } {  } {  } {B5} (  )
```

If you want to use only second octave then make a `shift` comment
with value `7`. So, you'll shift all notes for 7 tones.

For example, without the comment `A4` position will look like:
```
[  ] [  ] [  ] [  ] [  ]
[A4] [  ] {  } {  } {  }
{  } {  } {  } {  } (  )
```
With the comment:
```
[  ] [  ] [  ] [  ] [  ]
[  ] [  ] {  } {  } {  }
{  } {  } {A4} {  } (  )
```

# :information_source: Game notes

```
[C1] [D1] [E1] [F1] [G1]
[A1] [B1] {C2} {D2} {E2}
{F2} {G2} {A2} {B2} (C3)
```

```
C1   D1   E1   F1   G1   A1   B1   C2   D2   E2   F2   G2   A2   B2   C3
``` 
