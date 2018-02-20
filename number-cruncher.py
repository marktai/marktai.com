#!/usr/bin/python3
from pprint import pprint
import random
import pdb
random.seed(1368432912587494553158932)

def oop_shuffle(l, needs_change=False):
    while True:
        ret = random.sample(l, len(l))
        if needs_change and tuple(l) == tuple(ret):
            continue
        break
    return ret

def gen_unique_permutations_using_blocks(blocks, n, pred=None):
    perms = set()
    while len(perms) < n:
        new = ''.join(random.sample(blocks, 5))
        if pred is None or pred(new):
            perms.add(new)
    return oop_shuffle(list(perms))

# selects are in the format [ { 'abc', '123', ... }, {'213', 'abc', ... }, ... ]
# i actually need currying here :(
def can_password_happen(selects, p):
    p_triples = [p[3*i:3*i+3] for i in range(len(p)//3)]
    for p_triple, s in zip(p_triples, selects):
        if ''.join(p_triple) not in s:
            return False
    return True

code = 'XR5PB-EX0YN-WETZ8'

code_letters = list(code.replace('-', ''))

shuffled_code_letters = list(code_letters)
random.shuffle(shuffled_code_letters)
card_order = list(shuffled_code_letters)

triples = [code_letters[3*i:3*i+3] for i in range(len(code_letters)//3)]

num_choices_in_select = 6
selects = [[''.join(triple)] + [''.join(random.sample(code_letters, 3)) for _ in range(num_choices_in_select-1)] for triple in triples]
selects_set = [set(t) for t in selects]
selects_randomized = [oop_shuffle(s) for s in selects]

possible_passwords = gen_unique_permutations_using_blocks([block for blocks in selects_set for block in blocks], 500, lambda p: not can_password_happen(selects_set, p)) + [''.join(code_letters)]
random.shuffle(possible_passwords)

pprint(selects_randomized)
pprint(list(map(lambda x: ''.join(x), possible_passwords)))


pdb.set_trace()
