"use strict"

// A simulation of a game i' playing with my kids.

// There are 3 cows, sheeps, cats, and pigs each
const nAnimals = 3 * 4,

  // 5 portions of hay
  nHay = 5,

  // Roll a regular dice, return [1..n]
  roll = n => Math.floor(1 + Math.random() * n),

  // The cock will wake up all animals
  // Immutable state
  stepCock = state => {
    let c = Object.assign({}, state)
    c.awake = nAnimals
    return c
  },

  // Hay will reduce the number of hays by one
  stepHay = state => {
    let c = Object.assign({}, state)
    c.hay = c.hay - 1
    return c
  },

  // One animal falls asleep
  stepMoon = state => {
    let c = Object.assign({}, state)
    c.awake = c.awake - 1
    return c
  },

  // Return step function based on random dice
  stepFn = _ => [
      stepHay,
      stepCock,
      stepMoon, stepMoon, stepMoon, stepMoon
    ][roll(6) - 1],
    
  // The game is won if all animals are asleep
  isWon = state => state.awake == 0,

  // The game is lost if there is no more hay
  isLost = state => state.hay < 0,

  newGame = {
    // When starting, all animals are awake
    awake: nAnimals,

    // A couple of hay portions
    hay: nHay
  },

  // Play one game, return true if won
  oneGame = _ => {
    let state = newGame
    while (true) {
      // A game step
      var fn = stepFn()
      // Apply step function state
      state = fn(state)
      if (isWon(state)) return true
      if (isLost(state)) return false
    }
  },

  total = 100000,
  games = Array.apply(null, Array(total)).map(oneGame),
  wins = games.filter(b => b).length
  
console.log(`Played ${total} games winning ${wins}`)

// EOF