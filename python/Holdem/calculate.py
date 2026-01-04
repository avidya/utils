#!/usr/bin/python3
#
# author: kozz
# date: 2024/09
##
from core import *
from itertools import product
from itertools import chain
import time
from multiprocessing import Pool
import multiprocessing
from collections import Counter
import random

def checkValid(*decks):
    if len(decks) <= 1:
        raise Exception("At least two decks should be specified.")
    if len(set(map(lambda x:len(x), decks))) != 1:
        raise Exception("cards number in every deck should be aligned.")
    # Check for duplicate cards across all decks
    all_cards = list(chain(*decks))
    if len(all_cards) != len(set(all_cards)):
        duplicates = list(filter(lambda card: all_cards.count(card) > 1, set(all_cards)))
        raise Exception("Duplicate cards found across decks: " + str(duplicates))
    # Validate each card format
    invalid_cards = list(filter(lambda card: not checkValidCard(card), all_cards))
    if invalid_cards:
        raise Exception("Invalid card format found: " + str(invalid_cards))

ALL_CARDS=set(map(lambda t: t[0]+t[1], product(ALL_SUITS, ALL_RANK)))

def calSingle(jobs, pickFunc=fastPickHigh):
    scores=list(map(lambda deck:Stat(deck), jobs[1:]))
    for s in scores:
        s.Deck = pickFunc(list(chain(s.deck, jobs[0])))
    max_deck=reduce(lambda s, a: s if s.Deck >= a.Deck else a, scores)
    winners=list(filter(lambda s: s.Deck == max_deck.Deck, scores))
    return winners[0].deck if len(winners) == 1 else tuple(map(lambda w: w.deck, winners))

def calRates(*decks, concurrency=True, pickFunc=fastPickHigh, sample_factor=None):
    checkValid(*decks)
    result=None
    all_combinations = list(combinations(ALL_CARDS - set(chain(*decks)), 5))
    if sample_factor:
        all_combinations = random.sample(all_combinations, max(1, int(len(all_combinations) * sample_factor)))
    if concurrency:
        with Pool(multiprocessing.cpu_count()) as p:
            result=Counter(map(lambda r:str(r), \
                p.map_async(calSingle, map(lambda d:(d, *decks), all_combinations)).get())) 
    else:
        result=Counter(map(lambda r:str(r), map(lambda d:calSingle((d, *decks), pickFunc), all_combinations)))
    count=reduce(lambda s,a: s+a, result.values())
    for k, v in map(lambda x:(x, "win rates: {:.2f}%".format(x[1]/count*100)), result.items()):
        print(k, v)  

class Stat:
    def __init__(self, deck):
        self.deck = deck
        self.score = 0
        self.Deck = None

if __name__ == '__main__':

    # l=list(combinations(set(map(lambda t: t[0]+t[1], product(ALL_SUITS, ALL_RANK))), 7))[:80000]
    # x=-time.time()
    # for comb in l:
    #     fastPickHigh(comb)
    # print("FastPickHigh Cost: " + "{:.3f}".format(x+time.time()) +"s")

    # x=-time.time()
    # for comb in l:
    #     pickHigh(comb)
    # print("PickHigh Cost: " + "{:.3f}".format(x+time.time()) +"s") 

    x=-time.time()
    calRates(['C2', 'S2'],['HK', 'CK'], ['DK', 'SK'],  pickFunc=pickHigh, sample_factor=0.1)
    print("Cost: " + "{:.3f}".format(x+time.time()) +"s")