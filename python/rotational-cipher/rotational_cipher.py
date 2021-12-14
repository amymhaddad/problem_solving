from string import ascii_lowercase as lower, ascii_uppercase as upper

ALPHA_LENGTH = 26
START = 97

def rotate(text, key):
   # import pdb; pdb.set_trace()
    shift = ord(text) - START
 
    shift_val = (shift + key) % ALPHA_LENGTH

    new_val = shift_val + START

    return chr(new_val)
