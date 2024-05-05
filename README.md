# Kryptos K4 LLM Guesser

This is a tool that uses LLMs to attempt to guess the plaintext of K4 from the Kryptos sculpture. It just generates snippets of text that fit 1:1 in the unknown sections of K4. This is **not a serious attempt** to solve K4, but rather a fun way to generate something that looks like it could be a solution.

Technically, this is a template-based text generator. It can be used to generate text that fits a given template (even if the template is not related to Kryptos). However, some of the design choices were made with K4 in mind, like the ignoring of spaces and punctuation.

## Template format

```
Is this the {4} life? {6} just fantasy?
```

This will attempt to generate text that fits the template, where `{4}` and `{6}` are placeholders for the generated text. The number inside the curly braces is the length of the generated text, without accounting for spaces (`Is this` is 7 characters long, but only 6 if we ignore the space).

## Usage

Before running the project, a [llama.cpp](https://github.com/ggerganov/llama.cpp/) must be accessible. Any LLM that can generate text can be used, but optionally, an LLM that can do FIM (Fill in the middle) can be used, which may generate more coherent text as it can also leverage the context of the known parts.

```bash
  go install
```

```bash
  go run ./cmd/gen_correct/main.go --help
```

```
  -fill
        Fill in the middle (only works if the model supports it)
  -jobs int
        Number of concurrent LLM jobs (default 1)
  -min float
        Minimum probability to consider (default 0.01)
  -template string
        Template to use
  -url string
        URL of the llama.cpp server (default "http://localhost:8080/")
```

- `-fill` Can be used if the LLM supports FIM (Fill in the middle), which can generate more coherent text.
- `-jobs` Number of concurrent LLM jobs to run, it will make more requests in parallel.
- `-min` Minimum probability to consider, lower values will generate more text, but it may be less coherent.
- `-template` The template to use, see the format above.
- `-url` The URL of the llama.cpp server, by default it will use `http://localhost:8080/`. 

## Example

```bash
  go run ./cmd/gen_correct/main.go -template "Mama, I just {10}. Put a {3} against his head, pulled my {7} now he's {4}."
```

```
killed a man. Put a gun against his head, pulled my trigger now he's dead.
want to be in. Put a pin against his head, pulled my trigger now he's dead.
want to be an. Put a cap against his head, pulled my trigger now he's dead.
want to be in. Put a new against his head, pulled my hand out now he's gone.
want to be in. Put a new against his head, pulled my hand out now he's dead.
want to be in. Put a new against his head, pulled my hand out now he's been.
want to be in. Put a tag against his head, pulled my hair out now he's dead.
want to be in. Put a tag against his head, pulled my hair and now he's gone.
want to be in. Put a tag against his head, pulled my hair and now he's dead.
want to be an. Put a hat against his head, pulled my hand and now he's gone.
want to be an. Put a hat against his head, pulled my hand and now he's dead.
want to be an. Put a hat against his head, pulled my handker now he's dead.
want to be an. Put a hat against his head, pulled my hand out now he's dead.
...
```

> Results will vary depending on the LLM used.

## Kryptos Example

```
go run ./cmd/gen_correct/main.go -template "k1: between subtle shading and the absence of light lies the nuance of iqlusion\n\nk3: slowly desparatly slowly the remains of passage debris that encumbered the lower part of the doorway was removed with trembling hands i made a tiny breach in the upper left hand corner and then widening the hole a little i inserted the candle and peered in the hot air escaping from the chamber caused the flame to flicker but presently details of the room within emerged from the mist x can you see anything q\n\nk4:{21} east northeast {29} berlin clock {23}" -fill
```

```
 i can see a large room i see east northeast and northwest walls and a large alt berlin clock on the wall to the northwest
 i can see a large room with east northeast and northwest walls and a door to my berlin clock in the northwest wall and an
 i can see a large room with east northeast and northwest walls and a door to my berlin clock in the northwest wall and to
 i can see a large room and a east northeast corner of a room with a doorway to a n berlin clock tower and a large room with a
 i can see a large room i see east northeast and northwest walls and a large oak berlin clock on the wall to the northwest
 i can see a large room i see east northeast and northwest walls and a large alt berlin clock on the wall to the north a man
 i can see a large room with east northeast and northwest walls and a door to my berlin clock in the northwest wall and in
 i can see a large room with east northeast and northwest walls and a door to my berlin clock in the northwest wall and on
 i can see a large room and i east northeast of the room is a doorway and i can see berlin clocks and a large clock on the far
 i can see a large room and a east northeast corner of a room with a doorway to it berlin clocks are ticking in the room and
 i can see a large room and a east northeast corner of a room with a doorway to a n berlin clock tower and a large room to the
 i can see a large room and a east northeast corner of a room with a doorway to my berlin clocks and a large room with a fire
 i can see a large room and a east northeast corner of a room with a doorway to my berlin clocks and a large room with a door
 i can see a large room and a east northeast corner of a room with a doorway to an berlin clock and a large room with a table
 i can see a large room and a east northeast corner of a room with a doorway to an berlin clock and a large room with a large
 i can see a large room i see east northeast and northwest walls and a large oak berlin clock on the wall to the north a few
 i can see a large room i see east northeast and northwest walls and a large oak berlin clock on the wall to the north a des
 i can see a large room i see east northeast and northwest walls and a large alt berlin clock on the wall to the north a des
 i can see a large room i see east northeast and northwest walls and a large alt berlin clock on the wall to the north a few
 i can see a large room i see east northeast and northwest walls and a large alt berlin clock on the wall to the north a red
 i can see a large room i can east northeast and northwest and a small room to my berlin clockwise from the doorway i am in
 i can see a large room with east northeast and northwest walls and a door at nw berlin clockwise from the door is a small
 i can see a large room with east northeast and northwest walls and a door at nw berlin clockwise from the door is a large
 i can see a large room with east northeast and northwest walls and a door to my berlin clock in the northwest wall the ce
 i can see a large room and a east northeast corner of a room with a doorway to it berlin clocks are ticking in the room i do
 i can see a large room and a east northeast corner of a room with a doorway to it berlin clocks are ticking in the room i am
 i can see a large room and a east northeast corner of a room with a doorway to it berlin clocks are ticking in the room the
 i can see a large room and a east northeast corner of a room with a doorway to my berlin clocks and a large room with a long
 i can see a large room and a east northeast corner of a room with a doorway to my berlin clocks and a large room with a west
 i can see a large room and a east northeast corner of a room with a doorway to an berlin clock and a large room with a stove
 i can see a large room and a east northeast corner of a room with a doorway to an berlin clock and a large room with a stair
 i can see a large room and a east northeast corner of a room with a doorway to an berlin clock and a large room with a stain
 i can see a large room and a east northeast corner of a room with a doorway in it berlin clocks are ticking and the room sm
 i can see a large room and a east northeast corner of a room with a doorway in it berlin clocks are ticking and the room is
 i can see a large room i see east northeast and northwest walls and a large oak berlin clock on the wall to the north a man
 i can see a large room i am a east northeast of the doorway and the room is lit in berlin clockwork light the walls are of d
 i can see a large room i am a east northeast of the doorway and the room is lit in berlin clockwork light the walls are of a
 ...
```

Including the different already solved parts of Kryptos will yield different results, for example, if K2 is included, the results will use `x` a lot more.

```
go run ./cmd/gen_correct/main.go -template "between subtle shading and the absence of light lies the nuance of iqlusion\n\nit was totally invisible hows that possible ? they used the earths magnetic field x the information was gathered and transmitted undergruund to an unknown location x does langley know about this ? they should its buried out there somewhere x who knows the exact location ? only ww this was his last message x thirty eight degrees fifty seven minutes six point five seconds north seventy seven degrees eight minutes forty four seconds west x layer two\n\nslowly desparatly slowly the remains of passage debris that encumbered the lower part of the doorway was removed with trembling hands i made a tiny breach in the upper left hand corner and then widening the hole a little i inserted the candle and peered in the hot air escaping from the chamber caused the flame to flicker but presently details of the room within emerged from the mist x can you see anything q\n\n{21} east northeast {29} berlin clock {23}" -fill
```

```
i can see the floor and the east northeast wall x the ceiling is covered with a berlin clock x the west wall is covered in
i can see the floor and the east northeast wall x the ceiling is covered with a berlin clock x the west wall is covered by
i can see the floor and and east northeast wall x the ceiling is covered with a berlin clock x the west wall is covered in
i can see the floor and and east northeast wall x the ceiling is covered with a berlin clock x the west wall is covered by
...
```

Or you can just generate text for the unknown parts of K4.

```
go run ./cmd/gen_correct/main.go -template "{21} east northeast {29} berlin clock {23}"
```

```
The first thing that I did east northeast of the intersection of the two dirt berlin clock tower roads was to take a few
The first thing that I did east northeast of the intersection of the roads is berlin clock tutorial pdf the clock’s
The first thing that I did east northeast of the intersection of the roads is berlin clock tutorial pdf the clock and I
The first thing that I did east northeast of the intersection of the roads is berlin clock tutorial pdf the clock face
The first thing that I did east northeast of the intersection of the two main berlin clock roads was to take a photo of a
The first thing that I did east northeast of the intersection of the roads is berlin clock tutorial pdf the clock on my
The first thing that I did east northeast of the intersection of the roads in berlin clock tower square was to take the
The first thing that I did east northeast of the intersection of the roads in berlin clock tower square was to take out
The first thing that I did east northeast of the intersection of the roads to berlin clock tower and the old mill was to
The first thing that I did east northeast of the intersection of the roads at berlin clock tower was to take a photo of a
The first thing that I did east northeast of the intersection of the roads in berlin clock tower square was to look for
The first thing that came east northeast of the city of the city of the city is berlin clock tower in the city of the city
The first thing that came east northeast of the city of the city of the east of berlin clock tower is the east northeast
The first thing that I did east northeast of the intersection of the two dirt berlin clock tower roads was to take some
The first thing that I did east northeast of the intersection of the roads in berlin clock tower square was to look at a
The first thing that I did east northeast of the intersection of the east and berlin clock towers was to take a photo in
The first thing that I did east northeast of the intersection of the east and berlin clock towers was to take a photo of
The first thing that I did east northeast of the intersection of the east and berlin clock towers was to take a picture
The first thing that I did east northeast of the intersection of the two main berlin clock roads was to take a good long
The first thing that I did east northeast of the intersection of the two main berlin clock roads was to take a good look
The first thing that I did east northeast of the intersection of the two main berlin clock roads was to take a good walk
The first thing that I did east northeast of the intersection of the two main berlin clock roads was to take a good hard
The first thing that I did east northeast of the intersection of the roads is berlin clock tutorial pdf the clock in my
The first thing that I did east northeast of the intersection of the east and berlin clock towers was to take a few deep
The first thing that came east northeast of the city of the city of the east of berlin clock tower is the east of the east
The first thing that came east northeast of the city of the city of the east of berlin clock tower is the east of the city
The first thing that I did east northeast of the intersection of the road to H berlin clock tower and the road to the old
The first thing that I did east northeast of the intersection of the road to P berlin clock tower and the road to the old
The first thing that I did east northeast of the intersection of the road to S berlin clock tower and the road to the old
The first thing that I did east northeast of the intersection of the road to M berlin clock tower and the road to the old
The first thing that I did east northeast of the intersection of the road to L berlin clock tower and the road to the old
The first thing that I did east northeast of the intersection of the road and berlin clock tower was to take a photo of a
The first thing that came east northeast of the city of the city of the city of berlin clock tower is the first time that
...
```

## Conclusion

As you can probably see, this isn't very useful for solving K4.
