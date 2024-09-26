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

def cal(tasks, *decks):
    scores=list(map(lambda deck:Stat(deck), decks))
    count=0
    for c in tasks:
        for s in scores:
            s.Deck = fastPickHigh(list(chain(s.deck, c)))
        sorted(scores, key=lambda s:s.Deck, reverse=True)[0].score+=1
        count+=1
        if count % 10000 == 0:
            print("processed: " + str(count))
    return count, scores


def trivialCal(*decks):
    checkValid(*decks)
    count, scores = cal(combinations(ALL_CARDS - set(chain(*decks)), 5), *decks)
    for s in scores:
        print(str(s.deck) + " : " + "{:.2f}".format(s.score/count*100) + "%")

def calSingle(job):
    scores=list(map(lambda deck:Stat(deck), job[1:]))
    for s in scores:
        s.Deck = fastPickHigh(list(chain(s.deck, job[0])))
    return sorted(scores, key=lambda s:s.Deck, reverse=True)[0].deck

def concurrentCal(*decks):
    checkValid(*decks)
    result=None
    with Pool(multiprocessing.cpu_count()) as p:
        result=Counter(map(lambda r:tuple(r), \
            p.map_async(calSingle, map(lambda d:(d, *decks), (combinations(ALL_CARDS - set(chain(*decks)), 5)))).get()))
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
    concurrentCal(["HA", 'CT'], ['C2', 'S2'])
    # trivialCal(["HA", 'CT'], ['C2', 'S2'])
    print("Cost: " + "{:.3f}".format(x+time.time()) +"s")