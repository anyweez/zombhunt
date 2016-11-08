LITTLE ENDIAN
- Good way to find item names is to find them in a recipe.
- 0x80 3F 62 03 40 3F seems to be a boundary of some sort. Occurs quite a bit.
- Offset 1076 for name, terminated in 0x00 ?
    - Conditions before name, so name is not at stable position

* HEADER
- 9x 0xC842
- First 8 bytes are the same, at least in A15


- 9 empty bytes before most conditions (+thirst, +water, +radiation, beers, @wellness)
- condition list (in order):
    - +water
    - -water
    - waterLevel
    - +radiation
    - -radiation
    - +starve
    - +full
    - +thirst
    - +hydrated
    - beers
    - food
    - water
    - health
    - coretemp
    - +heatSource
    - radiation
    - @wellness
    - drink
- condition: lenth + name + 21 bytes

- Notifications:
    - BuffEntityUINotification, WaterBuffNotification
    - Name + 16 bytes
    - isPlayer + 21 or 33 bytes (12 blank bytes + more)

- Profile
    - Name + 7 bytes (values, not sure what). All values the same except last one (sam's = 4, mine = 8 in multiple games)



* ITEMS
- belt first? first nine slots
- inventory second?

- 0x0800 (8, # of belt slots) + belt + 0x2000 (32, # of inventory slots) + 32 slots

BE 10 00 00 00 00 00 00 00 00 00 00 01 00           small stone         (0x00BE)
14 15 00 00 00 00 00 00 00 00 00 00 02 00           grass               (0x0514)

BD 12 00 00 00 00 00 00 00 00 00 00 02 00           yucca fruit x2      (0x02BD)
23 15 00 00 00 00 00 00 00 00 00 00 02 00           Aloe vera plant x2  (0x0523)
BE 10 00 00 00 00 00 00 00 00 00 00 02 00           small stone x2      (0x00BE)
14 15 00 00 00 00 00 00 00 00 00 00 36 00           plant fibers x54    (0x0514)

0C 13 00 00 00 00 00 00 00 00 00 00 01 00           bottled water       (0x030c)
C9 12 00 00 00 00 00 00 00 00 00 00 01 00           can of chili        (0x02c9)

01      02      03      04      05      06      07      08      09      10      11      12      13      14
id     id     0x00    0x00    0x00    0x00    0x00    0x00    0x00    0x00    0x00    count?  count   0x00

- ID + 10 bytes + quantity + other stuff
- Byte 13 = quantity

- 0x03 + 14 empty bytes
    - rock = 17 bytes 
    - grass = 14 bytes

- questions:
    - which byte is the second count byte?




**** extra from game ****

02  Bottled Water
03  Can of chili
04  First aid bandage
05  Land claim block


02  Yucca fruit x2
03  Aloe vera plant x2 
04  Small stone x2 
06  Plant fibers x54

wellness 100
food 93
water 93