# Ebiten Beginner Guide

So you want to do game development in Go?

This document describes one of the paths you can take.

> This tutorial uses `bash` as a way to describe some actions in a programmer-friendly way. Keep in mind that these commands are not meant to be copy/pasted.

## Step 0 - Joining the Community

Join the official [Ebitengine Discord server](https://discord.com/invite/3tVdM5H8cC). If you ever get stuck, you can find help there.

## Step 1 - Running the Examples

```bash
$ git clone --depth 1 https://github.com/hajimehoshi/ebiten.git
$ cd ebiten
```

> You may need to [install the dependencies](https://ebitengine.org/en/documents/install.html) required by Ebitengine.

The examples are located under the `examples` folder. We'll start by running some of the examples.

```bash
# Run the game, check it out, then close the application and run another one.
$ go run ./examples/blocks
$ go run ./examples/flappy
$ go run ./examples/2048
$ go run ./examples/snake
```

There are ~80 examples in the folder. Most of them show how to solve one particular problem. Since they rarely use any third-party dependencies, you might expect there to be some boilerplate.

> Explore the examples, and look at the code. Try changing some parts and see how it goes.

The assets used for these examples can be found in `examples/resources` folder.

## Step 2 - Make a Tiny Game

I know that you want to create your dream game right away, but it might be impractical at this point. Try to make a simple game that would take you a couple of weeks to finish.

Here is a list of game ideas:

* Asteroids
* Space Invaders
* Snake

I don't recommend making a platformer or any other game that would require complicated physics. Take something that can make you comfortable with the engine. You'll explore the third-party libraries along the way.

Speaking of libraries...

## Step 3 - Explore the Libraries

The best place to start your search is [awesome-ebitengine](https://github.com/sedyh/awesome-ebitengine). That list is not comprehensive, but at least it's well-maintained.

Ebitengine implements a lot of good stuff, but many games will need some extra libraries if you want to avoid writing that code yourself.

Here is a list of things you might want to get sooner or later:

* [UI library](https://github.com/ebitenui/ebitenui)
* [Advanced input handling library](https://github.com/quasilyte/ebitengine-input)
* [Resource manager](https://github.com/quasilyte/ebitengine-resource)
* [2D vector math library](https://github.com/quasilyte/gmath)
* [Path finding library](https://github.com/quasilyte/pathing/)

> TODO: this list is far from being complete.

If you don't like some particular library, there are almost always several alternatives.

With examples and libraries, you should be able to finish your first game. If it's still too challenging, take one of the games from the examples and start modifying it. Add a menu, a couple of levels, and maybe new gameplay mechanics.

## Step 4 - Getting Serious

Your toolbox should evolve to make bigger games. Here is a list of stuff I find very useful:

* Visual level editors: [Tiled](https://www.mapeditor.org/) or [LDtk](https://ldtk.io/)
* Sprite editor: [aseprite](https://www.aseprite.org/)
* Sound effects generator: [sfxr](https://www.drpetter.se/project_sfxr.html) (there are other versions on the web)

If you want to get some free music for your game, take a look at [modarchive](https://modarchive.org/). Just make sure to give the appropriate credit for the author (pay attention to the license of the track). You can either convert the modular music files to `ogg` or use the [xm player](https://github.com/quasilyte/xm) to play them directly.

Check the [games](https://github.com/sedyh/awesome-ebitengine?tab=readme-ov-file#games) section of awesome-ebitengine to find some inspiration (these games have their sources published on GitHub). There is also [made-with-ebitengine](https://itch.io/queue/c/2030581/made-with-ebitengine) collection on `itch.io`.

There are tons of stuff to learn from this point onward. Step by step, you'll get better.
