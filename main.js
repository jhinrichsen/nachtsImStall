// A simulation of a game i'm playing with my kids.

'use strict'

const I = require('immutable')

// There are 3 cows, sheeps, cats, and pigs each
const nAnimals = 3 * 4,

  // 5 portions of hay
  nHay = 5,

  // Roll a regular dice, return [1..n]
  roll = n => Math.floor(1 + Math.random() * n),

  // Immutable state for all step functions

  // 'Cock' will wake up all animals
  stepCock = state => {
    let c = Object.assign({}, state)
    c.awake = nAnimals
    return c
  },

  // 'Hay' will reduce the number of hays by one
  stepHay = state => {
    let c = Object.assign({}, state)
    c.hay = c.hay - 1
    return c
  },

  // 'Moon' will let one animal fall asleep
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
  isWon = state => state.awake === 0,

  // The game is lost if there is no more hay
  isLost = state => state.hay < 0,

  newGame = {
    // When starting, all animals are awake
    awake: nAnimals,

    // All hay portions are available
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

  // Reduce game result into won/ lost accumulator
  statsReducer = (r, n) => n
    ? {won: r.won + 1, lost: r.lost}
    : {won: r.won, lost: r.lost + 1},

  // Large numbers of games to play will eventually blow call stack when
  // using plain Arrays' apply() and map(). Immutables laziness to the rescue!
  total = 1e7,

  stats = I.Range(1, total)
    .map(oneGame)
    .reduce(statsReducer, {won: 0, lost: 0})

console.log(`Winning ${stats.won} out of ${total} games.`)

// EOF
