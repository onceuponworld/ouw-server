# design

player actions and npc actions are separated, architecturally speaking, npc actions don't need to go over the WAN, they should be internal calls.  this allows for a more focused effort to process player requests.