#!/usr/bin/python3
#
# author: kozz
# date: 2024/09
##
from functools import reduce
from functools import total_ordering
from itertools import combinations
from itertools import dropwhile
from collections import Counter

ALL_SUITS=['H', 'S', 'C', 'D']
ALL_RANK=['2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A']
#ALL_RANK=['2', 'T', 'J', 'Q', 'K', 'A']


def decode(s):
    if s == 'T':
        return 10
    elif s == 'J':
        return 11
    elif s == 'Q':
        return 12
    elif s == 'K':
        return 13
    elif s == 'A':
        return 14
    elif s in ['2','3','4','5','6','7','8','9']:
        return int(s)
    else:
        raise Exception("invalid input: " + s)

##
# every card in a deck noted as 'XY'.
# the 'X' must be one of the four `suits``: `H` represent 'Heart', `S` represent 'Shade', `C` represent 'Club' and `D` represent 'Diamond'
# the valid 'Y' must be a `rank`` from ['2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A']
##
def checkValidCard(card):
    return isinstance(card, str) and len(card) == 2 and \
        card[0] in ALL_SUITS and card[1] in ALL_RANK

def checkValidDeck(deck):
    if len(deck) != 5:
        raise Exception("deck number must be an array with five cards: " + str(deck))
    if len(set(deck)) != 5:
        raise Exception("duplicated cards found: " + str(deck))
    if not reduce(lambda s, a: s and checkValidCard(a), deck, True):
        raise Exception("invalid card style: " + str(deck))

## 
# grouped structure:
# original deck: ['H2', 'HT', 'CT', 'S5', 'ST']
# after grouping, turned into: [['HT', 'CT', 'ST'], ['S5'], ['H2']]
##
def groupByRank(deck):
    groupedDeck={}
    for card in deck:
        if card[1] in groupedDeck:
            groupedDeck[card[1]].append(card)
        else:
            groupedDeck[card[1]]=[card]
    return sorted(groupedDeck.values(), key=lambda i:decode(i[0][1])+len(i)*100, reverse=True)

_pick_from_groupedDeck=lambda groupedDeck:list(map(lambda a:a[0], groupedDeck))

_make_judge=lambda deck_pattern, rank, aux=lambda x:True:\
    lambda groupedDeck: Deck(rank, _pick_from_groupedDeck(groupedDeck)) \
            if aux(groupedDeck) and list(map(lambda a:len(a), groupedDeck)) == deck_pattern \
            else None

## judge func generator. \n
# we create five judger(`four of a kind`, `full house`, `three of a kind`, `two pair` and `one pair`) sequentially
# when any of such judge hit, a `Deck` will be instantiated and returned immediately, and this is why they be called 'trivial'.
_trivial_judge=list(map(lambda p:_make_judge(p[0], p[1]), [([4,1],3), ([3,2],4), ([3,1,1],7), ([2,2,1],8), ([2,1,1,1],9)]))

def isStraight(groupedDeck):
    if len(groupedDeck) < 5:
        return None
    elif decode(groupedDeck[0][0][1]) - decode(groupedDeck[4][0][1]) == 4:
        return Deck(6, _pick_from_groupedDeck(groupedDeck))
    elif groupedDeck[0][0][1] == 'A' and groupedDeck[1][0][1] == '5':
        return Deck(6, _pick_from_groupedDeck(groupedDeck[1:] + [groupedDeck[0]]))
    else:
        return None

def makeDeck(deck):
    checkValidDeck(deck)
    groupedDeck=groupByRank(deck)

    # although i tried very hard in expressing in some kind of lazy-eval mode, \n
    # but failed unfortunately at last & ended up in such an eagerly, ugly & inefficient way..
    l=list(filter(lambda x:x, map(lambda judge:judge(groupedDeck), _trivial_judge)))
    if len(l) > 0:
        return l[0]
    else:
        d1=_make_judge([1,1,1,1,1], 5, aux=lambda groupedDeck: len(set(map(lambda s: s[0][0], groupedDeck))) == 1)(groupedDeck)
        d2=isStraight(groupedDeck)
        return Deck(2, d2.deck) if d1 and d2 else (d1 if d1 else (d2 if d2 else Deck(10, _pick_from_groupedDeck(groupedDeck))))

def pickHigh(deck):
    if len(deck) != 7:
        raise Exception("invalid games!")
    return reduce(lambda s, a: s if s >= a else a, map(lambda d: makeDeck(d), combinations(deck, 5)))

def isStraight2(s_deck):
    l=list(filter(lambda x: decode(s_deck[x][1])-decode(s_deck[x+4][1]) == 4, range(len(s_deck)-4)))
    return s_deck[l[0]:l[0]+5] if len(l) > 0 else \
        (s_deck[-4:] + [s_deck[0]] if (s_deck[0][1]=='A' and s_deck[-4][1]=='5') else None)

def fastPickHigh(deck):
    if len(deck) != 7:
        raise Exception("invalid games!")
    suit=Counter(map(lambda x:x[0], deck)).most_common(1)[0]
    isFlush=(suit[1]>=5)
    if isFlush:
        s=sorted(filter(lambda x:suit[0] in x, deck), key=lambda x:decode(x[1]), reverse=True)
        isFlush=s[:5]
        d=isStraight2(s)
        if d:
            return Deck(2, d)
    groupedDeck=groupByRank(deck)
    if len(groupedDeck[0]) == 4:
        return Deck(3, [groupedDeck[0][0], sorted(_pick_from_groupedDeck(groupedDeck[1:]), key=lambda x:decode(x[1]), reverse=True)[0]])
    elif len(groupedDeck[0]) == 3 and len(groupedDeck[1]) >= 2:
        return Deck(4, [groupedDeck[0][0], groupedDeck[1][0]])
    elif isFlush:
        return Deck(5, isFlush)
    d=isStraight2(sorted(_pick_from_groupedDeck(groupedDeck), key=lambda x:decode(x[1]), reverse=True))
    if d:
        return Deck(6, d)
    elif len(groupedDeck[0]) == 3:
        return Deck(7, [groupedDeck[0][0]] + sorted(_pick_from_groupedDeck(groupedDeck[1:]), key=lambda x:decode(x[1]), reverse=True)[:2]) 
    elif len(groupedDeck[0]) == 2 and len(groupedDeck[1]) == 2:
        return Deck(8, [groupedDeck[0][0], groupedDeck[1][0], sorted(_pick_from_groupedDeck(groupedDeck[2:]), key=lambda x:decode(x[1]), reverse=True)[0]])
    elif len(groupedDeck[0]) == 2:
        return Deck(9, [groupedDeck[0][0]] + sorted(_pick_from_groupedDeck(groupedDeck[1:]), key=lambda x:decode(x[1]), reverse=True)[:3])
    else:
        return Deck(10, _pick_from_groupedDeck(groupedDeck)[:5])

@total_ordering
class Deck:
    def __init__(self, rank, deck):
        self.rank = rank
        self.deck = deck

    def __lt__(self, oppo):
        if self.rank != oppo.rank:
            return self.rank > oppo.rank
        else:
            for i in range(0, len(self.deck)):
                if self.deck[i][1] != oppo.deck[i][1]:
                    return decode(self.deck[i][1]) < decode(oppo.deck[i][1])
            return False
    
    def __eq__(self, oppo):
        return self.rank == oppo.rank and list(map(lambda c:c[1], self.deck)) == list(map(lambda c:c[1], oppo.deck))
