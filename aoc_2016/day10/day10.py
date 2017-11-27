import re

def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

agents = {}
todo = []

def getAgent(agentDef):
    aId = ''.join(map(str, agentDef))

    if agents.get(aId) == None: agents[aId] = Agent(aId)

    return agents[aId]

class Agent:
    def __init__(self, i):
        self.id = i
        self.low = None
        self.high = None
        self.cargo = []

    def setLow(self, agent): self.low = agent

    def setHigh(self, agent): self.high = agent

    def grab(self, mchip):
        self.cargo.append(int(mchip))

        if len(self.cargo) == 2: todo.append(self.id)

    def deliver(self):
        v1, v2 = sorted(self.cargo)
        self.low.grab(v1)
        self.high.grab(v2)

        # 1 star
        if v1==17 and v2==61: print self.id

    def __str__(self):
        return "%s gives low to %s and high to %s" % (self.id, self.low.id, self.high.id)

for cmd in io().split('\n'):
    m = re.findall(r'(value|bot|output) (\d+)', cmd, re.M|re.I)

    t, v = m[0]

    if t == 'value':
        bot = getAgent(m[1])
        bot.grab(v)
    else:
        bot = getAgent(m[0])
        bot.setLow(getAgent(m[1]))
        bot.setHigh(getAgent(m[2]))

for botId in todo:
    agents[botId].deliver()

# 2 star
print reduce(lambda x, y: x * y, (agents['output'+str(k)].cargo[0] for k in [0,1,2]))
