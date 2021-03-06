# Need For Speed

Welcome to Need For Speed on Exercism's Go Track.
If you need help running the tests or submitting your code, check out `HELP.md`.
If you get stuck on the exercise, check out `HINTS.md`, but try and solve it without using those first :)

## Introduction

In Go, a `struct` is a sequence of named elements called _fields_, each field having a name and type. The name of a field must be unique within the struct. `Structs` can be compared with the _class_ in the Object Oriented Programming paradigm.

You create a new struct by using the `struct` keyword, a **_built-in type_**, and explicitly define the name and type of the fields as shown in the example below.

```go
type StructName struct{
    field1 fieldType1
    field2 fieldType2
}
```

Struct fields are accessed using a `.` notation.

```go
someStruct.someField = "someValue"
```

## Instructions

In this exercise you'll be organizing races between various types of remote controlled cars. Each car has its own speed and battery drain characteristics.

Cars start with full (100%) batteries. Each time you drive the car using the remote control, it covers the car's speed in meters and decreases the remaining battery percentage by its battery drain.

If a car's battery is below its battery drain percentage, you can't drive the car anymore.

Each race track has its own distance. Cars are tested by checking if they can finish the track without running out of battery.

## 1. Creating a remote controlled car

Allow creating a remote controller car by defining a function `NewCar` that takes the speed of the car in meters, and the battery drain percentage as its two parameters (both of type `int`) and returns a `Car` instance:

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
// Output: Car{speed: 5, batteryDrain: 2, battery:100}
```

## 2. Creating a race track

Define another struct type called `Track` with the field `distance` of type integer.
Allow creating a race track by defining a function `NewTrack` that takes the track's distance in meters as its sole parameter (which is of type `int`):

```go
distance := 800
raceTrack := NewTrack(distance)
// Output: Track{distance: 800}
```

## 3. Drive the car

Implement the `Drive` function that updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage:

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
car = Drive(car)
// Output: Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
```

## 4. Check if a remote control car can finish a race

To finish a race, a car has to be able to drive the race's distance. This means not draining its battery before having crossed the finish line. Implement the `CanFinish` function that takes a `Car` and a `Track` instance as its parameter and returns `true` if the car can finish the race; otherwise, return `false`:

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

distance := 100
raceTrack := NewTrack(distance)

CanFinish(car, raceTrack)
// Output: true
```

## Source

### Created by

- @tehsphinx

### Contributed to by

- @oanaOM