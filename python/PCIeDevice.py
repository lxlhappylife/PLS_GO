
from ctypes import *
from numpy.ctypeslib import ndpointer
lib_pcie = cdll.LoadLibrary('./pcieutils.so')


class PCIeDevice(object):
    def __init__(self, bus, device, function):
        self.bus = bus
        self.device = device
        self.function = function

    def getHostBDF(self):
        lib_pcie.GetHostBDF.argtypes = [c_uint8, c_uint8, c_uint8]
        lib_pcie.GetHostBDF.restype = ndpointer(dtype = c_uint8, shape = (3,))
        RC = lib_pcie.GetHostBDF(self.bus,  
                                 self.device,
                                 self.function)
        print("%x:%x.%x" % (RC[0], RC[1], RC[2]))


if __name__ == "__main__":
    dut = PCIeDevice(0xa, 0x0, 0x0)
    dut.getHostBDF()
