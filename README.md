# 2D Convex Hull API in Go

Given a set of points in 2D space, this API will return a valid Convex Hull polygon using a Gift Wrapping algorithm written in Go.

## Definition

A Convex Hull is calculated from a set of points (with n >= 3) and defines the minimum convex polygon so that any point of the set is either inside or at the border of the polygon.
More informations on [Wikipedia](https://en.wikipedia.org/wiki/Convex_hull).

![Illustration of a Convex Hull from Wikipedia](https://upload.wikimedia.org/wikipedia/commons/thumb/b/bc/ConvexHull.png/220px-ConvexHull.png)

## Why

Convex Hull is often used in [video games](https://www.cs.sfu.ca/~haoz/pubs/liu_zhang_gi08.pdf), allowing to find simple collision polygon from complex shapes (see below sprite). This optimization allows to use [simpler and faster algorithms](https://www.toptal.com/game/video-game-physics-part-ii-collision-detection-for-solid-objects) without impacting user experience.

Furthermore, convex polygons are in some cases a good compromise between collision box simplicity and [collision mesh](https://developer.valvesoftware.com/wiki/Collision_mesh) accuracy.

![Convex Hull of a video game character](https://i.imgur.com/BhKRtRW.png)

*Original sprite from [LuizMelo](https://luizmelo.itch.io/), The Convex Hull is drawn in red, it can be used as a collision polygon.*

Finding such polygon can be usefull in other fields like [shape analysis](https://brilliant.org/wiki/convex-hull/), [linear programming](https://cs.au.dk/~gerth/advising/thesis/bo-mortensen.pdf) and even [cooking](http://veronising.blogspot.com/2008/04/cooking-for-nerds-ingredient-polyhedron.html). 

Here are other usefull links to learn more about applications of Convex Hull:

[Collision detection in video game physics](https://www.toptal.com/game/video-game-physics-part-ii-collision-detection-for-solid-objects)

[Applications of Convex Hull](https://brilliant.org/wiki/convex-hull/)

## How

There are multiple algorithms to determine a convex hull, the one I used in this program is the [Gift Wrapping algorithm](https://en.wikipedia.org/wiki/Gift_wrapping_algorithm), whose time complexity is `O(mn)`. Other algorithms can be found on [this excellent pdf](https://www.cs.jhu.edu/~misha/Spring16/06.pdf).

### Example

For this example, we assume that the points are not collinear. I'm not going to detail the math here, but you can find more details on the code itself, on the [Wikipedia page](https://en.wikipedia.org/wiki/Convex_hull) and on this [video](https://youtu.be/YNyULRrydVI).

I made a tool to have a visual representation of the algorithm (you can find it in the `/public` folder). The screenshots below were taken from this tool.

Given this set of points:

![Step 0](https://i.imgur.com/A8VCrml.png)

The algorithm will first take the leftmost point (point with minimal x-value):

![Step 1](https://i.imgur.com/f0VCoH9.png)

Then rotate an imaginary line around this point until it crosses another point:

![Step 2](https://i.imgur.com/lZOqdeV.png)

By repeating this process until we get back to the starting point, we get the Convex Hull:

![Step 3](https://i.imgur.com/pujWKZi.png)

![Step 4](https://i.imgur.com/lF5lBX9.png)

## The API

The API is opened on port 8080 by default and has two routes:

```GET /``` : A static web page containing a tool to experiment with the API.

```POST /convex2d``` : The actual route that takes a list of points in the body (JSON format) and returns the list of points to construct the 2D Convex Hull (JSON format too).

Example of input:

```json
{ "points": [
    { "x": 239, "y": 308},
    { "x": 305, "y": 241},
    { "x": 377, "y": 321},
    { "x": 300, "y": 365},
    { "x": 302, "y": 301}
]}
```

The API returns:

```json
[
    { "x": 239, "y": 308 },
    { "x": 305, "y": 241 },
    { "x": 377, "y": 321 },
    { "x": 300, "y": 365 }
]
```

![Simple example](https://i.imgur.com/4SkThq6.png)

## How to start the application

If you want to test the API and the tool that comes with it, there are two built version that you can run: one Linux and one Windows version. You can run executables without any requirement, therefore you do **not** need Go to be installed.

The application will start the API on port 8080, and will close it when the application is terminated.

### Linux
```bash
# Clone repo
git clone https://github.com/BaptisteMiq/convex-hull-go.git
cd convex-hull-go

# Start with
./convex-hull
```

### Windows
```batch
git clone https://github.com/BaptisteMiq/convex-hull-go.git
cd convex-hull-go
start convex-hull.exe
```
or

Download ZIP and start `convex-hull.exe`.

## If you want to build from source...
*This program was written in Go 1.13.*

You will first need to clone the repository:

```bash
https://github.com/BaptisteMiq/convex-hull-go.git
cd convex-hull-go
```

Then install the required Go package (gin):
```bash
# You might need this library
sudo apt-get install libxi-dev

go get .
```

You can start the API with:
```bash
go run main.go
```

To build the app:
```bash
env GOOS=linux go build -trimpath   # Linux
env GOOS=windows go build -trimpath # Windows
```


## The tool

This tool can be accessed at http://localhost:8080 (when launched locally).

Click on the canvas to add a point at mouse position. The API will be called automatically and the Convex Hull will be drawn.

![GIF demo of the tool](https://media.giphy.com/media/RyWpwFci7Xn8Wfuhgd/giphy.gif?cid=790b7611eb2a149d17487521cf0772fb9951650d861fed2d&rid=giphy.gif&ct=g)

## Author

Made by Baptiste Miquel under the MIT license.