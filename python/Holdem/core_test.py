import unittest
import core
import calculate
import time
from itertools import product
from itertools import combinations

class TestCore(unittest.TestCase):

    def test_checkValidCard(self):
        self.assertTrue(not core.checkValidCard("X"))
        self.assertTrue(not core.checkValidCard("HS"))
        self.assertTrue(not core.checkValidCard("H10"))
        self.assertTrue(core.checkValidCard("HT"))
        self.assertTrue(not core.checkValidCard("s4"))
        self.assertTrue(core.checkValidCard("S4"))
        self.assertTrue(not core.checkValidCard("S_4"))

    def test_checkValidDeck(self):
        with self.assertRaises(Exception):
            core.checkValidDeck(["X2", "HA", "H5"])
        with self.assertRaises(Exception):
            core.checkValidDeck(["X2", "HA", "H5", "H3", "HK"])
        with self.assertRaises(Exception):
            core.checkValidDeck(["H2", "HA", "H10", "H3", "HK"])
        with self.assertRaises(Exception):
            core.checkValidDeck(["H2", "H2", "H5", "H3", "HK"])
        with self.assertRaises(Exception):
            core.checkValidDeck(["X2", "HA", "H5", "H3", "HK"])
        self.assertTrue(not core.checkValidDeck(["H2", "HA", "H5", "H3", "HK"]))

    # def test_isFlush(self):
    #     self.assertTrue(core.isFlush(core.shuffle(["X2", "HA", "H5", "H3", "HK"])) == None)
    #     self.assertTrue(core.isFlush(core.shuffle(["X2", "XA", "X5", "X3", "X4"])) == ["XA", 'X5', 'X4', 'X3', 'X2'])

    # def test_isStraight(self):
    #     self.assertTrue(core.isStraight(core.shuffle(["H4", "VA", "H6", "X2", "X5"])) == None)
    #     self.assertTrue(core.isStraight(core.shuffle(["XQ", "VA", "HT", "XJ", "XK"])) == ["VA", 'XK', 'XQ', 'XJ', 'HT'])
    #     self.assertTrue(core.isStraight(core.shuffle(["H4", "VA", "H3", "X2", "X5"])) == ["X5", 'H4', 'H3', 'X2', 'VA'])

    # def test_isFourOfAKind(self):
    #     self.assertTrue(core.isFourOfAKind(core.shuffle(["XQ", "VA", "HT", "XJ", "XK"])) == None)
    #     self.assertTrue(core.isFourOfAKind(core.shuffle(["XJ", "VQ", "HJ", "XQ", "_Q"])) == None)
    #     self.assertTrue(core.isFourOfAKind(core.shuffle(["XJ", "VQ", "HQ", "XQ", "XQ"])) == ["VQ", 'XJ'])
    #     self.assertTrue(core.isFourOfAKind(core.shuffle(["HJ", "H2", "DJ", "CJ", "SJ"])) == ["HJ", 'H2'])

    # def test_isFullHouse(self):
    #     self.assertTrue(core.isFullHouse(core.shuffle(["XQ", "VA", "HT", "XJ", "XK"])) == None)
    #     self.assertTrue(core.isFullHouse(core.shuffle(["XJ", "VQ", "HQ", "XQ", "_Q"])) == None)
    #     self.assertTrue(core.isFullHouse(core.shuffle(["XJ", "VQ", "HJ", "XQ", "_Q"])) == ["VQ", 'XJ'])

    # def test_isThreeOfAKind(self):
    #     self.assertTrue(core.isThreeOfAkind(core.shuffle(["XQ", "VA", "HQ", "XQ", "XA"])) == None)
    #     self.assertTrue(core.isThreeOfAkind(core.shuffle(["XQ", "V3", "HQ", "X3", "XA"])) == None)
    #     self.assertTrue(core.isThreeOfAkind(core.shuffle(["XQ", "VQ", "HQ", "XQ", "XA"])) == None)
    #     self.assertTrue(core.isThreeOfAkind(core.shuffle(["XQ", "VQ", "H7", "X6", "XA"])) == None)
    #     self.assertTrue(core.isThreeOfAkind(core.shuffle(["XQ", "V3", "HQ", "XQ", "XA"])) == ["XQ", 'XA', 'V3'])

    # def test_isTwoPair(self):
    #     self.assertTrue(core.isTwoPair(core.shuffle(["XQ", "V3", "H6", "XK", "XA"])) == None)
    #     self.assertTrue(core.isTwoPair(core.shuffle(["XQ", "V3", "HQ", "XQ", "XA"])) == None)
    #     self.assertTrue(core.isTwoPair(core.shuffle(["XQ", "VQ", "HQ", "XQ", "XA"])) == None)
    #     self.assertTrue(core.isTwoPair(core.shuffle(["XQ", "VQ", "H7", "X6", "XA"])) == None)
    #     self.assertTrue(core.isTwoPair(core.shuffle(["XQ", "V3", "HQ", "XA", "X3"])) == ["XQ", "V3", "XA"])
    
    # def test_isOnePair(self):
    #     self.assertTrue(core.isOnePair(core.shuffle(["XQ", "V3", "H6", "XK", "XA"])) == None)
    #     self.assertTrue(core.isOnePair(core.shuffle(["XQ", "V3", "HQ", "XQ", "XA"])) == None)
    #     self.assertTrue(core.isOnePair(core.shuffle(["XQ", "VQ", "HQ", "XQ", "XA"])) == None)
    #     self.assertTrue(core.isOnePair(core.shuffle(["XQ", "V3", "HQ", "XA", "X3"])) == None)
    #     self.assertTrue(core.isOnePair(core.shuffle(["XQ", "V3", "HQ", "XA", "X5"])) == ["XQ", "XA", "X5", "V3"])

    # def test_evaluate(self):
    #     with self.assertRaises(Exception):
    #         core.evaluate(["X2", "HA", "H5"])
    #     with self.assertRaises(Exception):
    #         core.evaluate(["H2", "H2", "H5", "H3", "HK"])
    #     self.assertTrue(core.evaluate(["H2", "H4", "H5", "H3", "HA"]) == (2, ["H5", 'H4', 'H3', 'H2', 'HA']))
    #     print("\nHit: " + str(core.interprete(core.evaluate(["H2", "H4", "H5", "H3", "HA"]))))
    #     self.assertTrue(core.evaluate(["H2", "H4", "H5", "S3", "HA"]) == (6, ["H5", 'H4', 'S3', 'H2', 'HA']))
    #     print("Hit: " + str(core.interprete(core.evaluate(["H2", "H4", "H5", "S3", "HA"]))))
    #     self.assertTrue(core.evaluate(["S5", "S4", "H5", "H3", "D5"]) == (7, ["S5", 'S4', 'H3']))
    #     print("Hit: " + str(core.interprete(core.evaluate(["S5", "S4", "H5", "H3", "D5"]))))
    #     self.assertTrue(core.evaluate(["S5", "S4", "H5", "H4", "D5"]) == (4, ["S5", 'S4']))
    #     print("Hit: " + str(core.interprete(core.evaluate(["S5", "S4", "H5", "H4", "D5"]))))
    #     self.assertTrue(core.evaluate(["SQ", "S3", "D6", "HT", "HA"]) == (10, ["HA", "SQ", "HT", "D6", "S3"]))
    #     print("Hit: " + str(core.interprete(core.evaluate(["HA", "SQ", "HT", "D6", "S3"]))))
    
    def test_compare(self):
        self.assertTrue(core.makeDeck(["H2", "H4", "H5", "H3", "HA"]) > core.makeDeck(["S5", "S4", "H5", "H3", "D5"]))
        self.assertTrue(core.makeDeck(["H2", "H4", "H5", "H3", "HA"]) > core.makeDeck(["S2", "H5", "S3", "SA", 'S4']))
        self.assertTrue(core.makeDeck(["H2", "H4", "H5", "H3", "HA"]) == core.makeDeck(["S2", "S5", "S3", "SA", 'S4']))
        self.assertTrue(core.makeDeck(["SJ", "SQ", "HJ", "CQ", "DQ"]) > core.makeDeck(["S2", "S8", "S3", "SA", 'S4']))
        self.assertTrue(core.makeDeck(["HQ", "SQ", "HJ", "CQ", "DQ"]) > core.makeDeck(["S2", "S8", "S3", "SA", 'S4']))
        self.assertTrue(core.makeDeck(["SJ", "SQ", "HJ", "CQ", "DQ"]) < core.makeDeck(["HQ", "SQ", "HJ", "CQ", "DQ"]))
        self.assertTrue(core.makeDeck(["SJ", "SQ", "S2", "S4", 'S7']) > core.makeDeck(["S2", "H5", "S3", "SA", 'D4']))
        self.assertTrue(core.makeDeck(["C4", "H4", "D5", "DT", "CT"]) > core.makeDeck(["H2", "DA", "CA", "S5", "S4"]))
        self.assertTrue(core.makeDeck(["H2", "H4", "H5", "HT", "HQ"]) > core.makeDeck(["CA", "SA", "SJ", "C7", "S2"]))
        self.assertTrue(core.makeDeck(["HT", "HJ", "HQ", "HK", "HA"]) >core.makeDeck(["CA", "H4", "SJ", "C7", "H2"]))
        self.assertTrue(core.makeDeck(["CT", "DJ", "DQ", "DK", "DA"]) >core.makeDeck(["CA", "CJ", "SJ", "C7", "S3"]))
        self.assertTrue(core.makeDeck(["C4", "HJ", "DQ", "HK", "DA"]) <core.makeDeck(["CA", "DT", "SJ", "C7", "D7"]))
        self.assertTrue(core.makeDeck(["C4", "C6", "C2", "C9","CT"]) >core.makeDeck(["CA", "DA", "D6", "D7", "S4"]))
        self.assertTrue(core.makeDeck(["C5", "C6", "C2", "C4", "C3"]) >core.makeDeck(["CA", "HA", "DA", "S7", "D7"]))
        self.assertTrue(core.makeDeck(["HA", "SA", "D8", "C8", "S8"]) < core.makeDeck(["C5", "C6", "C2", "C4", "C3"]))
        self.assertTrue(core.makeDeck(["H2", "H4", "H5", "HT", "H6"]) > core.makeDeck(["C4", "D5", "ST", "C6", "S2"]))
        self.assertTrue(core.makeDeck(["HQ", "HA", "HJ", "HK", "HT"]) ==core.makeDeck(["DQ", "DA", "DJ", "DK", "DT"]))
        self.assertTrue(core.makeDeck(["HA", "HQ", "HJ", "HK", "HT"])>core.makeDeck(["CA", "DA", "D8", "C8", "H8"]))
        self.assertTrue(core.makeDeck(["C9", "DA", "D8", "C8", "H8"]) <core.makeDeck(["C5", "C6", "C2", "C4", "C3"]))
        self.assertTrue(core.makeDeck(["C4", "DA", "H4", "C8", "H8"]) <core.makeDeck(["C5", "C6", "D8", "D4", "C7"]))
        self.assertTrue(core.makeDeck(["HA", "CA", "S8", "D8", "C8"]) >core.makeDeck(["H2", "DA", "SA", "S5", "S4"]))
        self.assertTrue(core.makeDeck(["H8", "H9", "HT", "HJ", 'HQ']) > core.makeDeck(["H5", "H6", "S7", "S8", 'S9']))
        self.assertTrue(core.makeDeck(["C2", "S2", "H2", "D2", "ST"]) < core.makeDeck(["H5", "H6", "H7", "H8", 'H9']))     

        self.assertTrue(core.makeDeck(["HA", "HK", "DQ", "DT", "C5"]) > core.makeDeck(["HA", "DJ", "C3", "S6", "S9"]))
        self.assertTrue(core.makeDeck(["HA", "HK", "DQ", "DT", "CJ"]) > core.makeDeck(["SK", "DJ", "CQ", "ST", "H9"]))
        self.assertTrue(core.makeDeck(["H7", "H6", "H5", "D4", "C3"]) > core.makeDeck(["HA", "DA", "C9", "ST", "S9"]))
        # self.assertTrue(core.makeDeck(["HA", "HA", "CT", "DT", "HT"]) > core.makeDeck(["HJ", "DJ", "C3", "S3", "S9"]))
        # self.assertTrue(core.makeDeck(["S9", "H9", "D9", "D6", "C5"]) > core.makeDeck(["HA", "HA", "C3", "S6", "HT"]))
        self.assertTrue(core.makeDeck(["HK", "SK", "D9", "DT", "C4"]) > core.makeDeck(["HA", "DK", "C5", "S6", "S9"]))
        # self.assertTrue(core.makeDeck(["H3", "H9", "DQ", "D6", "C5"]) > core.makeDeck(["HJ", "D8", "CQ", "S6", "S9"]))
        self.assertTrue(core.makeDeck(["H3", "H6", "D3", "D6", "C5"]) > core.makeDeck(["HA", "DA", "CK", "ST", "H8"]))
        self.assertTrue(core.makeDeck(["SA", "SK", "SQ", "ST", "SJ"]) > core.makeDeck(["HK", "HQ", "HJ", "HT", "H9"]))

    def test_pickHigh(self):
        self.assertTrue(core.pickHigh(["H2", 'SA', "H4", "H5", "H3", 'DA', "HA"]) == core.makeDeck(["C4", "C2", "C3", "CA", "C5"]))
    
    def test_isStraight2(self):
        self.assertTrue(core.isStraight2(["H2", 'S7', "H6", "H5", "H4", 'D3', "HA"]))
        self.assertTrue(core.isStraight2(["HA", 'S9', "H8", "H5", "H4", 'D3', "H2"]))

    def test_fastPickHigh(self):
        self.assertTrue(core.fastPickHigh(['DK', 'CK', 'C8', 'D8', 'SK', 'HK', 'SQ']) == core.pickHigh(['DK', 'CK', 'C8', 'D8', 'SK', 'HK', 'SQ']))
        self.assertTrue(core.fastPickHigh(['CT', 'ST', 'D8', 'H7', 'HJ', 'DT', 'HQ']) == core.pickHigh(['CT', 'ST', 'D8', 'H7', 'HJ', 'DT', 'HQ']))
        self.assertTrue(core.fastPickHigh(['H9', 'C3', 'DQ', 'H8', 'HA', 'S7', 'S4']) == core.pickHigh(['H9', 'C3', 'DQ', 'H8', 'HA', 'S7', 'S4']))       
        self.assertTrue(core.fastPickHigh(['H8', 'D8', 'DA', 'CQ', 'DT', 'SK', 'CJ']) == core.pickHigh(['H8', 'D8', 'DA', 'CQ', 'DT', 'SK', 'CJ']))       
        # for comb in list(combinations(set(map(lambda t: t[0]+t[1], product(core.ALL_SUITS, core.ALL_RANK))), 7))[:2000]:
        #     r = core.fastPickHigh(comb) == core.pickHigh(comb)
        #     if r:
        #         self.assertTrue(True)
        #     else:
        #         print(comb)

if __name__ == '__main__':
    unittest.main()