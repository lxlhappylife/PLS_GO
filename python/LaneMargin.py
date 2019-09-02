from ctypes import *

lib1 = cdll.LoadLibrary('./PcieLaneMarginCmds.so')
class LaneMargin(object):
    def __init__(self,bus,device,function):
        self.bus = bus
        self.device = device
        self.function = function

    def Precondition():
        pass

