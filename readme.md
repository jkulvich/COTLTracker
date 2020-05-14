Tool to automatic play on musical instruments in
"Sky: Children of the Light"

![alt text](./assets/proof2.gif)

# Requirements

- Android device with USB debugging
- Linux and ADB tools

# Flags

- track - Path to track file
- delay - Delay in ms between taps
- start - Number of block where to start

Simple example:
```
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