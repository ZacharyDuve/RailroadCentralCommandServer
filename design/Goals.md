# Goals

These are what features that we would like Central Command Server to have.

This is a living document that is expected to change over time.

## Coding Must haves

- All APIs that are served will be driven by a spec document
- Code will be documented and tested where makes sense

## Must have features

### Block control

This is a mountain of requirements that has been a thorn in my side...

The initial problem is that I need a way to route trains on the railroad. This means that I need to be able to do the following:

1) Know about turnouts (switches) on the railroad. This includes what type the turnout is as there are many different [types](https://www.trackopedia.com/en/encyclopedia/infrastructure/turnouts/types-of-turnouts/)
2) Being able to control the turnouts. This includes being able to connect to the Switch Machine Driver Servers
3) Connecting the idea of a logical turnout into set of switch machine drivers that set the positions. Some turnouts have only one set of points which means the relationship is 1 to 1 for turnout to driver but some turnouts have a 1 turnout to multiple switch machines drivers. Also need to factor in that some like a 3 way wye have overlapping points so that the points need to be driven in a given order.
4) Figure out how some nuanced routing device such as a turntable or transfer table are set. Are they turnouts?
    - They can connect track segments together though by their nature they can connect and disconnect parts of the railroad. The train (locomotive) needs to be on it before it can change position were true turnouts are set before the train, though with overlapping schedules that switches shouldn't throw all of the way through.
5) Have some idea of blocks. Section of track that are deemed as allowing a single train to pass through. Blocks have signals at ends. Also blocks need to know about the adjacent blocks to allow for correct signalling. 

#### Blocks

Technically a block isn't just between a section with 2 or more signals like so

Signal A            Signal B
 ==========================

Once a turnout is involved then the block goes from the signal to the junction

Signal A      J    Signal B
 ==========================
              \============
                   Signal C

In the case with junction, side not set will automatically set signal to STOP. In the prior example if the route is set A <-> B through the junction than Signals A and B will be set accordingly to allow through. Signal C will be set to STOP