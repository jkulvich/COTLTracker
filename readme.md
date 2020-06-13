Tool to automatic play on musical instruments in
"Sky: Children of the Light"

![alt text](./assets/proof2.gif)

YouTube Demo:
[River Flows in You | Horn](https://www.youtube.com/watch?v=-RD3mvBv8M8)
[Sparkle | Piano](https://www.youtube.com/watch?v=9vW_sGyi8EE)
[Zen Zen Zense | Piano](https://www.youtube.com/watch?v=WTTuqxaN5xg)

# Fast Start To Play

## Preparing
You'll need any **android device**, **PC** with any available OS and **USB cable** to connect your PC with your smartphone.
1. [**Enable USB debugging**](https://www.phonearena.com/news/How-to-enable-USB-debugging-on-Android_id53909) on your phone, It is in developer options.
2. [**Install ADB**](https://www.xda-developers.com/install-adb-windows-macos-linux/) on your PC.
3. **Plug your phone** with your PC and accept debug permissions if required.

## Configuring
See latest release with [prebuilds and tracks here](https://github.com/jkulvich/COTLTracker/releases).
1. **Download** one of these **prebuilt binary** app for your PC OS.
2. **Download the archive** with musical tracks for player and unpack it near the app.

## Running
1. On your PC **run a terminal**. If you're on the Windows, just press RMB when Shift pressed and select "Open command window here". You _should be in same directory_ where the app located.
2. **Execute the command**: `player --test` when the game running on your phone. _Don't forget take a musical instrument in your hands!_
3. If step 2 is ok and all tones was well, then just type: `player --track tracks/sparkle.txt` and hit Enter.

You're great!
You can stop the app with CTRL+C.

# Flags

- track - Path to track file
- delay - Delay in ms between taps
- start - Number of block where to start
- test - Runs taps test for all musical instrument buttons

Simple example:
```bash
./player --track ./tracks/zen_zen_zense.txt
```

# Block in tracker file

There are 3 blocks types:
- **Delay** - several presentation types are available:
    - **Number** - milliseconds: `200 500`
    - **tN** - milliseconds in formula 200 * N: `t t5 t10`.        
    - **Dash** - 200 * dash_count: `- -- ---`.        
- **Note** - note in char notation like: `C4, A2 G5`
- **Chord** - chord like: `Am, E, G`

Please, use dash or t for timings, so you can change
the track speed by changing the `timing` comment.

```
#!TIMING:200
#!SHIFT:2
```

# Tone transpose

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

## Prepared tracker files

You can found it in the ./track folder

## Game notes

```
[C1] [D1] [E1] [F1] [G1]
[A1] [B1] {C2} {D2} {E2}
{F2} {G2} {A2} {B2} (C3)
```

```
C1   D1   E1   F1   G1   A1   B1   C2   D2   E2   F2   G2   A2   B2   C3
``` 
