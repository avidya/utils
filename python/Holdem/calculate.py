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

def checkValid(*decks):
    if len(decks) <= 1:
        raise Exception("At least two decks should be specified.")
    if len(set(map(lambda x:len(x), decks))) != 1:
        raise Exception("cards number in every deck should be aligned.")

ALL_CARDS=set(map(lambda t: t[0]+t[1], product(ALL_SUITS, ALL_RANK)))

def calSingle(jobs, pickFunc=fastPickHigh):
    scores=list(map(lambda deck:Stat(deck), jobs[1:]))
    for s in scores:
        s.Deck = pickFunc(list(chain(s.deck, jobs[0])))
    return scores[0].deck if scores[0].Deck > scores[1].Deck else ("draw game" if scores[0].Deck == scores[1].Deck else scores[1].deck)

def calRates(*decks, concurrency=True, pickFunc=fastPickHigh):
    checkValid(*decks)
    result=None
    if concurrency:
        with Pool(multiprocessing.cpu_count()) as p:
            result=Counter(map(lambda r:str(r), \
                p.map_async(calSingle, map(lambda d:(d, *decks), (combinations(ALL_CARDS - set(chain(*decks)), 5)))).get())) 
    else:
        result=Counter(map(lambda r:str(r), map(lambda d:calSingle((d, *decks), pickFunc), (combinations(ALL_CARDS - set(chain(*decks)), 5)))))
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
    calRates(['C2', 'S2'],["HA", 'CT'], pickFunc=pickHigh)
    print("Cost: " + "{:.3f}".format(x+time.time()) +"s")